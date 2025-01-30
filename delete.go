package gocrawlhq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Delete(ctx context.Context, URLs []URL, localCrawls int) (err error) {
	expectedStatusCode := 204

	// build payload
	payload := DeletePayload{
		LocalCrawls: localCrawls,
		URLs:        URLs,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// build request
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.URLsEndpoint.String(), bytes.NewReader(jsonPayload))
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
