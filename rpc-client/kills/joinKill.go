package kills

import (
	"fmt"
	"log"
	"rpc-client/model"
	"rpc-client/resp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func JoinKill(c *gin.Context) {
	var u model.User
	var h model.Hoster
	var err error
	var s string

	//利用redis分布式锁保证秒杀秩序
	lockKey := "my_lock"            //锁名
	lockValue := "lock_value"       //锁值
	lockTimeout := 10 * time.Second //锁的超时时间
	// 尝试获取锁
	success, err := model.RDB.SetNX(c, lockKey, lockValue, lockTimeout).Result()
	if err != nil {
		panic(err)
	}
	if success {
		fmt.Println("获取锁成功")
		// 执行业务逻辑

		//token验证身份
		u.Token = c.PostForm("token")
		//获取商家token以实现商品更改
		h.Token, err = model.RDB.Get(c, "hostertoken").Result()
		if err != nil {
			log.Fatalln("get hostertoken error:", err)
		}
		//一个用户只能买一个商品
		s, err = model.RDB.Get(c, h.Token+"keep").Result()
		if err != nil {
			log.Fatalln("get keep error:", err)
		}
		h.Goods.Keep, _ = strconv.Atoi(s)
		if h.Goods.Keep > 0 {
			h.Goods.Keep--
		}
		model.RDB.Set(c, h.Token+"keep", h.Goods.Keep, 0)
		//付款操作
		u.IsBuy, _ = strconv.ParseBool(c.PostForm("isbuy"))
		if u.IsBuy {
			model.RDB.Set(c, "isbuy", true, 0)
		} else {
			model.RDB.Set(c, "isbuy", false, 0)
		}

		// 释放锁
		//创建一个Lua脚本对象，用于释放锁
		//Lua脚本用于实现Redis分布式锁的原子性操作，确保了在释放锁时的线程安全性
		unlockScript := redis.NewScript(`
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end
		`)
		//运行Lua脚本来释放锁
		_, err := unlockScript.Run(c, model.RDB, []string{lockKey}, lockValue).Result()
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("获取锁失败")
	}

	//释放锁后，如果用户还没有买下商品，则不能再操作并返回库存
	s, err = model.RDB.Get(c, "isbuy").Result()
	if err != nil {
		log.Fatalln("get isbuy error:", err)
	}
	u.IsBuy, _ = strconv.ParseBool(s)
	if u.IsBuy {
		resp.OKWithData(c, "kill success")
	} else {
		h.Goods.Keep++
		model.RDB.Set(c, h.Token+"keep", h.Goods.Keep, 0)
		c.JSON(200, gin.H{
			"status": "20001",
			"info":   "kill fail",
		})
	}
}
