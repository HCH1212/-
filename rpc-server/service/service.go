package service

import (
	"context"
	"errors"
	"rpc-server/dao"
	"rpc-server/model"
	"rpc-server/rpc/kitex_gen/userhoster"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// 实现用户和商家的注册登录和Token验证
func Register(u *userhoster.User, h *userhoster.Hoster) (err error) {
	if u.Name != "" {
		passwdByte, passwdErr := bcrypt.GenerateFromPassword([]byte(u.Passwd), bcrypt.DefaultCost)
		if passwdErr != nil {
			return passwdErr
		}
		u.Passwd = string(passwdByte)
		res := dao.DB.Create(&u)
		if res.Error != nil {
			return res.Error
		}
		dao.RDB.Set(context.Background(), "username", u.Name, 0)
		dao.RDB.Set(context.Background(), "userpasswd", u.Passwd, 0)
	} else if h.Name != "" {
		passwdByte, passwdErr := bcrypt.GenerateFromPassword([]byte(h.Passwd), bcrypt.DefaultCost)
		if passwdErr != nil {
			return passwdErr
		}
		h.Passwd = string(passwdByte)
		//再将商家专属flag存入
		h.Flag = model.HosterFlag
		res := dao.DB.Create(&h)
		if res.Error != nil {
			return res.Error
		}
		dao.RDB.Set(context.Background(), "hostername", h.Name, 0)
		dao.RDB.Set(context.Background(), "hosterpasswd", h.Passwd, 0)
		dao.RDB.Set(context.Background(), "flag", h.Flag, 0)
	}
	return nil
}

// 只设置Token，但并没有写入
func SetToken(u *userhoster.User, h *userhoster.Hoster) (s string, err error) {
	var cnt string = time.Now().Format("2006-01-02 15:04:05")
	key := []byte(cnt)
	if u.Name != "" {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name":   u.Name,
			"passwd": u.Passwd,
		})
		s, err = claims.SignedString(key)
	} else if h.Name != "" {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name":   h.Name,
			"passwd": h.Passwd,
		})
		s, err = claims.SignedString(key)
	}
	return
}

// 将token写入数据库和缓存
func WriteToken(u *userhoster.User, h *userhoster.Hoster, s string) (err error) {
	if u.Name != "" {
		res := dao.DB.Model(&u).Where("name = ?", u.Name).Update("token", s)
		if res.Error != nil {
			return res.Error
		}
		dao.RDB.Set(context.Background(), "usertoken", s, 0)
	} else if h.Name != "" {
		res := dao.DB.Model(&h).Where("name = ?", h.Name).Update("token", s)
		if res.Error != nil {
			return res.Error
		}
		dao.RDB.Set(context.Background(), "hostertoken", s, 0)
	}
	return nil
}

// 通过redis查找Token以快速登录
func Login(u *userhoster.User, h *userhoster.Hoster) (err error) {
	if u.Name != "" {
		token, err := dao.RDB.Get(context.Background(), "usertoken").Result()
		if err != nil {
			return err
		}
		if token == u.Token {
			return nil
		} else {
			return errors.New("user Token错误")
		}
	} else if h.Name != "" {
		token, err := dao.RDB.Get(context.Background(), "hostertoken").Result()
		if err != nil {
			return err
		}
		if token == h.Token {
			return nil
		} else {
			return errors.New("hoster Token错误")
		}
	}
	return nil
}
