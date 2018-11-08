package main

import (
	"msg/gateway/service/msgLogic"
	msgLogic2 "msg/msgLogic"
	"msg/msgLogic/app"
	"msg/msgLogic/httpServer"
	"msg/msgLogic/pb/gateway"
	"time"
)

var GatewayRpcClient gateway.GatewayClient

func main(){
	app.InitConfig()
	app.InitDB()
	app.InitRpcClient()
	go msgLogic2.RunRpcServer()

	go msgLogic.InitRpcClient()
	go httpServer.Start()

	for {
		time.Sleep(time.Minute)
	}
}
