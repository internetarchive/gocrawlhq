package gocrawlhq

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type IdentifyMessage struct {
	Project    string `json:"project"`
	Job        string `json:"job"`
	IP         string `json:"ip"`
	Hostname   string `json:"hostname"`
	Identifier string `json:"identifier"`
	Timestamp  int64  `json:"timestamp"`
	GoVersion  string `json:"goVersion"`
}

func (c *Client) InitWebsocketConn() (err error) {
	// Remove http or https from the HQAddress
	HQAddress := strings.ReplaceAll(c.HQAddress, "http://", "")
	HQAddress = strings.ReplaceAll(HQAddress, "https://", "")

	// Initialize the websocket connection
	headers := make(http.Header)
	headers.Add("X-Auth-Key", c.Key)
	headers.Add("X-Auth-Secret", c.Secret)
	headers.Add("User-Agent", "gocrawlhq/"+Version)

	if c.Identifier != "" {
		headers.Add("X-Identifier", c.Identifier)
	}

	dialer := &ws.Dialer{
		Header: ws.HandshakeHeaderHTTP(headers),
	}

	conn, _, _, err := dialer.Dial(context.Background(), "ws://"+path.Join(HQAddress, "api", "ws"))
	if err != nil {
		return err
	}

	c.WebsocketConn = &conn

	return nil
}

func (c *Client) Identify(msg *IdentifyMessage) (err error) {
	msg.Identifier = c.Identifier
	msg.Timestamp = time.Now().UTC().Unix()

	marshalled, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// Send the identify message
	err = wsutil.WriteClientMessage(*c.WebsocketConn, ws.OpText, []byte(`{"type":"identify","payload":`+string(marshalled)+`}`))
	if err != nil {
		return err
	}

	return nil
}
