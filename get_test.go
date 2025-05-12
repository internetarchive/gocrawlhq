package gocrawlhq

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_Get_Success(t *testing.T) {
	expectedKey := "test-key"
	expectedSecret := "test-secret"
	expectedIdentifier := "test-id"
	expectedUserAgent := "gocrawlhq/" + Version
	expectedSize := 2

	mockURLs := []URL{
		{Value: "https://example.com"},
		{Value: "https://test.com"},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check method
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET, got %s", r.Method)
		}
		// Check headers
		if r.Header.Get("X-Auth-Key") != expectedKey {
			t.Error("Missing or incorrect X-Auth-Key")
		}
		if r.Header.Get("X-Auth-Secret") != expectedSecret {
			t.Error("Missing or incorrect X-Auth-Secret")
		}
		if r.Header.Get("User-Agent") != expectedUserAgent {
			t.Error("Missing or incorrect User-Agent")
		}
		if r.Header.Get("X-Identifier") != expectedIdentifier {
			t.Error("Missing or incorrect X-Identifier")
		}
		// Check query param
		if r.URL.Query().Get("size") != strconv.Itoa(expectedSize) {
			t.Error("Missing or incorrect 'size' query param")
		}
		// Respond with mock data
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockURLs)
	}))
	defer ts.Close()

	parsedURL, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	client := &Client{
		Key:          expectedKey,
		Secret:       expectedSecret,
		Identifier:   expectedIdentifier,
		URLsEndpoint: parsedURL,
		HTTPClient:   http.DefaultClient,
	}

	urls, err := client.Get(context.Background(), expectedSize)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(urls) != expectedSize {
		t.Errorf("Expected %d URLs, got %d", expectedSize, len(urls))
	}
	if urls[0].Value != mockURLs[0].Value {
		t.Errorf("Unexpected URL[0]: %v", urls[0])
	}
}

func TestClient_Get_Empty(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent) // 204
	}))
	defer ts.Close()

	parsedURL, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	client := &Client{
		Key:          "k",
		Secret:       "s",
		URLsEndpoint: parsedURL,
		HTTPClient:   http.DefaultClient,
	}

	_, err = client.Get(context.Background(), 5)
	if !errors.Is(err, ErrFeedEmpty) {
		t.Errorf("Expected feed empty error, got: %v", err)
	}
}

func TestClient_Get_Non200(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError) // 500
	}))
	defer ts.Close()

	parsedURL, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	client := &Client{
		Key:          "k",
		Secret:       "s",
		URLsEndpoint: parsedURL,
		HTTPClient:   http.DefaultClient,
	}

	_, err = client.Get(context.Background(), 5)
	if err == nil || err.Error() != "non-200 status code: 500" {
		t.Errorf("Expected status code error, got: %v", err)
	}
}
