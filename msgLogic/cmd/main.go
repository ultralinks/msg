package main

import (
	"time"

	"msg/msgLogic/rpc"

	"msg/msgLogic/app"
	"msg/msgLogic/httpServer"
)

func main() {
	app.InitConfig()
	app.InitDB()
	go rpc.InitRpcClient()
	go rpc.RunRpcServer()

	go httpServer.Start()

	for {
		time.Sleep(time.Minute)
	}
}
