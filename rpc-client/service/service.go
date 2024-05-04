package service

import (
	"context"
	"fmt"
	"log"
	"rpc-client/model"
	"rpc-client/rpc"
	"rpc-client/rpc/kitex_gen/userhoster"
)

func UserRegister(u model.User) (token string, err error) {
	err = rpc.UserHoster.Register(context.Background(), &userhoster.User{
		Name:   u.Name,
		Passwd: u.Passwd,
	}, &userhoster.Hoster{
		Name: "",
	})
	if err != nil {
		log.Fatalln("register rpc user error :", err)
	}
	//注册成功就自动设置Token
	token, err = rpc.UserHoster.SetToken(context.Background(), &userhoster.User{
		Name:   u.Name,
		Passwd: u.Passwd,
	}, &userhoster.Hoster{
		Name: "",
	})
	if err != nil {
		fmt.Println(err)
		log.Fatalln("user set token error :", err)
	}
	return
}

func HosterRegister(h model.Hoster) (token string, err error) {
	err = rpc.UserHoster.Register(context.Background(), &userhoster.User{
		Name: "",
	}, &userhoster.Hoster{
		Name:   h.Name,
		Passwd: h.Passwd,
	})
	if err != nil {
		log.Fatalln("register rpc hoster error :", err)
	}
	//注册成功就自动设置Token
	token, err = rpc.UserHoster.SetToken(context.Background(), &userhoster.User{
		Name: "",
	}, &userhoster.Hoster{
		Name:   h.Name,
		Passwd: h.Passwd,
	})
	if err != nil {
		log.Fatalln("hoster set token error :", err)
	}
	return
}

func UserLogin(u model.User) (err error) {
	err = rpc.UserHoster.Login(context.Background(), &userhoster.User{
		Token: u.Token,
	}, &userhoster.Hoster{
		Name: "",
	})
	if err != nil {
		log.Fatalln("user login error :", err)
	}
	return
}

func HosterLogin(h model.Hoster) (err error) {
	err = rpc.UserHoster.Register(context.Background(), &userhoster.User{
		Name: "",
	}, &userhoster.Hoster{
		Token: h.Token,
	})
	if err != nil {
		log.Fatalln("hoster login error :", err)
	}
	return
}
