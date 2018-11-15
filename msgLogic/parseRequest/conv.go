package parseRequest

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	convService "msg/msgLogic/service/conv"
	convLinkService "msg/msgLogic/service/convLink"
	linkService "msg/msgLogic/service/link"
	"msg/msgLogic/service/model"
	"msg/msgLogic/util"
)

func ConvCreate(r Request) ([]string, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	convKey := param["convKey"].(string)
	name := param["name"].(string)
	links, err := getLinksByLinkKeys(param["linkKeys"].([]string))
	if err != nil {
		log.Println("get links by linkKeys err", err)
		return linkKeys, err
	}

	//conv
	now := time.Now()
	conv := model.Conv{
		Id:      util.GetRandomString(11),
		Key:     convKey,
		Name:    name,
		Created: now,
		Updated: now,
	}
	err = convService.Create(&conv)
	if err != nil {
		log.Println("create conv err", err)
	}

	//conv_link
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
		}
		convLinkService.Create(&convLink)

		linkKeys = append(linkKeys, link.Key)
	}

	return linkKeys, nil
}

func ConvList(r Request) ([]string, []byte, error) {
	linkKeys := make([]string, 0)
	responseByte := make([]byte, 0)

	link, err := linkService.GetByKey(r.LinkKey)
	if err != nil {
		log.Println("get link by key err", err)
		return linkKeys, responseByte, err
	}

	convs, err := convService.ListByLinkId(link.Id)
	if err != nil {
		log.Println("listConv err", err)
	}
	responseByte, err = json.Marshal(convs)
	if err != nil {
		log.Println("convsByte err", err)
	}

	linkKeys = append(linkKeys, link.Key)
	return linkKeys, responseByte, nil
}

func ConvDelete(r Request) ([]string, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	convId := param["convId"].(string)
	link, err := linkService.GetByKey(r.LinkKey)
	if err != nil {
		log.Println("get link by key err", err)
		return linkKeys, err
	}

	whereMap := map[string]interface{}{
		"conv_id": convId,
		"link_id": link.Id,
	}
	updateMap := map[string]interface{}{
		"is_ignore": 1,
		"updated":   time.Now(),
	}
	if err := convLinkService.Update(whereMap, updateMap); err != nil {
		log.Println("update conv_link is_ignore err", err)
	}

	linkKeys = append(linkKeys, link.Id)
	return linkKeys, nil
}

func ConvJoin(r Request) ([]string, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	convId := param["convId"].(string)
	link, err := linkService.GetByKey(r.LinkKey)
	if err != nil {
		log.Println("get link by key err", err)
		return linkKeys, err
	}

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
	if err := convLinkService.Create(&convLink); err != nil {
		log.Println("create conv_link err", err)
	}

	linkKeys = append(linkKeys, link.Id)
	return linkKeys, nil
}

func ConvLeave(r Request) ([]string, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	convId := param["convId"].(string)
	link, err := linkService.GetByKey(r.LinkKey)
	if err != nil {
		log.Println("get link by key err", err)
		return linkKeys, err
	}

	whereMap := map[string]interface{}{
		"conv_id": convId,
		"link_id": link.Id,
	}
	if err := convLinkService.Delete(whereMap); err != nil {
		log.Println("delete conv_link err", err)
	}

	linkKeys = append(linkKeys, link.Id)
	return linkKeys, nil
}

func ConvInviteLinks(r Request) ([]string, error) {
	linkKeys := make([]string, 0)

	param := r.Param
	convId := param["convId"].(string)
	links, err := getLinksByLinkKeys(param["linkKeys"].([]string))
	if err != nil {
		log.Println("get links by linkKeys err", err)
		return linkKeys, err
	}

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
	convId := param["convId"].(string)
	fromLink, err := linkService.GetByKey(r.LinkKey)
	if err != nil {
		log.Println("get link by key err", err)
		return linkKeys, err
	}

	fromConvLink, _ := convLinkService.Get(convId, fromLink.Id)
	if fromConvLink.IsOwner != 1 {
		return linkKeys, errors.New("only conv owner can remove conv")
	}

	links, err := getLinksByLinkKeys(param["linkKeys"].([]string))
	if err != nil {
		log.Println("get links by linkKeys err", err)
		return linkKeys, err
	}

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
