package api

import (
	"fmt"
	"rpc-client/model"
	"rpc-client/resp"
	"rpc-client/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var u model.User
	u.Name = c.PostForm("name")
	u.Passwd = c.PostForm("passwd")
	var err error
	if err = c.ShouldBind(&u); err != nil {
		fmt.Println(err)
		resp.ParamError(c)
		return
	}
	var token string
	if token, err = service.UserRegister(u); err != nil {
		resp.InternalError(c)
		return
	}
	resp.OKWithData(c, token)
}

func HosterRegister(c *gin.Context) {
	var h model.Hoster
	h.Name = c.PostForm("name")
	h.Passwd = c.PostForm("passwd")
	var err error
	if err = c.ShouldBind(&h); err != nil {
		resp.ParamError(c)
		return
	}
	var token string
	if token, err = service.HosterRegister(h); err != nil {
		resp.InternalError(c)
		return
	}
	resp.OKWithData(c, token)
}

func UserLogin(c *gin.Context) {
	var u model.User
	u.Token = c.PostForm("token")
	if err := c.ShouldBind(&u); err != nil {
		resp.ParamError(c)
		return
	}
	if err := service.UserLogin(u); err != nil {
		resp.InternalError(c)
		return
	}
	resp.OK(c)
}

func HosterLogin(c *gin.Context) {
	var h model.Hoster
	h.Token = c.PostForm("token")
	if err := c.ShouldBind(&h); err != nil {
		resp.ParamError(c)
		return
	}
	if err := service.HosterLogin(h); err != nil {
		resp.InternalError(c)
		return
	}
	resp.OK(c)
}
