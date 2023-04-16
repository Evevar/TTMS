package main

import (
	"TTMS/configs/consts"
	"TTMS/internal/play/dao"
	"TTMS/internal/play/service"
	play "TTMS/kitex_gen/play/playservice"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:10003")
	if err != nil {
		panic(err)
	}
	svr := play.NewServer(new(PlayServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.PlayServiceName}), // server name
		//server.WithMiddleware(mw.CommonMiddleware),                                          // middleWare
		//server.WithMiddleware(mw.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithRegistry(r),                                             // registry
	)
	dao.Init()
	service.InitStudioRPC()
	service.InitTicketRPC()
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
