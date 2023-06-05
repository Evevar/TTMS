package main

import (
	"TTMS/configs/consts"
	"TTMS/kitex_gen/order"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(consts.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,                                //禁用默认事务操作
			Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句
		},
	)
	if err != nil {
		panic(err)
	}
}

func main() {
	Init()
	o := order.Order{}
	DB.Model(&o).Where("user_id = ? and schedule_id = ? and seat_row = ? and seat_col = ?",
		1, 3, 3, 4).Order("id desc").Find(&o)
	fmt.Print(o.Id)
}
