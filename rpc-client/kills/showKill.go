package kills

import (
	"log"
	"net/http"
	"rpc-client/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 商家查看自己的商品
func ShowKill(c *gin.Context) {
	var h model.Hoster
	h.Token = c.PostForm("token")
	flag, err := model.RDB.Get(c, "flag").Result()
	if err != nil || flag != model.HosterFlag {
		log.Fatalln("CreatKill flag error :", err)
	}
	res, err := model.RDB.Get(c, h.Token+"name").Result()
	if err != nil {
		log.Fatalln("get name error:", err)
	}
	h.Goods.Name = res
	res, err = model.RDB.Get(c, h.Token+"price").Result()
	if err != nil {
		log.Fatalln("get price error:", err)
	}
	h.Goods.Price, _ = strconv.ParseFloat(res, 64)
	res, err = model.RDB.Get(c, h.Token+"keep").Result()
	if err != nil {
		log.Fatalln("get keep error:", err)
	}
	h.Goods.Keep, _ = strconv.Atoi(res)
	c.JSON(http.StatusOK, gin.H{
		"stauts": 10000,
		"info":   h.Goods,
	})
}
