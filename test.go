package main

import (
	"fmt"
	"time"
)

var ticketId2Timer map[int64]time.Timer

func main() {
	ticketId2Timer = make(map[int64]time.Timer)
	var ticket int64 = 1
	f1(ticket, time.Now())
	time.Sleep(time.Second * 5) //3秒模拟用户3分钟之后支付
	t := ticketId2Timer[ticket]
	t.Stop()
}
func f1(ticket int64, buy time.Time) { //buy 下单时间
	fmt.Println(buy.String())
	time.Sleep(time.Second * 1) //模拟异步执行创建订单所消耗的时间
	now := time.Now()
	dur := now.Sub(buy)                      //相对于抢票时间推迟10分钟作为最晚下单时间
	t := time.NewTimer(time.Second*10 - dur) //10秒相当于等待用户付款的10分钟
	ticketId2Timer[ticket] = *t
	t.Stop()
	select {
	case <-t.C: //距离抢票时间已经过去了10分钟，如果用户还未付款，则删除订单
		//模拟删除订单
		now = time.Now()
		fmt.Println("now = ", now)
		fmt.Println("buy = ", buy)
		fmt.Println("duration = ", now.Sub(buy))
	}

}
