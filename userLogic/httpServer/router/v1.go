package router

import (
	"github.com/gin-gonic/gin"
	"msg/userLogic/httpServer/handler/app"
	"msg/userLogic/httpServer/handler/org"
	"msg/userLogic/httpServer/handler/user"
)

func RegisterV1Router(r *gin.Engine) {
	commonGroup := r.Group("")
	{
		commonGroup.POST("/regByEmail", user.RegByEmail)
		commonGroup.POST("/loginByEmail", user.LoginByEmail)
	}
}

func RegisterV1RouterJWT(r *gin.Engine) {
	commonGroup := r.Group("")
	{
		commonGroup.POST("/org", org.Create)
		commonGroup.GET("/org/:id", org.Get)
		commonGroup.GET("/orgs", org.List)

		commonGroup.POST("/app", app.Create)
		commonGroup.GET("/app/:id", app.Get)
		commonGroup.GET("/apps/:orgId", app.List)
	}
}
