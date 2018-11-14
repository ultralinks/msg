package parseRequest

import (
	"encoding/json"
	"log"
	"time"

	convLinkService "msg/msgLogic/service/convLink"
	convSendService "msg/msgLogic/service/convSend"
	linkService "msg/msgLogic/service/link"
	"msg/msgLogic/service/model"
	msgService "msg/msgLogic/service/msg"
	msgConvService "msg/msgLogic/service/msgConv"
	"msg/msgLogic/util"
)

func MsgIm(request Request, requestByte []byte) {
	data, err := json.Marshal(request.Data)
	if err != nil {
		log.Println("json.Marshal data err", err)
	}

	param := request.Param
	convId := param["convId"].(string)
	fromLink, err := linkService.GetByKey(param["fromLinkKey"].(string))
	if err != nil {
		log.Println("get link by key err", err)
		return
	}

	//通过convId找到links
	links, err := getLinksByConvId(convId)
	if err != nil {
		log.Println("get links by convId err", err)
	}

	//存储msg
	err = storeMsg(links, string(data), fromLink.Id, convId)
	if err != nil {
		log.Println("storeMsg err", err)
	}

	//rpc到gateway发消息
	for _, link := range links {
		receiveSendData(link.Key, requestByte)
	}
}

func getLinksByConvId(convId string) ([]model.Link, error) {
	links, err := convLinkService.ListLink(convId)
	return *links, err
}

func storeMsg(links []model.Link, data, fromLinkId, convId string) error {
	now := time.Now()
	//msg
	msg := model.Msg{
		Id:         util.GetRandomString(11),
		Key:        util.GetRandomString(11),
		Data:       data,
		FromLinkId: fromLinkId,
		Created:    now,
		Updated:    now,
	}
	err := msgService.Create(&msg)
	if err != nil {
		return err
	}

	//msg_conv
	msgConv := model.MsgConv{
		MsgId:  msg.Id,
		ConvId: convId,
	}
	err = msgConvService.Create(&msgConv)
	if err != nil {
		return err
	}

	//conv_send
	for _, link := range links {
		err = convSendService.Create(&model.ConvSend{
			MsgId:    msg.Id,
			ConvId:   convId,
			ToLinkId: link.Id,
			Status:   0,
			Created:  now,
			Updated:  now,
		})
	}
	if err != nil {
		return err
	}

	return nil
}
