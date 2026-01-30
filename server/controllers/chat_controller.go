package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pawannn/juzchat/service"
)

type ControllerRepo struct {
	chatHub *service.ChatHub
}

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func InitControllers(chatHub *service.ChatHub) *ControllerRepo {
	return &ControllerRepo{
		chatHub: chatHub,
	}
}

func (c *ControllerRepo) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := service.NewClient(conn, c.chatHub)

	c.chatHub.Register(client)

	go client.WritePump()

	defer func() {
		c.chatHub.Unregister(client)
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// rate limit
		if !client.RateLimit() {
			continue
		}

		c.chatHub.Publish(msg)
	}
}

func (c *ControllerRepo) FetchAvailableChats(w http.ResponseWriter, r *http.Request) {
	chats, err := c.chatHub.FetchAllAvailabeChats(context.Background())
	if err != nil {
		http.Error(w, "failed to fetch chats", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chats)
}
