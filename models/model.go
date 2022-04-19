package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	ModifiedAt time.Time      `json:"modified_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}


