package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/whoismarcode/go-chat-room/global"
	"github.com/whoismarcode/go-chat-room/pkg/logging"
	"time"
)

var ctx = context.Background()

func Set(key string, data interface{}, expiration time.Duration) error {
	err := global.RedisClient.Set(ctx, key, data, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func ExpireAt(key string, expiration time.Duration) error {
	err := global.RedisClient.Expire(ctx, key, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func Get(key string) (string, error) {
	val, err := global.RedisClient.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		logging.Error("key does not exist:", key)
	case err != nil:
		logging.Error("key get failed, err:", key, err)
	}

	return val, err
}

func Exists(keys ...string) (bool, error) {
	result := false
	flag, err := global.RedisClient.Exists(ctx, keys...).Result()
	if err != nil {
		logging.Error("key check failed, err:", err)
		return false, err
	}

	if flag > 0 {
		result = true
	}

	return result, nil
}

func Delete(keys ...string) (bool, error) {
	_, err := global.RedisClient.Del(ctx, keys...).Result()
	if err != nil {
		logging.Error("key del failed err:", err)
		return false, err
	}

	return true, nil
}
