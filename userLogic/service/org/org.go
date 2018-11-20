package org

import (
	"msg/userLogic/app"
	"msg/userLogic/service/model"
)

func Create(org *model.Org) error {
	err := app.DB.Create(org).Error
	return err
}

func Get(orgId string) (*model.Org, error) {
	org := &model.Org{}
	d := app.DB.Where("id = ?", orgId).First(org)
	return org, d.Error
}

func ListByUserId(userId string) (*[]model.Org, error) {
	org := make([]model.Org, 0)
	err := app.DB.Table("org").Select("org.*").
		Joins("left join org_user on org_user.org_id = org.id").
		Where("org_user.user_id = ?", userId).Find(&org).Error
	return &org, err
}
