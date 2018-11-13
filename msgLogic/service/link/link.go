package link

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(link *model.Link) error {
	err := app.DB.Create(link).Error
	return err
}

func Get(linkId string) (*model.Link, error) {
	link := &model.Link{}
	d := app.DB.Where("id = ?", linkId).First(link)
	return link, d.Error
}

func GetByKey(linkKey string) (*model.Link, error) {
	link := &model.Link{}
	d := app.DB.Where("`key` = ?", linkKey).First(link)
	return link, d.Error
}