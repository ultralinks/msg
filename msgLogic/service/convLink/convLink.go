package convLink

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(convLink *model.ConvLink) error {
	d := app.DB.Create(convLink)
	return d.Error
}

func ListLink(convId int) (*[]model.Link, error) {
	links := make([]model.Link, 0)
	d := app.DB.Table("conv_link").Select("link.*").
		Joins("JOIN link on link.id = conv_link.link_id").
		Where("conv_link.conv_id = ?", convId).Find(&links)
	return &links, d.Error
}
