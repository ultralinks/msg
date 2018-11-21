package conv

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(conv *model.Conv) error {
	d := app.DB.Create(conv)
	return d.Error
}

func GetByTwoLinkId(linkId1, linkId2 string) (*model.Conv, error) {
	conv := &model.Conv{}
	sql := "select conv.* from conv" +
		"left join conv_link as cl1 on cl1.conv_id = conv.id" +
		"left join conv_link as cl2 on cl2.conv_id = conv.id" +
		"where conv.type = 'single' and cl1.link_id = ? and cl2.link_id = ?"
	d := app.DB.Raw(sql, linkId1, linkId2).First(conv)

	return conv, d.Error
}

func ListByLinkId(linkId string) (*[]model.Conv, error) {
	convs := make([]model.Conv, 0)
	d := app.DB.Table("conv").Select("conv.*").
		Joins("left join conv_link on conv_link.conv_id = conv.id").
		Where("conv_link.link_id = ?", linkId).
		Order("conv.created desc").Find(&convs)
	return &convs, d.Error
}
