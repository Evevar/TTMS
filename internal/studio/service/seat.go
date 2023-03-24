package service

import (
	"TTMS/internal/studio/dao"
	"TTMS/kitex_gen/studio"
	"context"
)

func AddSeatService(ctx context.Context, req *studio.AddSeatRequest) (*studio.AddSeatResponse, error) {
	seatInfo := &studio.Seat{
		StudioId: req.StudioId,
		Row:      req.Row,
		Col:      req.Col,
		Status:   req.Status,
	}
	err := dao.AddSeat(ctx, seatInfo)
	resp := &studio.AddSeatResponse{BaseResp: &studio.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}
