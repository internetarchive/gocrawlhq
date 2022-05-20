package gocrawlhq

import (
	"net/http"
	"net/url"
	"os"
	"path"
)

var (
	DiscoveredEndpoint *url.URL
	FinishedEndpoint   *url.URL
	FeedEndpoint       *url.URL

	Version = "1.1.4"
)

func Init(key, secret, project, HQAddress string) (c *Client, err error) {
	c = new(Client)

	c.Key = key
	c.Secret = secret
	c.Project = project
	c.HTTPClient = http.DefaultClient
	c.HQAddress = HQAddress

	hostname, err := os.Hostname()
	if err != nil {
		return c, err
	}

	c.Identifier = hostname + "-" + project

	DiscoveredEndpoint, err = url.Parse(c.HQAddress)
	if err != nil {
		return c, err
	}

	FinishedEndpoint, err = url.Parse(c.HQAddress)
	if err != nil {
		return c, err
	}

	FeedEndpoint, err = url.Parse(c.HQAddress)
	if err != nil {
		return c, err
	}

	DiscoveredEndpoint.Path = path.Join(DiscoveredEndpoint.Path, "api", "discovered")
	FinishedEndpoint.Path = path.Join(FinishedEndpoint.Path, "api", "finished")
	FeedEndpoint.Path = path.Join(FeedEndpoint.Path, "api", "feed", c.Project)

	return c, nil
}
