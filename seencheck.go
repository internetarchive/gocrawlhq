package gocrawlhq

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (c *Client) Seencheck(URLs []URL) (outputURLs []URL, err error) {
	expectedStatusCodes := []int{200, 204}

	jsonPayload, err := json.Marshal(URLs)
	if err != nil {
		return URLs, err
	}

	req, err := http.NewRequest("POST", c.SeencheckEndpoint.String(), bytes.NewReader(jsonPayload))
	if err != nil {
		return URLs, err
	}

	req.Header.Add("X-Auth-Key", c.Key)
	req.Header.Add("X-Auth-Secret", c.Secret)
	req.Header.Add("User-Agent", "gocrawlhq/"+Version)

	if c.Identifier != "" {
		req.Header.Add("X-Identifier", c.Identifier)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return URLs, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&outputURLs)
	if err != nil {
		return URLs, err
	}

	for _, expectedStatusCode := range expectedStatusCodes {
		if resp.StatusCode == expectedStatusCode {
			return outputURLs, nil
		}
	}

	return outputURLs, err
}
