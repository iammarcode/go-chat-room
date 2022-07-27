package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/iammarcode/go-chat-room/pkg/logging"
	"github.com/iammarcode/go-chat-room/pkg/util"
	"github.com/iammarcode/go-chat-room/response"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Failed("token required", c)
			logging.Error("token required")
			c.Abort()
			return
		}
		if _, err := util.VerifyToken(token); err != nil {
			response.Failed("unauthorized", c)
			logging.Error("invalid token err: ", err)
			c.Abort()
			return
		}

		c.Next()
	}
}
