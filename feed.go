package gocrawlhq

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Feed(size int) (feedResponse *FeedResponse, err error) {
	expectedStatusCode := 200
	feedResponse = new(FeedResponse)

	// build request
	req, err := http.NewRequest("GET", FeedEndpoint.String(), nil)
	if err != nil {
		return feedResponse, err
	}

	req.Header.Add("X-Auth-Key", c.Key)
	req.Header.Add("X-Auth-Secret", c.Secret)
	req.Header.Add("User-Agent", "gocrawlhq/"+Version)

	// execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return feedResponse, err
	}
	defer resp.Body.Close()

	// check response status code
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
