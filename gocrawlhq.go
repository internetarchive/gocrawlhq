package gocrawlhq

import (
	"net/http"
	"net/url"
	"path"
)

var (
	discoveredEndpoint *url.URL
	finishedEndpoint   *url.URL
	feedEndpoint       *url.URL

	Version = "1.0.0"
)

func Init(key, secret, project, HQAddress string) (c *Client, err error) {
	c = new(Client)

	c.Key = key
	c.Secret = secret
	c.Project = project
	c.HTTPClient = http.DefaultClient
	c.HQAddress = HQAddress

	discoveredEndpoint, err := url.Parse(c.HQAddress)
	if err != nil {
		return c, err
	}

	finishedEndpoint, err := url.Parse(c.HQAddress)
	if err != nil {
		return c, err
	}

	feedEndpoint, err := url.Parse(c.HQAddress)
	if err != nil {
		return c, err
	}

	discoveredEndpoint.Path = path.Join(discoveredEndpoint.Path, "discovered")
	finishedEndpoint.Path = path.Join(finishedEndpoint.Path, "finished")
	feedEndpoint.Path = path.Join(feedEndpoint.Path, "feed")

	return c, nil
}
