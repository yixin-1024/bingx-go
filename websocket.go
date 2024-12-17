package bingxgo

import (
	"time"

	"github.com/gorilla/websocket"
)

type WebsocketClient struct {
	// client  *Client
	// baseURL string
	conn *websocket.Conn
}

func (c *WebsocketClient) Subscribe(streams []string, handler func([]byte)) error {
	msg := struct {
		Method string   `json:"method"`
		Params []string `json:"params"`
		ID     int64    `json:"id"`
	}{
		Method: "SUBSCRIBE",
		Params: streams,
		ID:     time.Now().UnixNano(),
	}

	return c.conn.WriteJSON(msg)
}
