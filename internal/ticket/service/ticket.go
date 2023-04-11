package service

import (
	"TTMS/configs/consts"
	"TTMS/internal/ticket/dao"
	"TTMS/internal/ticket/redis"
	"TTMS/kitex_gen/ticket"
	"context"
	"errors"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

var nc *nats.Conn
var js nats.JetStreamContext
var buyPublisher nats.PubAckFuture

func NatsInit() {
	// 连接到nats的服务器
	conn, err := nats.Connect(consts.NatsAddress)
	nc = conn
	if err != nil {
		log.Panic("1", err)
	}

	// 初始化JetStream功能
	js, err = conn.JetStream(nats.Context(context.Background()))
	if err != nil {
		log.Panic("2", err)
	}
	//streamName, subject, subject1 := "stream", "order.buy", "order.return"
	streamName, subject := "stream", "order.buy"
	// 判断Stream是否存在，如果不存在，那么需要创建这个Stream，否则会导致pub/sub失败
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Println("3", err) // 如果不存在，这里会有报错
	}
	if stream == nil {
		log.Printf("creating stream %q and subject %q", streamName, subject)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{"stream.>"},
			MaxAge:   3 * 24 * time.Hour,
		})
		if err != nil {
			log.Panicln("4", err)
		}
	}
	//select {}
}

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
	key := fmt.Sprintf("%d:%d:%d", req.ScheduleId, req.SeatRow, req.SeatCol)
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
	fmt.Println(1, " ", nc, " ", nc.IsClosed())
	//暂时不改变票的状态
	//err = redis.BuyTicket(ctx, key)
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
		return resp, err
	}
	//发送创建订单消息
	fmt.Println(2, " ", nc, " ", nc.IsClosed())
	_, err = js.Publish("stream.order.buy",
		[]byte(fmt.Sprintf("%d:%s:%s", req.UserId, key, time.Now().Format("2006-01-02 15:04:05"))))
	fmt.Println(err)
	if err != nil {
		log.Println(err)
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	}
	//go dao.BuyTicket(ctx, int(req.ScheduleId), int(req.SeatRow), int(req.SeatCol))
	fmt.Println(3)
	return resp, nil
}

func ReturnTicket(ctx context.Context, req *ticket.ReturnTicketRequest) (resp *ticket.ReturnTicketResponse, err error) {

	return
}
