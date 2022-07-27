package global

import (
	"github.com/go-eyas/toolkit/amqp"
	"github.com/go-redis/redis/v8"
	"github.com/iammarcode/go-chat-room/config"
	"gorm.io/gorm"
)

type Message struct {
	Client *amqp.MQ
	Queue  *amqp.Queue
	Exchange *amqp.Exchange
}

var (
	Config      config.Config
	Db          *gorm.DB
	RedisClient *redis.Client
	Mq          *Message
)
