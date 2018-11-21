package msg

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(msg *model.Msg) error {
	d := app.DB.Create(msg)
	return d.Error
}

func Get(msgId string) (*model.Msg, error) {
	msg := &model.Msg{}
	d := app.DB.Where("id = ?", msgId).First(msg)
	return msg, d.Error
}

func GetByConvId(convId string) (*model.Msg, error) {
	msg := &model.Msg{}
	d := app.DB.Table("msg").Select("msg.*").
		Joins("left join msg_conv on msg_conv.msg_id = msg.id").
		Where("msg_conv.conv_id = ?", convId).
		Order("msg.created desc").First(msg)
	return msg, d.Error
}

func ListMsg(convId string, offset, limit int) (*[]model.Msg, error) {
	msgs := make([]model.Msg, 0)
	if limit == 0 {
		limit = 20
	}

	d := app.DB.Table("msg").Select("msg.*").
		Joins("left join msg_conv on msg_conv.msg_id = msg.id").
		Where("msg_conv.conv_id = ?", convId).
		Offset(offset).Limit(limit).Order("msg.created desc").
		Find(&msgs)

	return &msgs, d.Error
}
