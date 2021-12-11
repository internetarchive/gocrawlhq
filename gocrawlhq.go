package gocrawlhq

import (
	"net/http"
	"path"
)

var (
	discoveredEndpoint string
	finishedEndpoint   string
	feedEndpoint       string

	Version = "1.0.0"
)

func Init(key, secret, project, HQAddress string) (c Client, err error) {
	c.Key = key
	c.Secret = secret
	c.Project = project
	c.HTTPClient = http.DefaultClient
	c.HQAddress = HQAddress

	discoveredEndpoint = path.Join(c.HQAddress + "discovered")
	finishedEndpoint = path.Join(c.HQAddress + "finished")
	feedEndpoint = path.Join(c.HQAddress + "feed")

	return c, nil
}
