package consts

import "time"

const (
	EtcdAddress       = "127.0.0.1:2379"
	NatsAddress       = "nats://127.0.0.1:4222"
	UserServiceName   = "userSvr"
	StudioServiceName = "studioSvr"
	PlayServiceName   = "playSvr"
	TicketServiceName = "ticketSvr"
	OrderServiceName  = "orderSvr"
	MySQLDefaultDSN   = "TTMS:TTMS@tcp(localhost:3307)/TTMS?charset=utf8mb4&parseTime=True"
	WebServerPort     = "8080"
	RedisPassword     = "redis"
	OrderDelayTime    = time.Minute * 1
	RedisLockTimeOut  = time.Second * 5
)
