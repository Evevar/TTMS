package service

/*
+---------------------------------------------------------------------------+---------------------------------------------------------------+
|							BuyTicketMsg									|						ReturnTicketMsg							|
|---------------------------------------------------------------------------+---------------------------------------------------------------+
|	UserId	|	ScheduleId	|	SeatRow	|	SeatCol	|	Time	|	Price	|	UserId	|	ScheduleId	|	SeatRow	|	SeatCol	|	Time	|
+---------------------------------------------------------------------------+---------------------------------------------------------------+
*/
import (
	"TTMS/configs/consts"
	"TTMS/internal/ticket/dao"
	"TTMS/internal/ticket/nats"
	"TTMS/internal/ticket/redis"
	"TTMS/kitex_gen/order"
	"TTMS/kitex_gen/order/orderservice"
	"TTMS/kitex_gen/play"
	"TTMS/kitex_gen/play/playservice"
	"TTMS/kitex_gen/ticket"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"time"
)

var playClient playservice.Client
var orderClient orderservice.Client
var Loc *time.Location

func OrderPlayRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := orderservice.NewClient(
		consts.OrderServiceName,
		//client.WithMiddleware(mw.CommonMiddleware),
		//client.WithInstanceMW(mw.ClientMiddleware),
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
	orderClient = c
}
func InitPlayRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := playservice.NewClient(
		consts.PlayServiceName,
		//client.WithMiddleware(mw.CommonMiddleware),
		//client.WithInstanceMW(mw.ClientMiddleware),
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
	playClient = c
}
func LoadLocation() {
	var err error
	Loc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println("Error loading location:", err)
		return
	}
}
func BatchAddTicketService(ctx context.Context, req *ticket.BatchAddTicketRequest) (resp *ticket.BatchAddTicketResponse, err error) {
	//fmt.Println(req.ScheduleId, req.Price, req.PlayName, req.StudioId, req.List)
	err = dao.BatchAddTicket(ctx, req.ScheduleId, req.Price, req.PlayName, req.StudioId, req.List)
	for _, s := range req.List {
		redis.AddTicket(int(req.ScheduleId), int(s.Row), int(s.Col), req.Price)
	}
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
func TicketIsExist(ctx context.Context, ScheduleId int64, SeatRow int32, SeatCol int32) (bool, error, string) {
	key := fmt.Sprintf("%d;%d;%d", ScheduleId, SeatRow, SeatCol)
	result, err := redis.TicketIsExist(key) //只查看票的状态，不抢票
	if result {
		return true, nil, "redis"
	}
	if err == nil { //redis没找到，去mysql再找一下
		t := dao.GetTicket(ctx, ScheduleId, SeatRow, SeatCol)
		if t.Id > 0 && t.Status == 0 {
			return true, nil, "mysql"
		}
	}
	return false, errors.New("票已被抢"), ""
}
func BuyTicket(ctx context.Context, ScheduleId int64, SeatRow int32, SeatCol int32, source string) {
	if source == "redis" {
		redis.BuyTicket(ctx, fmt.Sprintf("%d;%d;%d", ScheduleId, SeatRow, SeatCol))
	}
	//source=="mysql"，无论是否更新redis，mysql是一定要更新的
	go dao.BuyTicket(ctx, ScheduleId, SeatRow, SeatCol)

}
func BuyTicketService(ctx context.Context, req *ticket.BuyTicketRequest) (resp *ticket.BuyTicketResponse, err error) {
	resp = &ticket.BuyTicketResponse{BaseResp: &ticket.BaseResp{}}
	//查看票是否还存在（没有被别人买）
	key := fmt.Sprintf("%d;%d;%d", req.ScheduleId, req.SeatRow, req.SeatCol)

	result, err, source := TicketIsExist(ctx, req.ScheduleId, req.SeatRow, req.SeatCol) //只查看票的状态，不抢票
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
		return resp, nil
	}

	//票已经被买,或者是同时抢票但没有抢到分布式锁
	if !result || !redis.AcquireLock(fmt.Sprintf("lock;%s", key)) {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = errors.New("票已经被买").Error()
		return resp, nil
	}
	defer redis.ReleaseLock(fmt.Sprintf("lock;%s", key)) //释放锁

	//判断是否为有效的’买票时间‘,多人抢票时，只让抢到分布式锁的用户进行时间检查
	schedule, err := playClient.GetSchedule(ctx, &play.GetScheduleRequest{Id: req.ScheduleId})

	deadline, _ := time.Parse("2006-01-02 15:04:05", schedule.Schedule.ShowTime)
	deadline = deadline.In(Loc).Add(-8 * time.Hour)

	log.Println("now = ", time.Now().Format("2006-01-02 15:04:05"))
	log.Println("showtime = ", deadline)
	log.Println("until = ", time.Until(deadline))
	if time.Until(deadline) < 10*time.Minute { //距离开场已经不足10分钟，停止售票
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = errors.New("已停止售票").Error()
		return resp, nil
	}

	//抢到分布式锁,执行买票流程
	//选择是否更新redis，并强制更新mysql
	BuyTicket(ctx, req.ScheduleId, req.SeatRow, req.SeatCol, source)

	//成功抢到票,发送创建订单消息
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("time = ", t)
	_, err = nats.JS.Publish("stream.order.buy",
		[]byte(fmt.Sprintf("%d;%s;%s;%s", req.UserId, key, t,
			redis.GetTicketPrice(ctx, fmt.Sprintf("%d;price", req.ScheduleId)))))
	fmt.Println(err)
	if err != nil {
		log.Println(err)
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func ReturnTicket(ctx context.Context, ScheduleId int64, SeatRow int32, SeatCol int32) {
	if ttl := redis.TicketTTL(ctx, fmt.Sprintf("%d;%d;%d", ScheduleId, SeatRow, SeatCol)); ttl > 0 {
		redis.ReturnTicket(ctx, fmt.Sprintf("%d;%d;%d", ScheduleId, SeatRow, SeatCol), ttl)
	}
	go dao.ReturnTicket(ctx, ScheduleId, SeatRow, SeatCol)
}
func ReturnTicketService(ctx context.Context, req *ticket.ReturnTicketRequest) (resp *ticket.ReturnTicketResponse, err error) {
	//先判断是否为有效的’退票时间‘
	schedule, err := playClient.GetSchedule(ctx, &play.GetScheduleRequest{Id: req.ScheduleId})
	resp = &ticket.ReturnTicketResponse{BaseResp: &ticket.BaseResp{}}

	deadline, _ := time.Parse("2006-01-02 15:04:05", schedule.Schedule.ShowTime)
	deadline = deadline.In(Loc).Add(-8 * time.Hour)
	if time.Until(deadline) < 0 { //演出已经开始，停止退票
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = errors.New("已停止退票").Error()
		return resp, nil
	}

	resp1, _ := orderClient.UpdateOrder(ctx, &order.UpdateOrderRequest{UserId: req.UserId, ScheduleId: req.ScheduleId,
		SeatRow: req.SeatRow, SeatCol: req.SeatCol})
	if resp1.BaseResp.StatusCode == 1 {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = resp1.BaseResp.StatusMessage
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	ReturnTicket(ctx, req.ScheduleId, req.SeatRow, req.SeatCol)
	return resp, nil
}
