package service

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebSocket(hub *ChatHub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := NewClient(conn, hub)
	hub.register <- client

	go client.WritePump()

	defer func() {
		hub.unregister <- client
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		if !client.rl.Allow() {
			continue // rate limit
		}

		hub.Publish(msg)
	}
}
