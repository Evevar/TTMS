package service

import (
	"TTMS/configs/consts"
	"TTMS/internal/order/dao"
	"TTMS/kitex_gen/order"
	"TTMS/kitex_gen/play"
	"TTMS/kitex_gen/play/playservice"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/nats-io/nats.go"
	"log"
	"strconv"
	"strings"
	"time"
)

var nc *nats.Conn
var playClient playservice.Client

func InitPlayRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := playservice.NewClient(
		consts.PlayServiceName,
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
	playClient = c
}
func NatsInit() {
	// 连接到nats的服务器
	var err error
	nc, err = nats.Connect(consts.NatsAddress)
	if err != nil {
		log.Panic("1", err)
	}

	// 初始化JetStream功能
	js, err := nc.JetStream(nats.Context(context.Background()))
	if err != nil {
		log.Panic("2", err)
	}
	//streamName, subject, subject1 := "stream", "order.buy", "order.return"
	streamName, subject := "stream", "order.buy"
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Println("3", err) // 如果不存在，这里会有报错
	}
	if stream == nil {
		//fmt.Sprintf("%s.%s", streamName, subject), fmt.Sprintf("%s.%s", streamName, subject1)
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
	consumer, err := js.AddConsumer(streamName, &nats.ConsumerConfig{Durable: "buyConsumer"})
	fmt.Println("consumer = ", consumer)
	sub, err := js.PullSubscribe(fmt.Sprintf("%s.%s", streamName, subject), "buyConsumer") //买票
	//fmt.Println("sub = ", sub, " err = ", err)
	//sub1, err := js.PullSubscribe(fmt.Sprintf("%s.%s", streamName, subject1), "returnConsumer") //退票
	////fmt.Println("sub1 = ", sub1, " err = ", err)

	for {
		msgs, _ := sub.Fetch(10)
		for _, msg := range msgs {
			switch msg.Subject {
			case "stream.order.buy":
				AddOrderHandler(msg)
			case "stream.order.return":
				ReturnOrderHandler(msg)
			}
		}
	}
}

func AddOrderHandler(msg *nats.Msg) {
	data := strings.Split(string(msg.Data), ";")
	fmt.Println(data)
	d0, _ := strconv.Atoi(data[0])
	d1, _ := strconv.Atoi(data[1])
	d2, _ := strconv.Atoi(data[2])
	d3, _ := strconv.Atoi(data[3])
	d5, _ := strconv.Atoi(data[5])
	//fmt.Println(d0, d1, d2, d3, data[4])
	err := dao.AddOrder(d0, d1, d2, d3, data[4], d5)
	if err != nil {
		log.Panicln(err)
	}
	err = msg.Ack()
	if err != nil {
		return
	}
}
func ReturnOrderHandler(msg *nats.Msg) {
	data := strings.Split(string(msg.Data), ";")
	fmt.Println(data)
	d0, _ := strconv.Atoi(data[0])
	d1, _ := strconv.Atoi(data[1])
	d2, _ := strconv.Atoi(data[2])
	d3, _ := strconv.Atoi(data[3])
	err := dao.UpdateOrder(d0, d1, d2, d3, data[4])
	if err != nil {
		log.Panicln(err)
	}
	err = msg.Ack()
	if err != nil {
		return
	}
}

func GetAllOrderService(ctx context.Context, req *order.GetAllOrderRequest) (resp *order.GetAllOrderResponse, err error) {
	resp = &order.GetAllOrderResponse{BaseResp: &order.BaseResp{}}
	resp.List, err = dao.GetAllOrder(ctx, int(req.UserId))
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	return resp, nil
}

func GetOrderAnalysisService(ctx context.Context, req *order.GetOrderAnalysisRequest) (resp *order.GetOrderAnalysisResponse, err error) {
	//通过RPC 根据playID查找scheduleId,再根据scheduleId统计票房
	fmt.Println("通过RPC 根据playID查找scheduleId,再根据scheduleId统计票房")
	resp = &order.GetOrderAnalysisResponse{BaseResp: &order.BaseResp{}}
	resp1, _ := playClient.PlayToSchedule(ctx, &play.PlayToScheduleRequest{Id: req.PlayId})
	if resp1.BaseResp.StatusCode == 1 {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = resp1.BaseResp.StatusMessage
		log.Println(resp1.BaseResp.StatusMessage)
		return resp, nil
	}
	o := &order.OrderAnalysis{
		PlayId:       resp1.Play.Id,
		PlayName:     resp1.Play.Name,
		Price:        int32(resp1.Play.Price),
		PlayArea:     resp1.Play.Area,
		PlayDuration: resp1.Play.Duration,
		StartData:    resp1.Play.StartDate,
		EndData:      resp1.Play.EndDate,
	}
	o.TotalTicket, err = dao.GetOrderAnalysis(ctx, resp1.ScheduleList)
	o.Sales = o.TotalTicket * int64(o.Price)
	resp.OrderAnalysis = o
	fmt.Println("o.TotalTicket = ", o.TotalTicket, " err = ", err)
	if err != nil {
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.BaseResp.StatusMessage = "success"
	}
	fmt.Println("resp = ", resp)
	return resp, nil
}
