package kills

import (
	"log"
	"rpc-client/model"
	"rpc-client/resp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatKill(c *gin.Context) {
	var h model.Hoster
	h.Token = c.PostForm("token")
	flag, err := model.RDB.Get(c, "flag").Result()
	if err != nil || flag != model.HosterFlag {
		log.Fatalln("CreatKill flag error :", err)
	}
	//验证完商家身份，接下来处理创建活动逻辑
	//一个商家只卖一种商品
	h.Goods.Name = c.PostForm("name")
	h.Goods.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	h.Goods.Keep, _ = strconv.Atoi(c.PostForm("keep"))
	h.Goods.Timeout, _ = strconv.Atoi(c.PostForm("timeout"))
	//直接存入缓存，以便快速获取
	model.RDB.Set(c, h.Token+"name", h.Goods.Name, time.Minute*time.Duration(h.Goods.Timeout))
	model.RDB.Set(c, h.Token+"price", h.Goods.Price, time.Minute*time.Duration(h.Goods.Timeout))
	model.RDB.Set(c, h.Token+"keep", h.Goods.Keep, time.Minute*time.Duration(h.Goods.Timeout))
	resp.OK(c)
}
