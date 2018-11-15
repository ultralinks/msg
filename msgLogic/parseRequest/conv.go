package parseRequest

import (
	"encoding/json"
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
		convLink := model.ConvLink{
			ConvId:   conv.Id,
			LinkId:   link.Id,
			IsOwner:  0,
			IsMute:   0,
			IsIgnore: 0,
			Created:  now,
		}
		convLinkService.Create(&convLink)
	}

	linkKeys = append(linkKeys, r.LinkKey)
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
