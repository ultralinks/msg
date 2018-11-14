package msg

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(msg *model.Msg) error {
	d := app.DB.Create(msg)
	return d.Error
}

func ListMsg(convId string) (*[]model.Msg, error) {
	msgs := make([]model.Msg, 0)
	d := app.DB.Table("msg_conv").Select("msg.*").
		Joins("JOIN msg on msg.id = msg_conv.msg_id").
		Where("msg_conv.conv_id = ?", convId).Find(&msgs)
	return &msgs, d.Error
}
