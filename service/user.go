package service

import (
	"encoding/json"
	"github.com/whoismarcode/go-chat-room/models"
	"github.com/whoismarcode/go-chat-room/pkg/cache"
	"github.com/whoismarcode/go-chat-room/pkg/logging"
)

type UserService struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

func (u UserService) Register() error {
	user := models.User{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Avatar:   u.Avatar,
	}

	return user.Create()
}

func (u UserService) Update() error {
	user := models.User{
		Model:    models.Model{ID: u.ID},
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Avatar:   u.Avatar,
	}

	return user.Update()
}

func (u UserService) DeleteById() error {
	user := models.User{
		Model: models.Model{ID: u.ID},
	}

	return user.DeleteById()
}

func (u UserService) GetById(id int) (*models.User, error) {
	user := models.User{
		Model: models.Model{ID: u.ID},
	}

	return user.GetById()
}

func (u UserService) Login() (bool, error) {
	user := models.User{
		Username: u.Username,
		Password: u.Password,
	}

	return user.CheckUsernamePassword()
}

func (u UserService) UserList() (*[]models.User, error) {
	var userList *[]models.User
	data, err := cache.Get("userList")
	if err != nil {
		// query from db
		user := models.User{}
		userList, err = user.GetAll()
		if err != nil {
			return nil, err
		} else {
			return userList, nil
		}
	}

	// query from cache
	err = json.Unmarshal([]byte(data), &userList)
	if err != nil {
		logging.Error(err)
		return nil, err
	}

	return userList, nil
}
