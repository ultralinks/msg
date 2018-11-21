package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"msg/userLogic/app"
	_ "msg/userLogic/docs"
	"msg/userLogic/httpServer/router"
	"msg/userLogic/httpServer/router/middleware"
	"msg/userLogic/rpc"
)

func main() {
	app.InitConfig()
	app.InitDB()
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

	r.Use(middleware.CORSMiddleware())
	router.RegisterV1Router(r)
	r.Use(middleware.JWTMiddleware())
	router.RegisterV1RouterJWT(r)

	fmt.Println("starting server ...")
	r.Run(app.Config.Http.Domain + ":" + app.Config.Http.Port)
}
