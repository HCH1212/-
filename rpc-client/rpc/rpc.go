package rpc

import (
	"log"
	"rpc-client/rpc/kitex_gen/userhoster/userhoster"

	"github.com/cloudwego/kitex/client"
)

var UserHoster userhoster.Client

func InitRpc() {
	u, err := userhoster.NewClient("uh", client.WithHostPorts("localhost:8888"))
	if err != nil {
		log.Fatalln("can't connent service :", err)
	}
	UserHoster = u
}
