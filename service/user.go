package service

import (
	"github.com/whoismarcode/go-chat-room/logging"
	"github.com/whoismarcode/go-chat-room/models"
)

type User struct {
	Id       int
	Name     string
	Password string
	Avatar   string
}

func (userService User) Create() error {
	err := models.Create(userService)
	if err != nil {
		logging.Error(err)
		return err
	}

	return nil
}

func (userService User) Update() error {
	err := models.Update(userService)
	if err != nil {
		logging.Error(err)
		return err
	}

	return nil
}

func (userService User) DeleteById(id int) error {
	err := models.DeleteById(id)
	if err != nil {
		logging.Error(err)
		return err
	}

	return nil
}

func (userService User) GetById(id int) (*models.User, error) {
	user, err := models.GetById(id)
	if err != nil {
		logging.Error(err)
		return nil, err
	}

	return user, nil
}
