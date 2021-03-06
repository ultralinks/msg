package client

import (
	"net/url"
	"github.com/gorilla/websocket"
	"log"
)

func NewClient() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: "localhost:9000", Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("websocket connect err", err)
	}
	return c
}
