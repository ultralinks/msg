package gateway

import (
	"fmt"

	"msg/gateway/rpc"
)

var HubObj = NewHub()

//token is a user, data is send by socket
type SendData struct {
	Key  string
	Data []byte
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	Sendcast   chan *SendData
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		Sendcast:   make(chan *SendData, 10000),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			Client_INCR()
		case client := <-h.unregister:
			LinkDeviceMap.Leave(client)

			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			Client_DESC()
		case request := <-h.broadcast:
			fmt.Println("**read from broadcast", string(request[:]))
			linkKeys, data := rpc.ParseMsg(request)

			for _, linkKey := range linkKeys {
				sendData := &SendData{
					Key:  linkKey,
					Data: data,
				}
				h.Sendcast <- sendData
			}
		case sendData := <-h.Sendcast:
			fmt.Println("send from sendCast", sendData.Key, string(sendData.Data)[:])
			for _, client := range LinkDeviceMap.linkKey2Device[sendData.Key] {
				client.send <- sendData.Data
			}
		}
	}
}
