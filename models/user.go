package models

import (
	"github.com/whoismarcode/go-chat-room/global"
	"github.com/whoismarcode/go-chat-room/service"
)

type User struct {
	Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

func Create(userService service.User) error {
	user := User{Name: userService.Name, Password: userService.Password, Avatar: userService.Avatar}
	err := global.Db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func Update(userService service.User) error {
	err := global.Db.Model(&User{}).Where("id = ?", userService.Id).Updates(User{Name: userService.Name, Password: userService.Password, Avatar: userService.Avatar}).Error
	if err != nil {
		return err
	}

	return nil
}

func GetById(id int) (*User, error) {
	// TODO: redis

	var user User
	err := global.Db.Find(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteById(id int) error {
	err := global.Db.Delete(&User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func GetAll() (*[]User, error) {
	// TODO: redis

	var users []User
	err := global.Db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}
