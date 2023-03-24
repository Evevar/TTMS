package service

import (
	"TTMS/internal/studio/dao"
	"TTMS/kitex_gen/studio"
	"context"
)

func AddStudioService(ctx context.Context, req *studio.AddStudioRequest) (resp *studio.AddStudioResponse, err error) {
	s, err := dao.AddStudio(ctx, req.Name, req.RowsCount, req.ColsCount)
	resp = &studio.AddStudioResponse{BaseResp: &studio.BaseResp{}}
	if err == nil {
		//全部位置默认都放上座位
		err = dao.AddBatchSeat(ctx, s)
	}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func GetAllStudioService(ctx context.Context, req *studio.GetAllStudioRequest) (resp *studio.GetAllStudioResponse, err error) {
	studios, err := dao.GetAllStudio(ctx, int(req.Current), int(req.PageSize))
	resp = &studio.GetAllStudioResponse{BaseResp: &studio.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	resp.List = studios
	return resp, nil
}

func UpdateStudioService(ctx context.Context, req *studio.UpdateStudioRequest) (resp *studio.UpdateStudioResponse, err error) {
	err = dao.UpdateStudio(ctx, req.Id, req.RowsCount, req.ColsCount, req.Name)
	resp = &studio.UpdateStudioResponse{BaseResp: &studio.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func DeleteStudioService(ctx context.Context, req *studio.DeleteStudioRequest) (resp *studio.DeleteStudioResponse, err error) {
	err = dao.DeleteStudio(ctx, req.Id)
	resp = &studio.DeleteStudioResponse{BaseResp: &studio.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
