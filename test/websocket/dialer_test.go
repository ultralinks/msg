package websocket

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gorilla/websocket"
)

// https://stackoverflow.com/questions/47637308/create-unit-test-for-ws-in-golang

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    defer c.Close()
    for {
        mt, message, err := c.ReadMessage()
        if err != nil {
            break
        }

        responseMsg := string(message) + " back"
        err = c.WriteMessage(mt, []byte(responseMsg))
        if err != nil {
            break
        }
    }
}

func TestDialer(t *testing.T) {
    s := httptest.NewServer(http.HandlerFunc(echo))
    defer s.Close()

    u := "ws" + strings.TrimPrefix(s.URL, "http")
    ws, _, err := websocket.DefaultDialer.Dial(u, nil)
    if err != nil {
        t.Errorf("%v", err)
    }
    defer ws.Close()

    for i := 0; i < 10; i++ {
        if err := ws.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
            t.Fatalf("%v", err)
        }

        _, p, err := ws.ReadMessage()
        if err != nil {
            t.Fatalf("%v", err)
        }
        if  string(p) != "hello back" {
            t.Fatalf("bad message")
        }
    }
}
