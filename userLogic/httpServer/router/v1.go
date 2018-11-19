package router

import (
	"github.com/gin-gonic/gin"
	"msg/userLogic/httpServer/handler/user"
)

func RegisterV1Router(r *gin.Engine) {
	userGroup := r.Group("user")
	{
		userGroup.POST("/regByEmail", user.RegByEmail)
		userGroup.POST("/loginByEmail", user.LoginByEmail)
	}
}
