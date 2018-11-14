package conv

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(conv *model.Conv) error {
	d := app.DB.Create(conv)
	return d.Error
}

func ListByLinkId(linkId string) (*[]model.Conv, error) {
	convs := make([]model.Conv, 0)
	d := app.DB.Table("conv").Select("conv.*").
		Joins("conv_link on conv_link.link_id = link.id").
		Where("link.id = ?", linkId).Find(&convs)
	return &convs, d.Error
}
