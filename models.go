package gocrawlhq

import (
	"net"
	"net/http"
	"net/url"
)

type Client struct {
	Key               string
	Secret            string
	Project           string
	HQAddress         string
	Identifier        string
	URLsEndpoint      *url.URL
	SeencheckEndpoint *url.URL
	ResetEndpoint     *url.URL
	HTTPClient        *http.Client
	WebsocketConn     *net.Conn
}

type URL struct {
	ID        string `json:"id" db:"id"`
	Value     string `json:"value" db:"value"`
	Via       string `json:"via,omitempty" db:"via"`
	Host      string `json:"host,omitempty" db:"host"`
	Path      string `json:"path,omitempty" db:"path"`
	Type      string `json:"type,omitempty" db:"type"`
	Crawler   string `json:"crawler,omitempty" db:"crawler"`
	Status    string `json:"status" db:"status"`
	LiftOff   int64  `json:"lift_off" db:"lift_off"`
	Timestamp int64  `json:"timestamp" db:"timestamp"`
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
	URLs            []URL `json:"urls"`
	BypassSeencheck bool  `json:"bypassSeencheck"`
}

type FinishedPayload struct {
	LocalCrawls int   `json:"localCrawls"`
	URLs        []URL `json:"urls"`
}
