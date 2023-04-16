package service

import (
	"TTMS/internal/ticket/dao"
	"TTMS/internal/ticket/nats"
	"TTMS/internal/ticket/redis"
	"TTMS/kitex_gen/ticket"
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func BatchAddTicketService(ctx context.Context, req *ticket.BatchAddTicketRequest) (resp *ticket.BatchAddTicketResponse, err error) {
	//fmt.Println(req.ScheduleId, req.Price, req.PlayName, req.StudioId, req.List)
	err = dao.BatchAddTicket(ctx, req.ScheduleId, req.Price, req.PlayName, req.StudioId, req.List)
	resp = &ticket.BatchAddTicketResponse{BaseResp: &ticket.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func UpdateTicketService(ctx context.Context, req *ticket.UpdateTicketRequest) (resp *ticket.UpdateTicketResponse, err error) {
	err = dao.UpdateTicket(ctx, req.ScheduleId, req.SeatRow, req.SeatCol, req.Price, req.Status)
	resp = &ticket.UpdateTicketResponse{BaseResp: &ticket.BaseResp{}}
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	fmt.Println("resp = ", resp)
	return resp, nil
}

func GetAllTicketService(ctx context.Context, req *ticket.GetAllTicketRequest) (resp *ticket.GetAllTicketResponse, err error) {
	resp = &ticket.GetAllTicketResponse{BaseResp: &ticket.BaseResp{}}
	resp.List, err = dao.GetAllTicket(ctx, req.ScheduleId)
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func BuyTicketService(ctx context.Context, req *ticket.BuyTicketRequest) (resp *ticket.BuyTicketResponse, err error) {
	key := fmt.Sprintf("%d;%d;%d", req.ScheduleId, req.SeatRow, req.SeatCol)
	resp = &ticket.BuyTicketResponse{BaseResp: &ticket.BaseResp{}}
	result, err := redis.TicketIsExist(key)
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
		return resp, nil
	}
	if !result {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = errors.New("票已经被抢").Error()
		return resp, nil
	}
	//成功买到票
	//fmt.Println(1, " ", nc, " ", nc.IsClosed())
	//暂时不改变票的状态
	err = redis.BuyTicket(ctx, key)
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
		return resp, err
	}
	//发送创建订单消息
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("time = ", t)
	_, err = nats.JS.Publish("stream.order.buy",
		[]byte(fmt.Sprintf("%d;%s;%s;%s", req.UserId, key, t,
			redis.GetTicketPrice(ctx, fmt.Sprintf("%s;price", key)))))
	fmt.Println(err)
	if err != nil {
		log.Println(err)
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	go dao.BuyTicket(ctx, int(req.ScheduleId), int(req.SeatRow), int(req.SeatCol))
	return resp, nil
}

func ReturnTicketService(ctx context.Context, req *ticket.ReturnTicketRequest) (resp *ticket.ReturnTicketResponse, err error) {
	key := fmt.Sprintf("%d;%d;%d", req.ScheduleId, req.SeatRow, req.SeatCol)
	resp = &ticket.ReturnTicketResponse{BaseResp: &ticket.BaseResp{}}
	err = redis.ReturnTicket(ctx, key)
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
		return resp, nil
	}
	//发送创建订单消息
	_, err = nats.JS.Publish("stream.order.return",
		[]byte(fmt.Sprintf("%d;%s;%s", req.UserId, key, time.Now().Format("2006-01-02 15:04:05"))))
	fmt.Println(err)
	if err != nil {
		log.Println(err)
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	go dao.ReturnTicket(ctx, int(req.ScheduleId), int(req.SeatRow), int(req.SeatCol))
	return resp, nil
}
