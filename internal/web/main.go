package main

import (
	"TTMS/configs/consts"
	"TTMS/internal/web/rpc"
	"TTMS/pkg/jwt"
)

func main() {
	rpc.InitRPC()
	jwt.InitRedis()
	r := InitRouter()
	r.Run("127.0.0.1:" + consts.WebServerPort)
}
