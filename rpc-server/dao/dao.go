package dao

import (
	"context"
	"fmt"
	"log"
	"rpc-server/model"

	"github.com/go-redis/redis/v8"
	_ "github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
)

func init() {
	InitDB()
	InitRedis()
}

// 初始化gorm
func InitDB() {
	var (
		username = "root"
		password = "041212"
		host     = "localhost"
		port     = "3306"
		dbName   = "rpc"
		dbUrl    = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	)
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln("open sql error : ", err) //打印错误并退出
	}
	DB = db
	db.AutoMigrate(&model.User{}) //可自动创建、更新、删除表
	db.AutoMigrate(&model.Hoster{})
}

// 初始化redis
func InitRedis() {
	rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 10, //连接池大小,一个Redis实例在同一时刻可以处理的最大连接数,默认10
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalln("redis connect error: ", err)
	} else {
		RDB = rdb
	}
}
