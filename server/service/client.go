package service

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait   = 5 * time.Second
	sendBuf     = 256
	maxMessages = 200
	chatChannel = "juzchat:global"
	chatKey     = "juzchat:messages"
	chatTTL     = 3 * time.Hour
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
	rl   *RateLimiter
	hub  *ChatHub
}

func NewClient(conn *websocket.Conn, hub *ChatHub) *Client {
	return &Client{
		conn: conn,
		send: make(chan []byte, sendBuf),
		rl:   NewRateLimiter(20), // 20 msgs/sec
		hub:  hub,
	}
}

func (c *Client) WritePump() {
	for msg := range c.send {
		c.conn.SetWriteDeadline(time.Now().Add(writeWait))
		if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.conn.Close()
}

func (c *Client) RateLimit() bool {
	return c.rl.Allow()
}
