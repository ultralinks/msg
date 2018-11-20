package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"msg/userLogic/app"
	"msg/userLogic/util"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if !strings.Contains(authorization, "Bearer") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := string([]byte(authorization)[7:])
		if token == "null" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		mc, err := util.ParseJwtToken(app.Config.Secret.JwtKey, token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		fmt.Println("jwt middlewareuser_id: ", mc["user_id"].(string))
		c.Request.Header.Set("user_id", mc["user_id"].(string))
		c.Next()
	}
}
