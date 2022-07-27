package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/iammarcode/go-chat-room/models"
	"github.com/iammarcode/go-chat-room/pkg/util"
	"github.com/iammarcode/go-chat-room/response"
	"github.com/iammarcode/go-chat-room/service"
)

func Login(c *gin.Context) {
	// check body
	var loginReq models.LoginRequest
	if err := c.ShouldBindBodyWith(&loginReq, binding.JSON); err != nil {
		response.Failed("Bad Request body", c)
		return
	}

	// check name and password
	userService := service.UserService{
		Username: loginReq.Username,
		Password: loginReq.Password,
	}
	result, err := userService.Login()
	if err != nil || result == false {
		response.Failed("Login failed", c)
		return
	}

	// gen token
	token, err := util.GenerateToken(loginReq.Username, loginReq.Password)
	if err != nil {
		response.Failed("Token generation failed", c)
		return
	}
	data := models.LoginResponseData{Token: token}

	response.Success("Login Success", data, c)
}

func RefreshToken(context *gin.Context) {

}
