package rpc

import (
	"log"

	"google.golang.org/grpc"
	pbMsgLogic "msg/gateway/pb/msgLogic"
)

var MsgLogicRpcClient pbMsgLogic.MsgLogicClient

func InitRpcClient() {
	address := "127.0.0.1:10009"
	log.Println("gateway rpc client start", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // 连接
	if err != nil {
		log.Println("messageRpc error", err)
	}
	// 初始化客户端
	MsgLogicRpcClient = pbMsgLogic.NewMsgLogicClient(conn)
}
