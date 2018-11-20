package app

import (
	a "msg/userLogic/app"
	"msg/userLogic/service/model"
)

func Create(app *model.App) error {
	return a.DB.Create(app).Error
}

func Get(appId string) (*model.App, error) {
	app := &model.App{}
	d := a.DB.Where("id = ?", appId).First(app)
	return app, d.Error
}

func ListByOrgId(orgId string) (*[]model.App, error) {
	app := make([]model.App, 0)
	err := a.DB.Table("app").Where("org_id = ?", orgId).Find(&app).Error
	return &app, err
}
