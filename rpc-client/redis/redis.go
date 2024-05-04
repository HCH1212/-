package redis

import (
	"context"
	"log"
	"rpc-client/model"

	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 10, //连接池大小,一个Redis实例在同一时刻可以处理的最大连接数,默认10
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalln("redis connect error: ", err)
	} else {
		model.RDB = rdb
	}
}
