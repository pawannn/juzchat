package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type ChatHub struct {
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	clients    map[*Client]struct{}
	redis      *redis.Client
	ctx        context.Context
}

func NewChatHub(rdb *redis.Client) *ChatHub {
	return &ChatHub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte, 1024),
		clients:    make(map[*Client]struct{}),
		redis:      rdb,
		ctx:        context.Background(),
	}
}

func (h *ChatHub) Run() {
	go h.subscribeRedis()

	for {
		select {
		case c := <-h.register:
			h.clients[c] = struct{}{}

		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				close(c.send)
			}

		case msg := <-h.broadcast:
			for c := range h.clients {
				select {
				case c.send <- msg:
				default:
					delete(h.clients, c)
					close(c.send)
				}
			}
		}
	}
}

func (h *ChatHub) Register(c *Client) {
	h.register <- c
}

func (h *ChatHub) Unregister(c *Client) {
	h.unregister <- c
}

func (h *ChatHub) Publish(msg []byte) {
	h.redis.Publish(h.ctx, chatChannel, msg)
	h.redis.LPush(h.ctx, chatKey, msg)
	h.redis.LTrim(h.ctx, chatKey, 0, maxMessages-1)
	h.redis.Expire(h.ctx, chatKey, chatTTL)
}

func (h *ChatHub) subscribeRedis() {
	sub := h.redis.Subscribe(h.ctx, chatChannel)
	ch := sub.Channel()

	for msg := range ch {
		h.broadcast <- []byte(msg.Payload)
	}
}

func (h *ChatHub) FetchAllAvailabeChats(ctx context.Context) ([]json.RawMessage, error) {
	msgs, err := h.redis.LRange(ctx, chatKey, 0, maxMessages-1).Result()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	chats := make([]json.RawMessage, len(msgs))
	for i, m := range msgs {
		chats[len(msgs)-1-i] = json.RawMessage(m)
	}

	return chats, nil
}
