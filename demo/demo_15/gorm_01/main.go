package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TBBank struct {
	ID uint
	Name string
	balance int
}

func main() {
	dsn := "root:Ycu061036@tcp(docker.cloaks.cn:3306)/db_test?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Silent,
			Colorful: true,
		},
	)

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 自动建表
	// err = db.AutoMigrate(&TBBank{})
	// if err != nil {
		// panic(err)
	// }
}