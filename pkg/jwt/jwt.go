package jwt

import (
	"github.com/gin-gonic/gin"
	util2 "github.com/whoismarcode/go-chat-room/pkg/util"
	"github.com/whoismarcode/go-chat-room/response"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Failed("token required", c)
			c.Abort()
			return
		}
		if _, err := util2.VerifyToken(token); err != nil {
			response.Failed("unauthorized", c)
			c.Abort()
			return
		}

		c.Next()
	}
}
