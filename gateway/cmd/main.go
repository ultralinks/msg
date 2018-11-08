package main

import (
	"log"
	"msg/gateway"
	"msg/gateway/app"
	gatewayService "msg/gateway/service/gateway"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin/json"
)

func main() {
	go gateway.HubObj.Run()
	go gatewayService.RunRpcServer()
	app.InitRpcClient()

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(gateway.CountVal)
		if err != nil {
			log.Println("json marshal CountVal", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		gateway.ServeWs(gateway.HubObj, w, r)
	})

	log.Println("server start")
	err := http.ListenAndServe("0.0.0.0:9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
