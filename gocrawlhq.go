package gocrawlhq

import (
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

var (
	Version = "1.2.15"
)

func Init(key, secret, project, address, identifier string, timeout int) (c *Client, err error) {
	c = new(Client)

	if timeout == 0 {
		timeout = 5
	}

	c.Key = key
	c.Secret = secret
	c.Project = project
	c.HQAddress = address
	c.HTTPClient = &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	if identifier == "" {
		// Initialize the identifier
		hostname, err := os.Hostname()
		if err != nil {
			return c, err
		}

		c.Identifier = hostname + "-" + project
	} else {
		c.Identifier = identifier
	}

	// Initialize the websocket connection
	err = c.InitWebsocketConn()
	if err != nil {
		return c, err
	}

	// Initialize the endpoints
	c.URLsEndpoint, err = url.Parse(c.HQAddress)
	if err != nil {
		return c, err
	}

	c.SeencheckEndpoint, err = url.Parse(c.HQAddress)
	if err != nil {
		return c, err
	}

	c.ResetEndpoint, err = url.Parse(c.HQAddress)
	if err != nil {
		return c, err
	}

	c.ProjectEndpoint, err = url.Parse(c.HQAddress)
	if err != nil {
		return c, err
	}

	c.URLsEndpoint.Path = path.Join(c.URLsEndpoint.Path, "api", "projects", c.Project, "urls")
	c.SeencheckEndpoint.Path = path.Join(c.SeencheckEndpoint.Path, "api", "projects", c.Project, "seencheck")
	c.ResetEndpoint.Path = path.Join(c.ResetEndpoint.Path, "api", "projects", c.Project, "reset")
	c.ProjectEndpoint.Path = path.Join(c.ProjectEndpoint.Path, "api", "projects", c.Project)

	return c, nil
}
