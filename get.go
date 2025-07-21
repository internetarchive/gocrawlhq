package gocrawlhq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var ErrFeedEmpty = errors.New("gocrawlhq: feed is empty")

func (c *Client) Get(ctx context.Context, size int) (URLs []URL, err error) {
	// build request
	req, err := NewAPIRequest(c, ctx, http.MethodGet, c.URLsEndpoint.String(), nil)
	if err != nil {
		return URLs, err
	}

	q := req.URL.Query()
	q.Add("size", strconv.Itoa(size))
	req.URL.RawQuery = q.Encode()

	// execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return URLs, err
	}
	defer resp.Body.Close()

	// check response status code for 'empty' or 204
	if resp.StatusCode == http.StatusNoContent {
		return URLs, nil
	}

	// check response status code for 200
	if resp.StatusCode != http.StatusOK {
		return URLs, fmt.Errorf("%w: %d", ErrUnexpectedStatusCode, resp.StatusCode)
	}

	// decode response body
	err = json.NewDecoder(resp.Body).Decode(&URLs)
	if err != nil {
		return URLs, err
	}

	return URLs, err
}
