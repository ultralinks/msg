package parseRequest

import (
	"log"
	"time"

	convSendService "msg/msgLogic/service/convSend"
	linkService "msg/msgLogic/service/link"
	"msg/msgLogic/service/model"
	msgService "msg/msgLogic/service/msg"
	msgConvService "msg/msgLogic/service/msgConv"
	"msg/msgLogic/util"
)

type MsgItemResp struct {
	ConvId      string    `json:"convId"`
	Id          string    `json:"id"`
	Key         string    `json:"key"`
	Type        string    `json:"type"`
	Content     string    `json:"content"`
	FromLinkKey string    `json:"fromLinkKey"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

type MsgHistoryResp struct {
	ConvId string        `json:"convId"`
	Msgs   []MsgItemResp `json:"msgs"`
}

func MsgIm(r Request) ([]string, MsgItemResp, error) {
	linkKeys := make([]string, 0)
	var msgItemResp MsgItemResp

	param := r.Param
	convId, ok := param["convId"].(string)
	if !ok {
		convId = ""
	}
	msgKey, ok := param["msgKey"].(string)
	if !ok {
		msgKey = ""
	}
	msgType, ok := param["msgType"].(string)
	if !ok {
		msgType = "text"
	}
	msgContent, ok := param["msgContent"].(string)
	if !ok {
		msgContent = ""
	}

	//通过convId找到links
	links, err := getLinksByConvId(convId)
	if err != nil {
		log.Println("get links by convId err", err)
		return linkKeys, msgItemResp, err
	}

	//存储msg
	msg, err := storeMsg(links, msgType, msgContent, r.LinkKey, convId, msgKey)
	if err != nil {
		log.Println("storeMsg err", err)
		return linkKeys, msgItemResp, err
	}

	//response
	msgItemResp = MsgItemResp{
		ConvId:      convId,
		Id:          msg.Id,
		Key:         msg.Key,
		Type:        msg.Type,
		Content:     msg.Content,
		FromLinkKey: msg.FromLinkKey,
		Created:     msg.Created,
		Updated:     msg.Updated,
	}
	for _, link := range links {
		linkKeys = append(linkKeys, link.Key)
	}
	return linkKeys, msgItemResp, nil
}

func MsgRead(r Request) ([]string, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	msgId, ok := param["msgId"].(string)
	if !ok {
		msgId = ""
	}
	toLinkKey, ok := param["toLinkKey"].(string)
	if !ok {
		toLinkKey = ""
	}
	toLink, _ := linkService.GetByKey(toLinkKey)

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

func MsgListHistory(r Request) ([]string, MsgHistoryResp, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	convId, ok := param["convId"].(string)
	if !ok {
		convId = ""
	}
	offsetParam, ok := param["offset"].(float64)
	if !ok {
		offsetParam = 0
	}
	offset := int(offsetParam)
	limitParam, ok := param["limit"].(float64)
	if !ok {
		limitParam = 20
	}
	limit := int(limitParam)
	link, _ := linkService.GetByKey(r.LinkKey)

	//listMsg
	msgs, err := msgService.ListMsg(convId, offset, limit)
	if err != nil {
		log.Println("listMsg err", err)
	}

	//response
	msgsResp := make([]MsgItemResp, 0)
	for _, m := range *msgs {
		msgItemResp := MsgItemResp{
			ConvId:      convId,
			Id:          m.Id,
			Key:         m.Key,
			Type:        m.Type,
			Content:     m.Content,
			FromLinkKey: m.FromLinkKey,
			Created:     m.Created,
			Updated:     m.Updated,
		}
		msgsResp = append(msgsResp, msgItemResp)
	}
	msgHistoryResp := MsgHistoryResp{
		ConvId: convId,
		Msgs:   msgsResp,
	}
	linkKeys = append(linkKeys, link.Key)
	return linkKeys, msgHistoryResp, err
}

func getLinksByConvId(convId string) ([]model.Link, error) {
	links, err := linkService.ListLink(convId)
	return *links, err
}

func storeMsg(links []model.Link, msgType, content, fromLinkKey, convId, msgKey string) (model.Msg, error) {
	now := time.Now()
	//msg
	msg := model.Msg{
		Id:          util.GetRandomString(11),
		Key:         msgKey,
		Type:        msgType,
		Content:     content,
		FromLinkKey: fromLinkKey,
		Created:     now,
		Updated:     now,
	}
	err := msgService.Create(&msg)
	if err != nil {
		return msg, err
	}

	//msg_conv
	msgConv := model.MsgConv{
		MsgId:  msg.Id,
		ConvId: convId,
	}
	err = msgConvService.Create(&msgConv)
	if err != nil {
		return msg, err
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
		return msg, err
	}

	return msg, nil
}
