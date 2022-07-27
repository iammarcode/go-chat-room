package models

import (
	"fmt"
	"github.com/iammarcode/go-chat-room/global"
	"github.com/iammarcode/go-chat-room/pkg/logging"
	"github.com/iammarcode/go-chat-room/pkg/util"
	"gorm.io/gorm"
)

type Model struct {
	ID         int     `gorm:"primaryKey" json:"id"`
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

func Create(table string, data map[string]interface{}) error {
	model, exist := Mapper[table]
	if !exist {
		logging.Error("find label key err")
		return gorm.ErrRecordNotFound
	}
	err := global.Db.Model(model).Create(data).Error
	if err != nil {
		logging.Error(fmt.Sprintf("create %s err: %v", table, err))
		return err
	}

	return nil
}

func Update(table string, data map[string]interface{}) error {
	model, exist := Mapper[table]
	if !exist {
		logging.Error("find label key err")
		return gorm.ErrRecordNotFound
	}
	err := global.Db.Model(model).Updates(data).Error
	if err != nil {
		logging.Error(fmt.Sprintf("update %s err: %v", table, err))
		return err
	}

	return nil
}
