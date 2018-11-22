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

//消息记录
type MsgHistory struct {
	ConvId string      `json:"convId"`
	Msgs   []model.Msg `json:"msgs"`
}

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

	//通过convId找到links
	links, err := getLinksByConvId(convId)
	if err != nil {
		log.Println("get links by convId err", err)
		return linkKeys, err
	}

	//存储msg
	err = storeMsg(links, string(data), r.LinkKey, convId, msgKey)
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
	toLink, _ := linkService.GetByKey(param["toLinkKey"].(string))

	//update convSend status
	whereMap := map[string]interface{}{
		"msg_id":     msgId,
		"to_link_id": toLink.Id,
	}
	updateMap := map[string]interface{}{
		"status":  1,
		"updated": time.Now(),
	}
	err := convSendService.Update(whereMap, updateMap)
	if err != nil {
		log.Println("update read status err", err)
	}

	linkKeys = append(linkKeys, toLink.Key)
	return linkKeys, err
}

func MsgListHistory(r Request) ([]string, MsgHistory, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	convId := param["convId"].(string)
	offset := int(param["offset"].(float64))
	limit := int(param["limit"].(float64))
	link, _ := linkService.GetByKey(r.LinkKey)

	//listMsg
	msgs, err := msgService.ListMsg(convId, offset, limit)
	if err != nil {
		log.Println("listMsg err", err)
	}

	msgHistory := MsgHistory{
		ConvId: convId,
		Msgs:   *msgs,
	}
	linkKeys = append(linkKeys, link.Key)
	return linkKeys, msgHistory, err
}

func getLinksByConvId(convId string) ([]model.Link, error) {
	links, err := linkService.ListLink(convId)
	return *links, err
}

func storeMsg(links []model.Link, data, fromLinkKey, convId, msgKey string) error {
	now := time.Now()
	//msg
	msg := model.Msg{
		Id:          util.GetRandomString(11),
		Key:         msgKey,
		Data:        data,
		FromLinkKey: fromLinkKey,
		Created:     now,
		Updated:     now,
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
