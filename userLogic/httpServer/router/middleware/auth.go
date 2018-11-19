package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminUserId := c.GetHeader("User_id")
		if adminUserId != "" {
			c.Next()
		} else {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{"error": "user not found"})
		}
	}
}
