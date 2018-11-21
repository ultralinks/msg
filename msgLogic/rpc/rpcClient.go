package rpc

import (
	"log"

	"google.golang.org/grpc"
	userPb "msg/msgLogic/pb/userLogic"
)

var UserRpcClient userPb.UserClient

func InitRpcClient() {
	appClient()
}

func appClient() {
	address := "127.0.0.1:10010"
	log.Println("message rpc client start", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // 连接
	if err != nil {
		log.Println("messageRpc error", err)
	}
	// 初始化客户端
	UserRpcClient = userPb.NewUserClient(conn)
}
