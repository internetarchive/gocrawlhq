package gocrawlhq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Seencheck(ctx context.Context, URLs []URL) (outputURLs []URL, err error) {
	jsonPayload, err := json.Marshal(URLs)
	if err != nil {
		return URLs, err
	}

	endpoint := c.SeencheckEndpoint.String()
	if c.AltSeencheckURL != "" {
		endpoint = c.AltSeencheckURL
	}

	req, err := NewAPIRequest(c, ctx, http.MethodPost, endpoint, bytes.NewReader(jsonPayload))
	if err != nil {
		return URLs, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return URLs, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&outputURLs)
		if err != nil {
			return URLs, err
		}
	} else if resp.StatusCode != http.StatusNoContent {
		return URLs, fmt.Errorf("%w: %d", ErrUnexpectedStatusCode, resp.StatusCode)
	}

	return outputURLs, err
}
