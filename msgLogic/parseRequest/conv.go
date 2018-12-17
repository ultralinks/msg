package parseRequest

import (
	"errors"
	"log"
	"time"

	convService "msg/msgLogic/service/conv"
	convLinkService "msg/msgLogic/service/convLink"
	convSendService "msg/msgLogic/service/convSend"
	linkService "msg/msgLogic/service/link"
	"msg/msgLogic/service/model"
	msgService "msg/msgLogic/service/msg"
	"msg/msgLogic/util"
)

//listConv中的每一项
type ConvItem struct {
	ConvId    string       `json:"convId"`
	Conv      model.Conv   `json:"conv"`
	Links     []model.Link `json:"links"`
	LastMsg   model.Msg    `json:"lastMsg"`
	UnreadNum int          `json:"unreadNum"`
}

func ConvCreate(r Request) ([]string, ConvItem, error) {
	linkKeys := make([]string, 0)

	link, _ := linkService.GetByKey(r.LinkKey)

	param := r.Param
	convKey, ok := param["convKey"].(string)
	if !ok {
		convKey = ""
	}
	convType, ok := param["convType"].(string)
	if !ok {
		convType = "single"
	}
	name, ok := param["convName"].(string)
	if !ok {
		name = "unknown"
	}
	avt, ok := param["convAvt"].(string)
	if !ok {
		avt = ""
	}

	//convLinkKeys
	convLinkKeys := make([]string, 0)
	linkKeyInterfaces, ok := param["linkKeys"].([]interface{})
	if !ok {
		linkKeyInterfaces = make([]interface{}, 0)
	}
	for _, l := range linkKeyInterfaces {
		convLinkKeys = append(convLinkKeys, l.(string))
	}
	links, err := getLinksByLinkKeys(convLinkKeys)
	if err != nil {
		log.Println("get links by linkKeys err", err)
	}

	//single type conv if existed
	if convType == "single" {
		existedConv, _ := convService.GetByTwoLinkId(links[0].Id, links[1].Id)
		if existedConv.Id != "" {
			lastMsg, _ := msgService.GetByConvId(existedConv.Id)
			existedConvItem := ConvItem{
				ConvId:    existedConv.Id,
				Conv:      *existedConv,
				Links:     links,
				LastMsg:   *lastMsg,
				UnreadNum: 0,
			}
			return []string{link.Key}, existedConvItem, nil
		}
	}

	//create conv
	now := time.Now()
	conv := model.Conv{
		Id:      util.GetRandomString(11),
		Key:     convKey,
		Name:    name,
		Avt:     avt,
		Type:    convType,
		Created: now,
		Updated: now,
	}
	err = convService.Create(&conv)
	if err != nil {
		log.Println("create conv err", err)
	}

	//create conv_link
	for _, link := range links {
		isOwner := 0
		if link.Key == r.LinkKey {
			isOwner = 1
		}
		convLink := model.ConvLink{
			ConvId:   conv.Id,
			LinkId:   link.Id,
			IsOwner:  isOwner,
			IsMute:   0,
			IsIgnore: 0,
			Created:  now,
			Updated:  now,
		}
		convLinkService.Create(&convLink)
		linkKeys = append(linkKeys, link.Key)
	}

	lastMsg, _ := msgService.GetByConvId(conv.Id)
	convItem := ConvItem{
		ConvId:    conv.Id,
		Conv:      conv,
		Links:     links,
		LastMsg:   *lastMsg,
		UnreadNum: 0,
	}
	return linkKeys, convItem, err
}

func ConvList(r Request) ([]string, []ConvItem, error) {
	linkKeys := make([]string, 0)

	link, _ := linkService.GetByKey(r.LinkKey)

	//list convs
	convs, err := convService.ListByLinkId(link.Id)
	if err != nil {
		log.Println("listConv err", err)
	}

	//response: convs add links and lastMsg
	convList := make([]ConvItem, 0)
	for _, conv := range *convs {
		links, _ := linkService.ListLink(conv.Id)
		lastMsg, _ := msgService.GetByConvId(conv.Id)
		unreadNum, _ := getUnreadNum(conv.Id, link.Id)
		convItem := ConvItem{
			ConvId:    conv.Id,
			Conv:      conv,
			Links:     *links,
			LastMsg:   *lastMsg,
			UnreadNum: unreadNum,
		}
		convList = append(convList, convItem)
	}

	linkKeys = append(linkKeys, link.Key)
	return linkKeys, convList, err
}

