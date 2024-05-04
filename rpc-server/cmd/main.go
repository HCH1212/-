package main

import (
	"log"
	_ "rpc-server/dao"
	"rpc-server/rpc"
	"rpc-server/rpc/kitex_gen/userhoster/userhoster"
)

func main() {
	svr := userhoster.NewServer(new(rpc.UserhosterImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
