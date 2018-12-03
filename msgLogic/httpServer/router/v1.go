package router

import (
	"github.com/gin-gonic/gin"
	"msg/msgLogic/httpServer/handler/link"
)

func RegisterV1Router(r *gin.Engine) {
	r.POST("/links", link.Link)
	r.GET("/link/:key", link.Get)
}