func ConvDelete(r Request) ([]string, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	convId, ok := param["convId"].(string)
	if !ok {
		convId = ""
	}
	link, _ := linkService.GetByKey(r.LinkKey)

	//update convLink
	whereMap := map[string]interface{}{
		"conv_id": convId,
		"link_id": link.Id,
	}
	updateMap := map[string]interface{}{
		"is_ignore": 1,
		"updated":   time.Now(),
	}
	err := convLinkService.Update(whereMap, updateMap)
	if err != nil {
		log.Println("update conv_link is_ignore err", err)
	}

	linkKeys = append(linkKeys, link.Key)
	return linkKeys, err
}

func ConvJoin(r Request) ([]string, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	convId, ok := param["convId"].(string)
	if !ok {
		convId = ""
	}
	link, err := linkService.GetByKey(r.LinkKey)
	if err != nil {
		log.Println("get link by key err", err)
		return linkKeys, err
	}

	//create convLink
	now := time.Now()
	convLink := model.ConvLink{
		ConvId:   convId,
		LinkId:   link.Id,
		IsOwner:  0,
		IsMute:   0,
		IsIgnore: 0,
		Created:  now,
		Updated:  now,
	}
	err = convLinkService.Create(&convLink)
	if err != nil {
		log.Println("create conv_link err", err)
	}

	linkKeys = append(linkKeys, link.Key)
	return linkKeys, err
}

func ConvLeave(r Request) ([]string, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	convId, ok := param["convId"].(string)
	if !ok {
		convId = ""
	}
	link, err := linkService.GetByKey(r.LinkKey)
	if err != nil {
		log.Println("get link by key err", err)
		return linkKeys, err
	}

	//delete convLink
	whereMap := map[string]interface{}{
		"conv_id": convId,
		"link_id": link.Id,
	}
	err = convLinkService.Delete(whereMap)
	if err != nil {
		log.Println("delete conv_link err", err)
	}

	linkKeys = append(linkKeys, link.Key)
	return linkKeys, err
}

func ConvInviteLinks(r Request) ([]string, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	convId, ok := param["convId"].(string)
	if !ok {
		convId = ""
	}
	links, err := getLinksByLinkKeys(param["linkKeys"].([]string))
	if err != nil {
		log.Println("get links by linkKeys err", err)
		return linkKeys, err
	}

	//create convLink
	now := time.Now()
	for _, link := range links {
		convLink := model.ConvLink{
			ConvId:   convId,
			LinkId:   link.Id,
			IsOwner:  0,
			IsMute:   0,
			IsIgnore: 0,
			Created:  now,
			Updated:  now,
		}
		convLinkService.Create(&convLink)
		linkKeys = append(linkKeys, link.Key)
	}

	return linkKeys, nil
}

func ConvRemoveLinks(r Request) ([]string, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	convId, ok := param["convId"].(string)
	if !ok {
		convId = ""
	}
	fromLink, _ := linkService.GetByKey(r.LinkKey)
	links, err := getLinksByLinkKeys(param["linkKeys"].([]string))
	if err != nil {
		log.Println("get links by linkKeys err", err)
		return linkKeys, err
	}

	//owner == 1
	fromConvLink, _ := convLinkService.Get(convId, fromLink.Id)
	if fromConvLink.IsOwner != 1 {
		return linkKeys, errors.New("only conv owner can remove conv")
	}

	//delete convLink
	for _, link := range links {
		whereMap := map[string]interface{}{
			"conv_id": convId,
			"link_id": link.Id,
		}
		convLinkService.Delete(whereMap)
		linkKeys = append(linkKeys, link.Key)
	}

	return linkKeys, nil
}

func getLinksByLinkKeys(linkKeys []string) ([]model.Link, error) {
	links := make([]model.Link, 0)
	var err error
	for _, linkKey := range linkKeys {
		link, err := linkService.GetByKey(linkKey)
		if err != nil {
			log.Println("get link by key err", err)
		}
		links = append(links, *link)
	}
	return links, err
}

func getUnreadNum(convId, linkId string) (int, error) {
	num, err := convSendService.CountUnread(convId, linkId)
	return num, err
}
