package msgConv

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(msgConv *model.MsgConv) error {
	d := app.DB.Create(msgConv)
	return d.Error
}
