package service

import (
	"encoding/json"
	"github.com/iammarcode/go-chat-room/models"
	"github.com/iammarcode/go-chat-room/pkg/cache"
	"github.com/iammarcode/go-chat-room/pkg/logging"
	"github.com/iammarcode/go-chat-room/pkg/mq"
)

type UserService struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

func (u UserService) Register() error {
	//user := models.User{
	//	Username: u.Username,
	//	Password: u.Password,
	//	Email:    u.Email,
	//	Avatar:   u.Avatar,
	//}
	if u.Avatar == "" {
		u.Avatar = "default.png"
	}

	err := mq.Pubulish(models.Message{"create", "user", map[string]interface{}{"username": u.Username, "password": u.Password, "avatar": u.Avatar, "email": u.Email}})
	if err != nil {
		logging.Error(err)
	}

	return err
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
	data, err := cache.Get(cache.KeyUserList)
	if err != nil {
		logging.Error(err)
		// query from db
		user := models.User{}
		userList, err = user.GetAll()
		if err != nil {
			logging.Error(err)
			return nil, err
		} else {
			// store in redis
			dataByte, err := json.Marshal(userList)
			if err != nil {
				logging.Error(err)
			} else {
				err = cache.Set(cache.KeyUserList, dataByte, cache.DefaultExpiration)
				if err != nil {
					logging.Error(err)
				}
			}

			return userList, nil
		}
	}

	// query from cache
	logging.Info("get from cache, key: ", cache.KeyUserList)
	err = json.Unmarshal([]byte(data), &userList)
	if err != nil {
		logging.Error(err)
		return nil, err
	}

	return userList, nil
}
