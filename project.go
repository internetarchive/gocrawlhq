package gocrawlhq

import (
	"context"
	"encoding/json"
	"net/http"
)

type Project struct {
	Paused           bool     `json:"paused"`
	Name             string   `json:"name"`
	Exclusions       []string `json:"exclusions"`
	SeencheckEnabled bool     `json:"seencheck_enabled"`
	SeencheckTTL     int      `json:"seencheck_ttl"`
	Stats            struct {
		Pending         int `json:"pending"`
		Processing      int `json:"processing"`
		CompletedSeeds  int `json:"completed_seeds"`
		CompletedAssets int `json:"completed_assets"`
	} `json:"stats"`
}

func (c *Client) GetProject(ctx context.Context) (p *Project, err error) {
	expectedStatusCodes := 200

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.ProjectEndpoint.String(), nil)
	if err != nil {
		return p, err
	}

	req.Header.Add("X-Auth-Key", c.Key)
	req.Header.Add("X-Auth-Secret", c.Secret)
	req.Header.Add("User-Agent", "gocrawlhq/"+Version)

	if c.Identifier != "" {
		req.Header.Add("X-Identifier", c.Identifier)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return p, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatusCodes {
		return p, err
	}

	err = json.NewDecoder(resp.Body).Decode(&p)
	return p, err
}
