package gocrawlhq

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) Reset(ctx context.Context) (err error) {
	expectedStatusCode := 202

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.ResetEndpoint.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-Auth-Key", c.Key)
	req.Header.Add("X-Auth-Secret", c.Secret)
	req.Header.Add("User-Agent", "gocrawlhq/"+Version)

	if c.Identifier != "" {
		req.Header.Add("X-Identifier", c.Identifier)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatusCode {
		return fmt.Errorf("non-%d status code: %d", expectedStatusCode, resp.StatusCode)
	}

	return nil
}

func (c *Client) ResetURL(ctx context.Context, ID string) (err error) {
	expectedStatusCode := 200

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.ResetEndpoint.String()+"/"+ID, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-Auth-Key", c.Key)
	req.Header.Add("X-Auth-Secret", c.Secret)
	req.Header.Add("User-Agent", "gocrawlhq/"+Version)

	if c.Identifier != "" {
		req.Header.Add("X-Identifier", c.Identifier)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatusCode {
		return fmt.Errorf("non-%d status code: %d", expectedStatusCode, resp.StatusCode)
	}

	return nil
}
