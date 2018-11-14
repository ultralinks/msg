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

func ListLink(convId string) (*[]model.Link, error) {
	links := make([]model.Link, 0)
	d := app.DB.Table("conv_link").Select("link.*").
		Joins("JOIN link on link.id = conv_link.link_id").
		Where("conv_link.conv_id = ?", convId).Find(&links)
	return &links, d.Error
}
