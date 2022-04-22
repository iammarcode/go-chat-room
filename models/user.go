package models

import (
	"github.com/whoismarcode/go-chat-room/global"
	"github.com/whoismarcode/go-chat-room/pkg/logging"
	"gorm.io/gorm"
)

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

func (u User) Create() error {
	if u.Avatar == "" {
		u.Avatar = "default.png"
	}
	user := User{Username: u.Username, Password: u.Password, Email: u.Email, Avatar: u.Avatar}
	err := global.Db.Create(&user).Error
	if err != nil {
		logging.Error(err)
		return err
	}

	return nil
}

func (u User) Update() error {
	user := User{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Avatar:   u.Avatar,
	}
	err := global.Db.Model(&User{}).Where("id = ?", u.ID).Updates(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u User) GetById() (*User, error) {
	var user User
	err := global.Db.Find(&user, "id = ?", u.ID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u User) DeleteById() error {
	err := global.Db.Delete(&User{}, u.ID).Error
	if err != nil {
		return err
	}

	return nil
}

func (u User) GetAll() (*[]User, error) {
	var users []User
	err := global.Db.Find(&users).Error
	if err != nil {
		logging.Error(err)
		return nil, err
	}

	return &users, nil
}

func (u User) CheckUsernamePassword() (bool, error) {
	var user User
	err := global.Db.Select("id").Where(User{Username: u.Username, Password: u.Password}).First(&user).Error
	// query exception
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Error(err)
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	// not found
	logging.Error("user not found")
	return false, nil
}
