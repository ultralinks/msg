package service

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(conv *model.Conv) error {
	d := app.DB.Create(conv)
	return d.Error
}
