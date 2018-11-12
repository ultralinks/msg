package msg

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(msg *model.Msg) error {
	d := app.DB.Create(msg)
	return d.Error
}
