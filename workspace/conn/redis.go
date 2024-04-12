package conn

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisPool *redis.Client

func InitRedis(address string, port string) error {
	fmt.Println("初始化redis连接...")

	RedisPool = redis.NewClient(&redis.Options{
		Addr:       address + ":" + port, // Redis server address
		Password:   "",                   // no password set
		DB:         0,                    // use default DB
		PoolSize:   10,                   // set the pool size to 100
		MaxRetries: 3,                    // set the maximum number of retries to 3
	})
	_, err := RedisPool.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis连接初始化失败")
		return err
	}
	fmt.Println("redis连接初始化完成")
	return nil

}
