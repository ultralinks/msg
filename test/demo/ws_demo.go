package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:9000", "http service address")

func loop() {
	for {
		u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
		q := u.Query()
		q.Set("authToken", "001")
		u.RawQuery = q.Encode()
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			continue
		}
		// 循环读写消息
		for {
			error := c.WriteMessage(websocket.TextMessage, []byte("hello, I am client"))
			if error != nil {
				fmt.Println("writeMessage error")
			}

			time.Sleep(time.Second)
			_, message, err := c.ReadMessage()
			if err != nil {
				break
			}
			fmt.Println("read message:", string(message[:]))
			time.Sleep(3 * time.Second)
			// log.Printf("recv: %s", message)
		}
		c.Close()
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	for i := 0; i < 1000; i++ {
		go loop()
	}

	for {
		time.Sleep(1 * time.Second)
	}
}
