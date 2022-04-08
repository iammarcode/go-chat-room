package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id uint32
	Name string
	Password string
}
