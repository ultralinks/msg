package parseRequest

import (
	"encoding/json"
	"errors"
	"log"
)

//ws request
type Request struct {
	Action  string                     `json:"action"`
	LinkKey string                     `json:"linkKey"`
	Param   map[string]interface{}     `json:"param"`
	Data    map[string]json.RawMessage `json:"data"`
}

//ws response
type Response struct {
	Action string      `json:"action"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

//解析websocket请求
func ParseRequest(requestByte []byte) ([]string, []byte, error) {
	request := Request{}
	json.Unmarshal(requestByte, &request)
	log.Println("request: ", request)

	var linkKeys []string
	var err error
	response := &Response{
		Action: request.Action,
		Status: "ok",
		Data:   request,
	}

	//handle ws request
	switch request.Action {
	case "msg-im":
		linkKeys, err = MsgIm(request)

	case "msg-read":
		linkKeys, err = MsgRead(request)

	case "msg-listHistory":
		linkKeys, response.Data, err = MsgListHistory(request)

	case "conv-create":
		linkKeys, response.Data, err = ConvCreate(request)

	case "conv-list":
		linkKeys, response.Data, err = ConvList(request)

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

	responseByte, _ := json.Marshal(*response)
	return linkKeys, responseByte, err
}
