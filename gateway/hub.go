package gateway

import "fmt"

var HubObj = NewHub()

//token is a user, data is send by socket
type SendData struct {
	token string
	data  []byte
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	sendcast   chan SendData
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
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
			//todo delete client & token in UserClient
			UserClientMap.Leave(client)

			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			Client_DESC()

			//receive from js
		case request := <-h.broadcast:
			fmt.Println("**read from broadcast", string(request[:]))
			//todo rpc to msgLogic, parse message

			var sendData = SendData{
				token: "001",
				data:  []byte("hi,I am server"),
			}
			for _, client := range UserClientMap.token2Client[sendData.token] {
				client.send <- sendData.data
			}
		case sendData := <-h.sendcast:
			fmt.Println("send from sendCast", sendData.token, string(sendData.data)[:])
			for _, client := range UserClientMap.token2Client[sendData.token] {
				client.send <- sendData.data
			}

			// case message := <-h.broadcast:
			// 	for client := range h.clients {
			// 		select {
			// 		case client.send <- message:
			// 		default:
			// 			close(client.send)
			// 			delete(h.clients, client)
			// 		}
			// 	}
		}
	}
}
