package gocrawlhq

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (c *Client) Feed(size int) (feedResponse *FeedResponse, err error) {
	expectedStatusCode := 200
	emptyStatusCode := 204
	feedResponse = new(FeedResponse)

	// build request
	req, err := http.NewRequest("GET", FeedEndpoint.String(), nil)
	if err != nil {
		return feedResponse, err
	}

	q := req.URL.Query()
	q.Add("size", strconv.Itoa(size))
	req.URL.RawQuery = q.Encode()

	req.Header.Add("X-Auth-Key", c.Key)
	req.Header.Add("X-Auth-Secret", c.Secret)

	if c.Identifier != "" {
		req.Header.Add("X-Identifier", c.Identifier)
	}

	req.Header.Add("User-Agent", "gocrawlhq/"+Version)

	// execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return feedResponse, err
	}
	defer resp.Body.Close()

	// check response status code for 'empty' or 204
	if resp.StatusCode == emptyStatusCode {
		return feedResponse, errors.New("gocrawlhq: feed is empty")
	}

	// check response status code for 200
	if resp.StatusCode != expectedStatusCode {
		return feedResponse, fmt.Errorf("non-%d status code: %d", expectedStatusCode, resp.StatusCode)
	}

	// decode response body
	err = json.NewDecoder(resp.Body).Decode(feedResponse)
	if err != nil {
		return feedResponse, err
	}

	return feedResponse, err
}
