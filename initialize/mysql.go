package initialize

import (
	"fmt"
	"github.com/whoismarcode/go-chat-room/global"
	"github.com/whoismarcode/go-chat-room/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func Mysql() {
	// connection mysql
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", global.Config.Mysql.Username, global.Config.Mysql.Password, global.Config.Mysql.Host, global.Config.Mysql.DbName), // data source name
		DefaultStringSize: 256,
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		logging.Fatal("Mysql initialize gorm.Open() err: %v", err)
	}

	// connection pool
	sqlDB, err := db.DB()
	if err != nil {
		logging.Fatal("Mysql initialize db.DB() err: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// expose db instance
	global.Db = db
}
