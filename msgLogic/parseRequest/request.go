package parseRequest

import (
	"context"
	"encoding/json"
	"log"
	"msg/msgLogic/app"
	"msg/msgLogic/pb/gateway"
	"time"
)

type Message struct {
	ConvId string
	from   string
	//msgType        string // text image video audio

	//text     string
	//image    string
	//video    string
	//audio    string
	//location string
	data string
}

type Request struct {
	Action string                     `json:"action"`
	Data   map[string]json.RawMessage `json:"data"`
	ConvId string                     `json:"convId"`
	From   string                     `json:"from"`
}

//type Response struct {
//	action  string
//	content Message
//	headers struct{}
//}

func ParseRequest(requestByte []byte) {
	request := Request{}
	json.Unmarshal(requestByte, &request)

	switch request.Action {
	// data.action==im, get ConvId and get to users token, then put response to gateway
	case "im":
		ConvId := request.ConvId
		linkKeys := getLinkKeyByConvId(ConvId)

		data := requestByte
		for _, token := range linkKeys {
			receiveSendData(token, data)
		}

	case "listHistory":

	case "listConv":

	}
}

func receiveSendData(token string, data []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//todo data 加工
	_, err := app.GatewayRpcClient.ReceiveSendData(ctx, &gateway.SendDataRequest{
		Token: token,
		Data:  data,
	})

	if err != nil {
		log.Println("rpc gatewayRpcClient", err)
	}

}

//mock param ConvId return linkKeys
func getLinkKeyByConvId(ConvId string) []string {
	if ConvId == "11111" {
		return []string{"001", "002"}
	}
	return []string{}
}
