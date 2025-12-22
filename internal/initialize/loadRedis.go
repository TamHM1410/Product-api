package initialize

import (
	"context"
	"fmt"
	"go-backend-api/internal/global"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func LoadRedis() *redis.Client {
	fmt.Println("✅ Connecting to Redis...")

	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     global.RedisConfig.Host + ":" + strconv.Itoa(global.RedisConfig.Port),
		Password: global.RedisConfig.Password,
		DB:       global.RedisConfig.DB,
	})

	err := global.Rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(fmt.Sprintf("❌ Failed to connect to Redis: %v", err))
	}
	fmt.Println("✅ Redis connection established")
	return global.Rdb

}
