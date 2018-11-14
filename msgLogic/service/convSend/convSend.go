package convSend

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(convSend *model.ConvSend) error {
	d := app.DB.Create(convSend)
	return d.Error
}

func Update(whereMap map[string]interface{}, updateMap map[string]interface{}) error {
	err := app.DB.Table("conv_send").Where(whereMap).Updates(updateMap).Error
	return err
}
