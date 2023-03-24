package dao

import (
	"TTMS/configs/consts"
	"TTMS/kitex_gen/user"
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
	err = DB.AutoMigrate(&user.User{})
	if err != nil {
		panic(err)
	}
}
