package parseRequest

import (
	"context"
	"encoding/json"
	"log"
	"msg/msgLogic/pb/gateway"
	"msg/msgLogic/rpcClient"
	"time"
)

type Message struct {
	conversationId string
	from           string
	msgType        string // text image video audio

	text     string
	image    string
	video    string
	audio    string
	location string
}

type Request struct {
	action  string
	content Message
	headers struct{}
}

type Response struct {
	action  string
	content Message
	headers struct{}
}

func ParseRequest(requestByte []byte) {
	request := Request{}
	json.Unmarshal(requestByte, &request)

	switch request.action {
	// data.action==im, get conversationId and get to users token, then put response to gateway
	case "im":
		conversationId := request.content.conversationId
		tokens := getUserTokensByConversationId(conversationId)

		data := requestByte
		for _, token := range tokens {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			//todo data 加工

			status, err := rpcClient.GatewayRpcClient.ReceiveSendData(ctx, &gateway.SendDataRequest{
				Token: token,
				Data:  data,
			})

			if err != nil {
				log.Println("rpc get user info error", err)
			}
			log.Println("msgLogic to gateway status", status)
		}

	case "listHistory":

	case "listConv":

	}
}

//mock param conversationId return tokens
func getUserTokensByConversationId(conversationId string) []string {
	return []string{"001", "002"}
}
