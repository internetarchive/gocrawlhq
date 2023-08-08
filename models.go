package gocrawlhq

import (
	"net"
	"net/http"
)

type Client struct {
	Key           string
	Secret        string
	Project       string
	HQAddress     string
	Identifier    string
	HTTPClient    *http.Client
	WebsocketConn *net.Conn
}

type URL struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value"`
	Path  string `json:"path"`
	Via   string `json:"via,omitempty"`
}

type FeedResponse struct {
	Project string `json:"project"`
	URLs    []URL  `json:"urls"`
}

type DiscoveredResponse struct {
	Project string `json:"project"`
	Type    string `json:"type"`
	URLs    []URL  `json:"urls,omitempty"`
}

type FinishedResponse struct {
	Project string `json:"project"`
}

type DiscoveredPayload struct {
	Type            string `json:"type"`
	URLs            []URL  `json:"urls"`
	BypassSeencheck bool   `json:"bypassSeencheck"`
	SeencheckOnly   bool   `json:"seencheckOnly"`
}

type FinishedPayload struct {
	LocalCrawls int   `json:"localCrawls"`
	URLs        []URL `json:"urls"`
}
