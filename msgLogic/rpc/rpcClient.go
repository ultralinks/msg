package rpc

import (
	"log"

	"google.golang.org/grpc"
	"msg/msgLogic/app"
	userPb "msg/msgLogic/pb/userLogic"
)

var UserRpcClient userPb.UserClient

func InitRpcClient() {
	appClient()
}

func appClient() {
	address := app.Config.UserRpc.Host + ":" + app.Config.UserRpc.Port
	log.Println("message rpc client start", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // 连接
	if err != nil {
		log.Println("messageRpc error", err)
	}
	// 初始化客户端
	UserRpcClient = userPb.NewUserClient(conn)
}
