package httpServer

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"msg/msgLogic/app"
	"msg/msgLogic/httpServer/handler/Conv"
	"msg/msgLogic/httpServer/handler/conv_link"
	_ "order-server/docs"
)

func Start() {
	//app.InitConfig()
	//app.InitDB()
	//app.InitRpcClient()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if gin.Mode() == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.POST("/conv", conv.Create)
	r.POST("/conv-link", conv_link.Create)

	r.Run(app.Config.Http.Domain + ":" + app.Config.Http.Port) // listen and serve on 0.0.0.0:8080
}
