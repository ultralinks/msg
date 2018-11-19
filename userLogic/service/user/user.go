package user

import (
	"time"

	"msg/userLogic/app"
	"msg/userLogic/service/model"
)

func Create(user *model.User) error {
	err := app.DB.Create(user).Error
	return err
}

func Get(whereMap map[string]interface{}) (*model.User, error) {
	user := &model.User{}
	d := app.DB.Where(whereMap).First(user)
	return user, d.Error
}

func UpdateLogined(userId string, logined time.Time) error {
	d := app.DB.Table("user").Where("id = ?", userId).Update("logined", logined)
	return d.Error
}
