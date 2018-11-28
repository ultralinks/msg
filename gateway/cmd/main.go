package main

import (
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof"

	"msg/gateway"
	"msg/gateway/rpc"
)

func main() {
	go gateway.HubObj.Run()
	go rpc.InitRpcClient()

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(gateway.CountVal)
		if err != nil {
			log.Println("json marshal CountVal", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	})

	//localhost/ws?token=001
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		gateway.ServeWs(gateway.HubObj, w, r)
	})

	log.Println("server start")
	err := http.ListenAndServe("0.0.0.0:12315", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
