package parseRequest

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Action  string                     `json:"action"`
	LinkKey string                     `json:"linkKey"`
	Param   map[string]interface{}     `json:"param"`
	Data    map[string]json.RawMessage `json:"data"`
}

func ParseRequest(requestByte []byte) ([]string, []byte, error) {
	request := Request{}
	json.Unmarshal(requestByte, &request)
	fmt.Println("request: ", request)

	var linkKeys []string
	var err error
	responseByte := requestByte

	//处理消息
	switch request.Action {
	case "msg-im":
		linkKeys, err = MsgIm(request)

	case "msg-read":
		linkKeys, err = MsgRead(request)

	case "msg-listHistory":
		linkKeys, responseByte, err = MsgListHistory(request)

	case "conv-create":
		linkKeys, err = ConvCreate(request)

	case "conv-list":
		linkKeys, responseByte, err = ConvList(request)

	case "conv-delete":

	case "conv-join":

	case "conv-leave":

	case "conv-inviteLinks":

	case "conv-removeLinks":

	}

	return linkKeys, responseByte, err
}

/*
func receiveSendData(key string, data []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := app.GatewayRpcClient.ReceiveSendData(ctx, &gateway_bak.SendDataRequest{
		Token: key,
		Data:  data,
	})
	if err != nil {
		log.Println("rpc gatewayRpcClient", err)
	}
}
*/
