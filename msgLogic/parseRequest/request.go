package parseRequest

import (
	"encoding/json"
	"errors"
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

	//处理request
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
		linkKeys, err = ConvDelete(request)

	case "conv-join":
		linkKeys, err = ConvJoin(request)

	case "conv-leave":
		linkKeys, err = ConvLeave(request)

	case "conv-inviteLinks":
		linkKeys, err = ConvInviteLinks(request)

	case "conv-removeLinks":
		linkKeys, err = ConvRemoveLinks(request)

	default:
		err = errors.New("error request action")
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
