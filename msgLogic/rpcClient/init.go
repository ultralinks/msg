package rpcClient

import (
"google.golang.org/grpc"
"log"
"msg/msgLogic/pb/gateway"
)

var GatewayRpcClient gateway.GatewayClient

func InitGatewayRpcClient() {
	address := "127.0.0.1:10000"
	log.Println("gateway rpc client start", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // 连接
	if err != nil {
		log.Println("messageRpc error", err)
	}
	// 初始化客户端
	GatewayRpcClient = gateway.NewGatewayClient(conn)
}