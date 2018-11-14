package main

import (
	//"msg/gateway/service/msgLogic"
	msgLogic2 "msg/msgLogic"
	"msg/msgLogic/app"
	"msg/msgLogic/httpServer"
	"time"
)

func main() {
	app.InitConfig()
	app.InitDB()
	go msgLogic2.RunRpcServer()
	go httpServer.Start()

	for {
		time.Sleep(time.Minute)
	}
}
