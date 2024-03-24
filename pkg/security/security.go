package security

import (
	"errors"
	"github.com/bautista00/Final_BackEnd_Go/pkg/web"
	"github.com/gin-gonic/gin"
	"os"
)


func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("auth-key")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			c.Abort()
			return
		}
		if token != os.Getenv("AUTHENTICATION_TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			c.Abort()
			return
		}
		c.Next()
	}
}