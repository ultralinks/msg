package linkToken

import (
	"msg/msgLogic/app"
	"msg/msgLogic/service/model"
)

func Get(linkId string) (*model.LinkToken, error) {
	lt := &model.LinkToken{}
	d := app.DB.Table("link_token").Where("link_id = ?", linkId).First(lt)
	return lt, d.Error
}
