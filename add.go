package gocrawlhq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Add(URLs []URL, bypassSeencheck bool) (err error) {
	expectedStatusCode := 201

	// build payload
	var URLsPayload []URL

	URLsPayload = append(URLsPayload, URLs...)

	payload := AddPayload{
		BypassSeencheck: bypassSeencheck,
		URLs:            URLsPayload,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// build request
	req, err := http.NewRequest(http.MethodPost, c.URLsEndpoint.String(), bytes.NewReader(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Add("X-Auth-Key", c.Key)
	req.Header.Add("X-Auth-Secret", c.Secret)
	req.Header.Add("User-Agent", "gocrawlhq/"+Version)

	if c.Identifier != "" {
		req.Header.Add("X-Identifier", c.Identifier)
	}

	// execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check response status code
	if resp.StatusCode != expectedStatusCode {
		return fmt.Errorf("non-%d status code: %d", expectedStatusCode, resp.StatusCode)
	}

	return err
}
