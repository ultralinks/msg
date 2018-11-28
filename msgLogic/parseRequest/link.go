package parseRequest

import (
	linkService "msg/msgLogic/service/link"
	"msg/msgLogic/service/model"
)

func LinkGet(r Request) ([]string, model.Link, error) {
	linkKeys := make([]string, 0)

	link, err := linkService.GetByKey(r.LinkKey)

	linkKeys = append(linkKeys, link.Key)
	return linkKeys, *link, err
}
