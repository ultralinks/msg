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

func ConvCreate(r Request, requestByte []byte) {
	param := r.Param
	convKey := param["convKey"].(string)
	name := param["name"].(string)
	links, err := getLinksByLinkKeys(param["linkKeys"].([]string))
	if err != nil {
		log.Println("get links by linkKeys err", err)
		return
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

	receiveSendData(r.LinkKey, requestByte)
}

func ConvList(r Request) {
	link, err := linkService.GetByKey(r.LinkKey)
	if err != nil {
		log.Println("get link by key err", err)
		return
	}

	convs, err := convService.ListByLinkId(link.Id)
	if err != nil {
		log.Println("listConv err", err)
	}
	convsByte, err := json.Marshal(convs)
	if err != nil {
		log.Println("convsByte err", err)
	}

	receiveSendData(link.Key, convsByte)
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
