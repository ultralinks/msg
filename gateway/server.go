package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin/json"
)

func main() {
	hub := newHub()
	go hub.run()
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(CountVal)
		if err != nil {
			log.Println("json marshal CountVal", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	log.Println("server start")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
