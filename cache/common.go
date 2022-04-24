package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

var (
	RedisClient *redis.Client

	ctx = context.Background()
)

// Redis 在中间件中初始化redis链接
func Redis(connString string) {
	db, _ := strconv.ParseUint(connString, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       int(db),
		PoolSize: 100, // 连接池大小
	})

	if err := client.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	RedisClient = client
}
