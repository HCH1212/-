package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Good struct {
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Keep    int     `json:"keep"`
	Timeout int     `json:"timeout"` //商品过期时间,单位为分钟
}

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
	Token  string `json:"token"`
	Book   Good   `json:"book"`
	IsBuy  bool   `json:"isbuy"`
}

type Hoster struct {
	gorm.Model
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
	Token  string `json:"token"`
	Flag   string `json:"flag"`
	Goods  Good   `json:"goods"`
}

var RDB *redis.Client

var HosterFlag = "wozhendeshishangjia" //商家专属Flag
