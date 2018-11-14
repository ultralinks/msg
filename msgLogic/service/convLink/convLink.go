package convLink

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(convLink *model.ConvLink) error {
	d := app.DB.Create(convLink)
	return d.Error
}
