package service

import (
	"TTMS/configs/consts"
	"TTMS/internal/order/dao"
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"strconv"
	"strings"
	"time"
)

var nc *nats.Conn
var js nats.JetStreamContext

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
	streamName, subject, subject1 := "stream", "order.buy", "order.return"
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
	fmt.Println("sub = ", sub, " err = ", err)
	sub1, err := js.PullSubscribe(fmt.Sprintf("%s.%s", streamName, subject1), "returnConsumer") //退票
	fmt.Println("sub1 = ", sub1, " err = ", err)
	go func() {
		for {
			msgs, _ := sub.Fetch(10)
			for _, msg := range msgs {
				data := strings.Split(string(msg.Data), ":")
				d0, _ := strconv.Atoi(data[0])
				d1, _ := strconv.Atoi(data[1])
				d2, _ := strconv.Atoi(data[2])
				d3, _ := strconv.Atoi(data[3])
				fmt.Println(d0, d1, d2, d3, data[4])
				err = dao.AddOrder(d0, d1, d2, d3, data[4])
				if err != nil {
					log.Panicln(err)
				}
				//publish, err := js.Publish("order.buy", []byte("success"))
				//if err != nil {
				//	log.Panicln(publish)
				//	return
				//}
				err = msg.Ack()
				if err != nil {
					return
				}
			}
		}
	}()
	for {
		msgs, _ := sub1.Fetch(10)
		for _, msg := range msgs {
			data := strings.Split(string(msg.Data), ":")
			d0, _ := strconv.Atoi(data[0])
			d1, _ := strconv.Atoi(data[1])
			d2, _ := strconv.Atoi(data[2])
			d3, _ := strconv.Atoi(data[3])
			fmt.Println(d0, d1, d2, d3, data[4])
			//err = dao.AddOrder(d0, d1, d2, d3, data[4])
			//if err != nil {
			//	log.Panicln(err)
			//}
			//之后换为具体逻辑
			err := msg.Ack()
			if err != nil {
				return
			}
		}
	}
}

//func AddOrderService(ctx context.Context) (resp *order.AddOrderResponse, err error) {
//
//	return
//}
