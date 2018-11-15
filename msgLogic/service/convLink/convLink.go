package convLink

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(convLink *model.ConvLink) error {
	d := app.DB.Create(convLink)
	return d.Error
}

func Get(convId, linkId string) (*model.ConvLink, error) {
	money := &model.ConvLink{}
	d := app.DB.Where("conv_id = ? AND link_id = ?", convId, linkId).First(money)
	return money, d.Error
}

func Update(whereMap map[string]interface{}, updateMap map[string]interface{}) error {
	err := app.DB.Table("conv_link").Where(whereMap).Updates(updateMap).Error
	return err
}

func Delete(whereMap map[string]interface{}) error {
	convLink := model.ConvLink{}
	err := app.DB.Where(whereMap).Delete(&convLink).Error
	return err
}
