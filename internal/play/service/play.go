package service

import (
	"TTMS/internal/play/dao"
	"TTMS/kitex_gen/play"
	"context"
)

func AddPlayService(ctx context.Context, req *play.AddPlayRequest) (resp *play.AddPlayResponse, err error) {
	PlayInfo := &play.Play{Name: req.Name, Type: req.Type, Area: req.Area,
		Rating: req.Rating, Duration: req.Duration, StartData: req.StartData, EndData: req.EndData, Price: req.Price}
	err = dao.AddPlay(ctx, PlayInfo)
	resp = &play.AddPlayResponse{BaseResp: &play.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func UpdatePlayService(ctx context.Context, req *play.UpdatePlayRequest) (resp *play.UpdatePlayResponse, err error) {
	PlayInfo := &play.Play{Id: req.Id, Name: req.Name, Type: req.Type, Area: req.Area,
		Rating: req.Rating, Duration: req.Duration, StartData: req.StartData, EndData: req.EndData, Price: req.Price}
	err = dao.UpdatePlay(ctx, PlayInfo)
	resp = &play.UpdatePlayResponse{BaseResp: &play.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func DeletePlayService(ctx context.Context, req *play.DeletePlayRequest) (resp *play.DeletePlayResponse, err error) {
	err = dao.DeletePlay(ctx, req.Id)
	resp = &play.DeletePlayResponse{BaseResp: &play.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func GetAllPlayService(ctx context.Context, req *play.GetAllPlayRequest) (resp *play.GetAllPlayResponse, err error) {
	resp = &play.GetAllPlayResponse{BaseResp: &play.BaseResp{}}
	resp.List, err = dao.GetAllPlay(ctx, int(req.Current), int(req.PageSize))
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func AddScheduleService(ctx context.Context, req *play.AddScheduleRequest) (resp *play.AddScheduleResponse, err error) {
	SInfo := &play.Schedule{PlayId: req.PlayId, StudioId: req.StudioId, ShowTime: req.ShowTime}
	err = dao.AddSchedule(ctx, SInfo)
	resp = &play.AddScheduleResponse{BaseResp: &play.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func UpdateScheduleService(ctx context.Context, req *play.UpdateScheduleRequest) (resp *play.UpdateScheduleResponse, err error) {
	SInfo := &play.Schedule{PlayId: req.PlayId, StudioId: req.StudioId, ShowTime: req.ShowTime}
	err = dao.UpdateSchedule(ctx, SInfo)
	resp = &play.UpdateScheduleResponse{BaseResp: &play.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func DeleteScheduleService(ctx context.Context, req *play.DeleteScheduleRequest) (resp *play.DeleteScheduleResponse, err error) {
	err = dao.DeleteSchedule(ctx, req.Id)
	resp = &play.DeleteScheduleResponse{BaseResp: &play.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
func GetAllScheduleService(ctx context.Context, req *play.GetAllScheduleRequest) (resp *play.GetAllScheduleResponse, err error) {
	resp = &play.GetAllScheduleResponse{BaseResp: &play.BaseResp{}}
	resp.List, err = dao.GetAllSchedule(ctx, int(req.Current), int(req.PageSize))
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
