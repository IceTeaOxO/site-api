package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func InitializeRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 伺服器地址
		Password: "",               // Redis 密碼，若無可不填
		DB:       0,                // 使用的 Redis 資料庫
	})

	// 測試 Redis 連接
	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis 連接成功！")
}

func GetFromCache(key string) (string, error) {
	val, err := RDB.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func SetToCache(key string, value string, expiration time.Duration) error {
	err := RDB.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func ConvertJsonToString(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
