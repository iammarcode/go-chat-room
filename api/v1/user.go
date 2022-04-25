package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/whoismarcode/go-chat-room/models"
	"github.com/whoismarcode/go-chat-room/pkg/logging"
	"github.com/whoismarcode/go-chat-room/response"
	"github.com/whoismarcode/go-chat-room/service"
)

func Register(c *gin.Context) {
	// check body
	var registerReq models.RegisterRequest
	if err := c.ShouldBindBodyWith(&registerReq, binding.JSON); err != nil {
		logging.Error(err)
		response.Failed("Bad Request body", c)
		return
	}

	// check name and password
	userService := service.UserService{
		Username: registerReq.Username,
		Password: registerReq.Password,
		Email:    registerReq.Email,
	}

	err := userService.Register()
	if err != nil {
		response.Failed("Unknown error", c)
		return
	}

	response.Success("Register Success", nil, c)
}

func UserList(c *gin.Context) {
	userService := service.UserService{}
	userList, err := userService.UserList()
	if err != nil {
		logging.Error(err)
		response.Failed("Unknown error", c)
	}

	response.Success("Success", *userList, c)
}
