package parseRequest

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"msg/msgLogic/app"
	"msg/msgLogic/pb/gateway"
	convLinkService "msg/msgLogic/service/convLink"
	linkService "msg/msgLogic/service/link"
	"msg/msgLogic/service/model"
	msgService "msg/msgLogic/service/msg"
	msgConvService "msg/msgLogic/service/msgConv"
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
	ConvId int                        `json:"convId"`
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
		fromLink, err := linkService.GetByKey(request.From)
		if err != nil {
			log.Println("get link by key err", err)
			return
		}

		links, err := getLinksByConvId(request.ConvId)
		if err != nil {
			log.Println("get links by convId err", err)
			return
		}

		err = storeMsg(request, fromLink.Id)
		if err != nil {
			log.Println("storeMsg err", err)
		}

		data := requestByte
		for _, link := range links {
			receiveSendData(link.Key, data)
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

//从 conv 表、和 link 表找出来，该convId 下有哪些 linkKey
func getLinksByConvId(convId int) ([]model.Link, error) {
	links, err := convLinkService.ListLink(convId)
	return *links, err
}

//存储消息: msg msg_conv conv_link
func storeMsg(r Request, fromLinkId int) error {
	data, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	now := time.Now()
	msg := &model.Msg{
		Data:       string(data),
		FromLinkId: fromLinkId,
		Created:    now,
		Updated:    now,
	}
	err = msgService.Create(msg)
	if err != nil {
		return err
	}

	msgConv := model.MsgConv{
		MsgId:  msg.Id,
		ConvId: r.ConvId,
	}
	err = msgConvService.Create(&msgConv)
	if err != nil {
		return err
	}

	return nil
}
