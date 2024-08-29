package gocrawlhq

import (
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

var (
	Version = "1.2.6"
)

func Init(key, secret, project, HQAddress string) (c *Client, err error) {
	c = new(Client)

	// Initialize the identifier
	hostname, err := os.Hostname()
	if err != nil {
		return c, err
	}

	c.Key = key
	c.Secret = secret
	c.Project = project

	c.HTTPClient = &http.Client{
		Timeout: time.Second * 5,
	}

	c.HQAddress = HQAddress
	c.Identifier = hostname + "-" + project

	// Initialize the websocket connection
	err = c.InitWebsocketConn()
	if err != nil {
		return c, err
	}

	// Initialize the endpoints
	c.DiscoveredEndpoint, err = url.Parse(c.HQAddress)
	if err != nil {
		return c, err
	}

	c.FinishedEndpoint, err = url.Parse(c.HQAddress)
	if err != nil {
		return c, err
	}

	c.FeedEndpoint, err = url.Parse(c.HQAddress)
	if err != nil {
		return c, err
	}

	c.DiscoveredEndpoint.Path = path.Join(c.DiscoveredEndpoint.Path, "api", "project", c.Project, "discovered")
	c.FinishedEndpoint.Path = path.Join(c.FinishedEndpoint.Path, "api", "project", c.Project, "finished")
	c.FeedEndpoint.Path = path.Join(c.FeedEndpoint.Path, "api", "project", c.Project, "feed")

	return c, nil
}
