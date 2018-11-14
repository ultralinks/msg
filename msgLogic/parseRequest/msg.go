package parseRequest

import (
	"encoding/json"
	"log"
	"time"

	convSendService "msg/msgLogic/service/convSend"
	linkService "msg/msgLogic/service/link"
	"msg/msgLogic/service/model"
	msgService "msg/msgLogic/service/msg"
	msgConvService "msg/msgLogic/service/msgConv"
	"msg/msgLogic/util"
)

func MsgIm(r Request, requestByte []byte) {
	data, err := json.Marshal(r.Data)
	if err != nil {
		log.Println("json.Marshal data err", err)
	}

	param := r.Param
	convId := param["convId"].(string)
	msgKey := param["msgKey"].(string)
	fromLink, err := linkService.GetByKey(r.LinkKey)
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
	err = storeMsg(links, string(data), fromLink.Id, convId, msgKey)
	if err != nil {
		log.Println("storeMsg err", err)
	}

	//rpc到gateway发消息
	for _, link := range links {
		receiveSendData(link.Key, requestByte)
	}
}

func MsgRead(r Request, requestByte []byte) {
	param := r.Param
	msgId := param["msgId"].(string)
	toLink, err := linkService.GetByKey(param["toLinkKey"].(string))
	if err != nil {
		log.Println("get link by key err", err)
		return
	}

	whereMap := map[string]interface{}{
		"msg_id":     msgId,
		"to_link_id": toLink.Id,
	}
	updateMap := map[string]interface{}{
		"status":  1,
		"updated": time.Now(),
	}
	if err := convSendService.Update(whereMap, updateMap); err != nil {
		log.Println("update read status err", err)
	}

	receiveSendData(toLink.Key, requestByte)
}

func MsgListHistory(r Request) {
	param := r.Param
	convId := param["convId"].(string)
	link, err := linkService.GetByKey(r.LinkKey)
	if err != nil {
		log.Println("get link by key err", err)
		return
	}

	msgs, err := msgService.ListMsg(convId)
	if err != nil {
		log.Println("listMsg err", err)
	}
	msgsByte, err := json.Marshal(msgs)
	if err != nil {
		log.Println("msgsByte err", err)
	}

	receiveSendData(link.Key, msgsByte)
}

func getLinksByConvId(convId string) ([]model.Link, error) {
	links, err := linkService.ListLink(convId)
	return *links, err
}

func storeMsg(links []model.Link, data, fromLinkId, convId, msgKey string) error {
	now := time.Now()
	//msg
	msg := model.Msg{
		Id:         util.GetRandomString(11),
		Key:        msgKey,
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
