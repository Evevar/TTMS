package redis

import (
	"TTMS/configs/consts"
	"TTMS/internal/ticket/dao"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var redisClient *redis.Client

func Init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6378",
		Password: consts.RedisPassword,
		DB:       1,
	})
	//err := InitAllTicket(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	go LockRenewal()
}

// InitAllTicket 将所有票加载到redis缓存中（已弃用）
func InitAllTicket(ctx context.Context) error {
	// 票信息 键值对 key="ScheduleId:SeatRow:SeatCol" value="Status"
	tickets, err := dao.GetAllTicket2(ctx)
	if err != nil {
		log.Panicln(err)
		return err
	}
	for _, ticket := range tickets {
		redisClient.Set(ctx, fmt.Sprintf("%d;%d;%d", ticket.ScheduleId, ticket.SeatRow, ticket.SeatCol), ticket.Status, 0)
		redisClient.Set(ctx, fmt.Sprintf("%d;%d;%d;price", ticket.ScheduleId, ticket.SeatRow, ticket.SeatCol), ticket.Price, 0)
	}
	return nil
}

// AddTicket 添加票缓存
func AddTicket(key string, price int32) {
	ctx := context.Background()
	redisClient.Set(ctx, key, 0, time.Hour)
	redisClient.Set(ctx, fmt.Sprintf("%s;price", key), price, time.Hour)
}

// AcquireLock 分布式锁，加锁
func AcquireLock(lockKey string) bool {
	//result, err := redis.String(conn.Do("SET", lockKey, lockValue, "NX", "EX", lockTimeout))
	result, err := redisClient.SetNX(context.Background(), lockKey, 1, time.Duration(consts.RedisLockTimeOut)).Result()
	if err != nil || !result {
		return false
	}
	return true
}

// ReleaseLock 分布式锁，释放锁
func ReleaseLock(lockKey string) bool {
	result, err := redisClient.Del(context.Background(), lockKey).Result()
	if err != nil || result != 1 {
		return false
	}
	return true
}

// LockRenewal 为分布式锁续期
func LockRenewal() {
	var cursor uint64 = 0
	ctx := context.Background()
	for range time.Tick(1 * time.Second) {
		keys, next, err := redisClient.Scan(ctx, cursor, "lock;*", 10000).Result()
		if err != nil {
			log.Println(err)
		}
		cursor = next
		for _, key := range keys {
			d, err := redisClient.TTL(ctx, key).Result()
			if err != nil {
				log.Println(err)
			}
			if d < 2*time.Second { //锁过期时间不足2s时，对锁进行续期
				redisClient.Expire(ctx, key, consts.RedisLockTimeOut)
			}
		}
	}

}
func TicketIsExist(key string) (bool, error) {
	fmt.Println("key = ", key)
	value, err := redisClient.Get(context.Background(), key).Result()
	fmt.Println("value = ", value, " err = ", err)
	if err == redis.Nil {
		return false, nil
	}
	if value == "0" { //0-待售（未被预定）
		return true, nil
	}
	//票已经被抢
	return false, errors.New("票已经被抢")
}

func BuyTicket(ctx context.Context, key string) {
	redisClient.Set(ctx, key, "9", time.Hour)
}
func ReturnTicket(ctx context.Context, key string, expire time.Duration) {
	redisClient.Set(ctx, key, "0", expire)
}
func CommitTicket(ctx context.Context, key string) {
	redisClient.Set(ctx, key, "1", time.Hour)
}
func GetTicketPrice(ctx context.Context, key string) string {
	price, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		log.Println("GetTicketPrice ", err)
	}
	return price
}
func TicketTTL(ctx context.Context, key string) time.Duration {
	d, err := redisClient.TTL(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}
	return d
}
