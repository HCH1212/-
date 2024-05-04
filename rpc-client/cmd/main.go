package main

import (
	"rpc-client/redis"
	"rpc-client/router"
	"rpc-client/rpc"
)

func main() {
	redis.InitRedis()
	rpc.InitRpc()
	router.InitRouters()
}
