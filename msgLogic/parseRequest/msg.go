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

func MsgIm(r Request) ([]string, error) {
	linkKeys := make([]string, 0)

	data, err := json.Marshal(r.Data)
	if err != nil {
		log.Println("json.Marshal data err", err)
		return linkKeys, err
	}

	param := r.Param
	convId := param["convId"].(string)
	msgKey := param["msgKey"].(string)
	fromLink, err := linkService.GetByKey(r.LinkKey)
	if err != nil {
		log.Println("get link by key err", err)
		return linkKeys, err
	}

	//通过convId找到links
	links, err := getLinksByConvId(convId)
	if err != nil {
		log.Println("get links by convId err", err)
		return linkKeys, err
	}

	//存储msg
	err = storeMsg(links, string(data), fromLink.Id, convId, msgKey)
	if err != nil {
		log.Println("storeMsg err", err)
		return linkKeys, err
	}

	for _, link := range links {
		linkKeys = append(linkKeys, link.Key)
	}
	return linkKeys, nil
}

func MsgRead(r Request) ([]string, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	msgId := param["msgId"].(string)
	toLink, err := linkService.GetByKey(param["toLinkKey"].(string))
	if err != nil {
		log.Println("get link by key err", err)
		return linkKeys, err
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

	linkKeys = append(linkKeys, toLink.Id)
	return linkKeys, nil
}

func MsgListHistory(r Request) ([]string, []byte, error) {
	linkKeys := make([]string, 0)
	responseByte := make([]byte, 0)

	param := r.Param
	convId := param["convId"].(string)
	offset := param["offset"].(int)
	limit := param["limit"].(int)
	link, err := linkService.GetByKey(r.LinkKey)
	if err != nil {
		log.Println("get link by key err", err)
		return linkKeys, responseByte, err
	}

	msgs, err := msgService.ListMsg(convId, offset, limit)
	if err != nil {
		log.Println("listMsg err", err)
	}
	responseByte, err = json.Marshal(msgs)
	if err != nil {
		log.Println("responseByte err", err)
	}

	linkKeys = append(linkKeys, link.Id)
	return linkKeys, responseByte, nil
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
