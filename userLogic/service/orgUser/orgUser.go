package orgUser

import (
	"msg/userLogic/app"
	"msg/userLogic/service/model"
)

func Create(orgUser *model.OrgUser) error {
	return app.DB.Create(orgUser).Error
}
