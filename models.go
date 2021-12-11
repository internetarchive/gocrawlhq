package gocrawlhq

import "net/http"

type Client struct {
	Key        string
	Secret     string
	Project    string
	HQAddress  string
	HTTPClient *http.Client
}

type URL struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type FeedResponse struct {
	URLs []URL `json:"urls"`
}

type DiscoveredResponse struct {
	Project string `json:"project"`
	Type    string `json:"type"`
}

type FinishedResponse struct {
	Project string `json:"project"`
}

type DiscoveredPayload struct {
	Project string   `json:"project"`
	Type    string   `json:"type"`
	URLs    []string `json:"urls"`
}

type FinishedPayload struct {
	Project string `json:"project"`
	URLs    []URL  `json:"urls"`
}
