package convSend

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(convSend *model.ConvSend) error {
	d := app.DB.Create(convSend)
	return d.Error
}
