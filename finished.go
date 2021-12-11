package gocrawlhq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Finished(URLs []URL, URLType string) (finishedResponse *FinishedResponse, err error) {
	expectedStatusCode := 200
	finishedResponse = new(FinishedResponse)

	// build payload
	payload := FinishedPayload{
		Project: c.Project,
		URLs:    URLs,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return finishedResponse, err
	}

	// build request
	req, err := http.NewRequest("POST", finishedEndpoint, bytes.NewReader(jsonPayload))
	if err != nil {
		return finishedResponse, err
	}

	req.Header.Add("X-Auth-Key", c.Key)
	req.Header.Add("X-Auth-Secret", c.Secret)
	req.Header.Add("User-Agent", "gocrawlhq/"+Version)

	// execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return finishedResponse, err
	}
	defer resp.Body.Close()

	// check response status code
	if resp.StatusCode != expectedStatusCode {
		return finishedResponse, fmt.Errorf("non-%d status code: %d", expectedStatusCode, resp.StatusCode)
	}

	// decode response body
	err = json.NewDecoder(resp.Body).Decode(finishedResponse)
	if err != nil {
		return finishedResponse, err
	}

	return finishedResponse, err
}
