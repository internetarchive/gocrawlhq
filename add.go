package gocrawlhq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Add(ctx context.Context, URLs []URL, bypassSeencheck bool) (err error) {
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
	req, err := NewAPIRequest(c, ctx, http.MethodPost, c.URLsEndpoint.String(), bytes.NewReader(jsonPayload))
	if err != nil {
		return err
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
