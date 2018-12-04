package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"msg/msgLogic/rpc"

	"msg/msgLogic/app"
	"msg/msgLogic/httpServer/router"
	_ "order-server/docs"
)

func main() {
	app.InitConfig()
	app.InitDB()
	go rpc.InitRpcClient()
	go rpc.RunRpcServer()

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

	address := app.Config.Http.Domain + ":" + app.Config.Http.Port
	r.Run(address) // listen and serve on 0.0.0.0:8080

	for {
		time.Sleep(time.Minute)
	}
}
