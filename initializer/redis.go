package initializer

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/iammarcode/go-chat-room/global"
)

func InitRedis() {
	// connection pool
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", global.Config.Redis.Host, global.Config.Redis.Port),
		Password: global.Config.Redis.Password,
		DB:       0, // use default DB
	})

	// expose redis
	global.RedisClient = rdb
}
