package cmd

import (
	"msg/msgLogic/pb/gateway"
	"msg/msgLogic/rpcClient"
)

var GatewayRpcClient gateway.GatewayClient

func main(){
	rpcClient.InitGatewayRpcClient()
}

