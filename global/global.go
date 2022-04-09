package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/whoismarcode/go-chat-room/config"
	"gorm.io/gorm"
)

var (
	Config config.Config
	Db     *gorm.DB
	Cache  *redis.Client
)
