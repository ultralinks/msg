package link

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Create(l *model.Link, lt *model.LinkToken) error {
	var count int64
	app.DB.Table("link").Where("`key` = ?", l.Key).Count(&count)
	if count != 0 {
		return nil
	}

	// 开始事务
	tx := app.DB.Begin()
	if err := tx.Create(l).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(lt).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func Get(linkId string) (*model.Link, error) {
	link := &model.Link{}
	d := app.DB.Where("id = ?", linkId).First(link)
	return link, d.Error
}

func GetByKey(linkKey string) (*model.Link, error) {
	link := &model.Link{}
	d := app.DB.Where("`key` = ?", linkKey).First(link)
	return link, d.Error
}

func ListLink(convId string) (*[]model.Link, error) {
	links := make([]model.Link, 0)
	d := app.DB.Table("conv_link").Select("link.*").
		Joins("left join link on link.id = conv_link.link_id").
		Where("conv_link.conv_id = ?", convId).Find(&links)
	return &links, d.Error
}
