package rpc

import (
	"TTMS/configs/consts"
	"TTMS/kitex_gen/studio"
	"TTMS/kitex_gen/studio/studioservice"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var studioClient studioservice.Client

func InitStudioRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := studioservice.NewClient(
		consts.StudioServiceName,
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	studioClient = c
}

func AddStudio(ctx context.Context, req *studio.AddStudioRequest) (*studio.AddStudioResponse, error) {
	resp, err := studioClient.AddStudio(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func GetAllStudio(ctx context.Context, req *studio.GetAllStudioRequest) (*studio.GetAllStudioResponse, error) {
	resp, err := studioClient.GetAllStudio(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func GetStudio(ctx context.Context, req *studio.GetStudioRequest) (*studio.GetStudioResponse, error) {
	resp, err := studioClient.GetStudio(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func UpdateStudio(ctx context.Context, req *studio.UpdateStudioRequest) (*studio.UpdateStudioResponse, error) {
	resp, err := studioClient.UpdateStudio(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func DeleteStudio(ctx context.Context, req *studio.DeleteStudioRequest) (*studio.DeleteStudioResponse, error) {
	resp, err := studioClient.DeleteStudio(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
