package parseRequest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"msg/msgLogic/app"
	"msg/msgLogic/pb/gateway"
)

type Request struct {
	Action  string                     `json:"action"`
	LinkKey string                     `json:"linkKey"`
	Param   map[string]interface{}     `json:"param"`
	Data    map[string]json.RawMessage `json:"data"`
}

func ParseRequest(requestByte []byte) {
	request := Request{}
	json.Unmarshal(requestByte, &request)
	fmt.Println("request: ", request)

	//处理消息
	switch request.Action {
	case "msg-im":
		MsgIm(request, requestByte)
	case "msg-read":
		MsgRead(request, requestByte)

	case "msg-listHistory":
		MsgListHistory(request)

	case "conv-create":
		ConvCreate(request, requestByte)

	case "conv-list":
		ConvList(request)

	case "conv-delete":

	case "conv-join":

	case "conv-leave":

	case "conv-inviteLinks":

	case "conv-removeLinks":

	}
}

func receiveSendData(key string, data []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//todo data 加工
	_, err := app.GatewayRpcClient.ReceiveSendData(ctx, &gateway_bak.SendDataRequest{
		Token: key,
		Data:  data,
	})

	if err != nil {
		log.Println("rpc gatewayRpcClient", err)
	}

}
