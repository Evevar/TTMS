package main

import (
	"TTMS/internal/web/rpc"
	"TTMS/pkg/consts"
)

func main() {
	rpc.InitRPC()
	r := InitRouter()
	r.Run("127.0.0.1:" + consts.WebServerPort)
}
