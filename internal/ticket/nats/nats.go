package nats

import (
	"TTMS/configs/consts"
	"TTMS/internal/ticket/dao"
	"TTMS/internal/ticket/redis"
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"strconv"
	"strings"
	"time"
)

var NC *nats.Conn
var JS nats.JetStreamContext

func Init() {
	// 连接到nats的服务器
	conn, err := nats.Connect(consts.NatsAddress)
	NC = conn
	if err != nil {
		log.Panic("1", err)
	}

	// 初始化JetStream功能
	JS, err = conn.JetStream(nats.Context(context.Background()))
	if err != nil {
		log.Panic("2", err)
	}
	//streamName, subject, subject1 := "stream", "order.buy", "order.return"
	streamName, subject, subject1, subject2, subject3 := "stream", "order.buy", "order.return", "ticket.commit", "ticket.timeout"
	// 判断Stream是否存在，如果不存在，那么需要创建这个Stream，否则会导致pub/sub失败
	stream, err := JS.StreamInfo(streamName)
	if err != nil {
		log.Println("3", err) // 如果不存在，这里会有报错
	}
	if stream == nil {
		_, err = JS.AddStream(&nats.StreamConfig{
			Name: streamName,
			Subjects: []string{fmt.Sprintf("%s.%s", streamName, subject), fmt.Sprintf("%s.%s", streamName, subject1),
				fmt.Sprintf("%s.%s", streamName, subject2), fmt.Sprintf("%s.%s", streamName, subject3)},
			MaxAge: 3 * 24 * time.Hour,
		})
		if err != nil {
			log.Panicln("4", err)
		}
	}
	consumer, err := JS.AddConsumer(streamName, &nats.ConsumerConfig{Durable: "TicketConsumer"})
	fmt.Println("consumer = ", consumer)
	sub, err := JS.PullSubscribe(fmt.Sprintf("%s.%s", streamName, subject), "TicketConsumer") //买票
	for {
		msgs, _ := sub.Fetch(10)
		for _, msg := range msgs {
			switch msg.Subject {
			case "stream.ticket.commit":
				commitTicketHandler(msg)
			case "stream.ticket.timeout":
				timeoutTicketHandler(msg)
			}
		}
	}
	//select {}
}
func commitTicketHandler(msg *nats.Msg) {
	data := strings.Split(string(msg.Data), ";")
	fmt.Println(data)
	d0, _ := strconv.Atoi(data[0])
	d1, _ := strconv.Atoi(data[1])
	d2, _ := strconv.Atoi(data[2])
	err := redis.CommitTicket(context.Background(), fmt.Sprintf("%d;%d;%d", d0, d1, d2))
	if err != nil {
		log.Panicln(err)
	}
	dao.CommitTicket(context.Background(), d0, d1, d2)

}
func timeoutTicketHandler(msg *nats.Msg) {
	data := strings.Split(string(msg.Data), ";")
	fmt.Println(data)
	d0, _ := strconv.Atoi(data[0])
	d1, _ := strconv.Atoi(data[1])
	d2, _ := strconv.Atoi(data[2])
	err := redis.ReturnTicket(context.Background(), fmt.Sprintf("%d;%d;%d", d0, d1, d2))
	if err != nil {
		log.Panicln(err)
	}
	dao.ReturnTicket(context.Background(), d0, d1, d2)
}
