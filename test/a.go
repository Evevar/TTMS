package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// 加载 CST 时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	t, err := time.Parse("2006-01-02 15:04:05", "2023-06-05 14:30:00")
	t = t.In(loc).Add(-8 * time.Hour)
	time.Until(t)
	log.Println("show_time = ", t)
	log.Println("now = ", time.Now().In(loc))
	log.Println(time.Until(t))
	log.Println(err)
}
