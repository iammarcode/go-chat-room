package cache

import (
	"context"
	"encoding/json"
	"github.com/whoismarcode/go-chat-room/global"
	"time"
)

// Set a key/value
func Set(key string, data interface{}, time time.Duration) error {
	ctx := context.Background()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = global.RedisClient.Set(ctx, key, value, time).Err()
	if err != nil {
		return err
	}

	return nil
}


// Get a value
func Get(key string, data interface{}, time time.Duration) error {
	ctx := context.Background()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = global.RedisClient.Set(ctx, key, value, time).Err()
	if err != nil {
		return err
	}

	return nil
}

