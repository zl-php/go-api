package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"longtu/cache"
	"time"
)

func RedisTest(ctx *gin.Context) {

	// 写入redis
	err := cache.RedisClient.Set(ctx, "zhoulei", "222222222", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// 读取redis
	val, err := cache.RedisClient.Get(ctx, "zhoulei").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key", val)
	}
}
