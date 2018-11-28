package client

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func NewClient() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: "localhost:12315", Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("websocket connect err", err)
	}
	return c
}
