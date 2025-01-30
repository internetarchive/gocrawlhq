package gocrawlhq

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func (c *Client) Seencheck(ctx context.Context, URLs []URL) (outputURLs []URL, err error) {
	jsonPayload, err := json.Marshal(URLs)
	if err != nil {
		return URLs, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.SeencheckEndpoint.String(), bytes.NewReader(jsonPayload))
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

	if resp.StatusCode == 200 {
		err = json.NewDecoder(resp.Body).Decode(&outputURLs)
		if err != nil {
			return URLs, err
		}
	} else if resp.StatusCode != 204 {
		return URLs, errors.New("unexpected status code: " + resp.Status)
	}

	return outputURLs, err
}
