package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func InitializeRedis() {
	envFilePath := ".env"
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 讀取環境變數的值
	redisHost := os.Getenv("REDIS_HOST")

	RDB = redis.NewClient(&redis.Options{
		// Redis 伺服器地址
		Addr:     fmt.Sprintf("%s:6379", redisHost),
		Password: "", // Redis 密碼，若無可不填
		DB:       0,  // 使用的 Redis 資料庫
	})

	// 測試 Redis 連接
	_, errPing := RDB.Ping(context.Background()).Result()
	if errPing != nil {
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
