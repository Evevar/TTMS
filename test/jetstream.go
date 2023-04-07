package main

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	// 连接到nats的服务器
	conn, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Panic("1", err)
	}
	defer conn.Close()

	// 初始化JetStream功能
	js, err := conn.JetStream(nats.Context(context.Background()))
	if err != nil {
		log.Panic("2", err)
	}
	streamName, subject := "stream", "order"
	// 判断Stream是否存在，如果不存在，那么需要创建这个Stream，否则会导致pub/sub失败
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Println("3", err) // 如果不存在，这里会有报错
	}
	if stream == nil {
		log.Printf("creating stream %q and subject %q", streamName, subject)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{subject},
			MaxAge:   3 * 24 * time.Hour,
		})
		if err != nil {
			log.Panicln("4", err)
		}
	}

	// 订阅消息
	sub, err := js.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Println("subscribe msg ", msg)
		err := msg.Ack()
		if err != nil {
			return
		}
	}, nats.AckAll(), nats.DeliverNew())
	if err != nil {
		log.Panic("5", err)
		return
	}
	defer sub.Unsubscribe()

	// 发送消息
	js.Publish(subject, []byte("Hello World! "+time.Now().Format(time.RFC3339)))

	time.Sleep(5 * time.Second)
	log.Println("Exiting...")

}
