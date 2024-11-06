package gocrawlhq

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (c *Client) Feed(size int, strategy string) (URLs []URL, err error) {
	expectedStatusCode := 200
	emptyStatusCode := 204

	// build request
	req, err := http.NewRequest("GET", c.URLsEndpoint.String(), nil)
	if err != nil {
		return URLs, err
	}

	q := req.URL.Query()
	q.Add("size", strconv.Itoa(size))
	q.Add("strategy", strategy)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("X-Auth-Key", c.Key)
	req.Header.Add("X-Auth-Secret", c.Secret)
	req.Header.Add("User-Agent", "gocrawlhq/"+Version)

	if c.Identifier != "" {
		req.Header.Add("X-Identifier", c.Identifier)
	}

	// execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return URLs, err
	}
	defer resp.Body.Close()

	// check response status code for 'empty' or 204
	if resp.StatusCode == emptyStatusCode {
		return URLs, errors.New("gocrawlhq: feed is empty")
	}

	// check response status code for 200
	if resp.StatusCode != expectedStatusCode {
		return URLs, fmt.Errorf("non-%d status code: %d", expectedStatusCode, resp.StatusCode)
	}

	// decode response body
	err = json.NewDecoder(resp.Body).Decode(URLs)
	if err != nil {
		return URLs, err
	}

	return URLs, err
}
