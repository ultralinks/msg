package msg

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(msg *model.Msg) error {
	d := app.DB.Create(msg)
	return d.Error
}

func ListMsg(convId string, offset, limit int) (*[]model.Msg, error) {
	msgs := make([]model.Msg, 0)
	if limit == 0 {
		limit = 10
	}

	d := app.DB.Table("msg_conv").Select("msg.*").
		Joins("left join msg on msg.id = msg_conv.msg_id").
		Where("msg_conv.conv_id = ?", convId).
		Offset(offset).Limit(limit).Order("msg.created desc").
		Find(&msgs)

	return &msgs, d.Error
}
