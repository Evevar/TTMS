package redis

import (
	"TTMS/configs/consts"
	"TTMS/internal/ticket/dao"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var redisClient *redis.Client

func Init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6378",
		Password: consts.RedisPassword,
		DB:       1,
	})
	err := InitAllTicket(context.Background())
	if err != nil {
		panic(err)
	}
}

func TicketIsExist(key string) (bool, error) {
	fmt.Println("key = ", key)
	value, err := redisClient.Get(context.Background(), key).Result()
	fmt.Println("value = ", value, " err = ", err)
	if err == redis.Nil {
		return false, errors.New("票信息错误，不存在该票")
	}
	if value == "0" { //0-待售（未被预定）
		return true, nil
	}
	//票已经被抢
	return false, errors.New("票已经被抢")
}

// InitAllTicket 将所有票加载到redis缓存中
func InitAllTicket(ctx context.Context) error {
	// 票信息 键值对 key="ScheduleId:SeatRow:SeatCol" value="Status"
	tickets, err := dao.GetAllTicket2(ctx)
	if err != nil {
		log.Panicln(err)
		return err
	}
	for _, ticket := range tickets {
		redisClient.Set(ctx, fmt.Sprintf("%d:%d:%d", ticket.ScheduleId, ticket.SeatRow, ticket.SeatCol), ticket.Status, 0)
	}
	return nil
}
func BuyTicket(ctx context.Context, key string) error {
	return redisClient.Set(ctx, key, "9", 0).Err()
}
