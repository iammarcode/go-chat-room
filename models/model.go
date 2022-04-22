package models

import (
	"github.com/whoismarcode/go-chat-room/pkg/logging"
	"github.com/whoismarcode/go-chat-room/pkg/util"
	"gorm.io/gorm"
)

type Model struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	CreatedAt  *string `json:"created_at"`
	ModifiedAt *string `json:"modified_at"`
	DeletedAt  *string `json:"deleted_at"`
}

func (model *Model) BeforeCreate(tx *gorm.DB) error {
	if model.CreatedAt == nil {
		tx.Statement.SetColumn("CreatedAt", util.NowTimeString())
	}
	return nil
}

func (model *Model) BeforeDelete(tx *gorm.DB) error {
	err := tx.Set("DeletedAt", util.NowTimeString()).Error
	if err != nil {
		logging.Error(err)
		return err
	}
	return nil
}

func (model *Model) BeforeUpdate(tx *gorm.DB) error {
	err := tx.Set("ModifiedAt", util.NowTimeString()).Error
	if err != nil {
		logging.Error(err)
		return err
	}
	return nil
}
