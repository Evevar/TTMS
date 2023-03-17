package main

import (
	"TTMS/internal/web/rpc"
	"TTMS/pkg/consts"
	"TTMS/pkg/jwt"
)

func main() {
	rpc.InitRPC()
	jwt.InitRedis()
	r := InitRouter()
	r.Run("127.0.0.1:" + consts.WebServerPort)
}
