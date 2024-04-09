package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var Rdb *redis.Client

func InitRedis() error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := Rdb.Ping(context.Background()).Result() // 检验是否链接成功
	if err != nil {
		log.Println(err) // 输出错误日志
		return err
	}

	return nil
}
