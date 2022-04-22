package initializer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whoismarcode/go-chat-room/api/v1"
	"github.com/whoismarcode/go-chat-room/global"
	"github.com/whoismarcode/go-chat-room/pkg/jwt"
)

func Router() {
	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	router := gin.New()

	// TODO: beautify gin.Logger() format
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// By default gin.DefaultWriter = os.Stdout
	//router.Use(gin.Logger())


	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	apiV1 := router.Group("/v1")

	// unauthorized
	apiV1.POST("/register", v1.Register)
	apiV1.POST("/login", v1.Login)

	// authorized
	apiV1.Use(jwt.JwtAuth())
	{
		apiV1.POST("/refreshToken", v1.RefreshToken)
		apiV1.GET("/users", v1.UserList)
	}

	router.Run(fmt.Sprintf("%s:%s", global.Config.Server.Host, global.Config.Server.Port))
}
