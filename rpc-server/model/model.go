package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identify int8   `json:"identify"`
	Name     string `json:"name"`
	Passwd   string `json:"passwd"`
	Token    string `json:"token"`
}

type Hoster struct {
	gorm.Model
	Identify int8   `json:"identify"`
	Name     string `json:"name"`
	Passwd   string `json:"passwd"`
	Token    string `json:"token"`
	Flag     string `json:"flag"`
}

var HosterFlag = "wozhendeshishangjia" //商家专属Flag
