package gocrawlhq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Discovered(URLs []string, URLType string, bypassSeencheck bool) (discoveredResponse *DiscoveredResponse, err error) {
	expectedStatusCode := 201
	discoveredResponse = new(DiscoveredResponse)

	// build payload
	var URLsPayload []URL
	for _, URLString := range URLs {
		URLsPayload = append(URLsPayload, URL{Value: URLString})
	}

	payload := DiscoveredPayload{
		Project:         c.Project,
		BypassSeencheck: bypassSeencheck,
		Type:            URLType,
		URLs:            URLsPayload,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return discoveredResponse, err
	}

	// build request
	req, err := http.NewRequest("POST", DiscoveredEndpoint.String(), bytes.NewReader(jsonPayload))
	if err != nil {
		return discoveredResponse, err
	}

	req.Header.Add("X-Auth-Key", c.Key)
	req.Header.Add("X-Auth-Secret", c.Secret)
	req.Header.Add("User-Agent", "gocrawlhq/"+Version)

	// execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return discoveredResponse, err
	}
	defer resp.Body.Close()

	// check response status code
	if resp.StatusCode != expectedStatusCode {
		return discoveredResponse, fmt.Errorf("non-%d status code: %d", expectedStatusCode, resp.StatusCode)
	}

	// decode response body
	err = json.NewDecoder(resp.Body).Decode(discoveredResponse)
	if err != nil {
		return discoveredResponse, err
	}

	return discoveredResponse, err
}
