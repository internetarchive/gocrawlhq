package gocrawlhq

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) Reset(ctx context.Context) (err error) {
	req, err := NewAPIRequest(c, ctx, http.MethodPost, c.ResetEndpoint.String(), nil)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("%w: %d", ErrUnexpectedStatusCode, resp.StatusCode)
	}

	return nil
}

func (c *Client) ResetURL(ctx context.Context, ID string) (err error) {
	req, err := NewAPIRequest(c, ctx, http.MethodPost, c.ResetEndpoint.String()+"/"+ID, nil)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %d", ErrUnexpectedStatusCode, resp.StatusCode)
	}

	return nil
}
