package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"msg/userLogic/app"
	_ "msg/userLogic/docs"
	"msg/userLogic/httpServer/router"
)

func main() {
	app.InitConfig()

	app.InitDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if gin.Mode() == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	router.RegisterV1Router(r)

	r.Run(app.Config.Http.Domain + ":" + app.Config.Http.Port)
}
