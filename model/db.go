package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USERNAME     = "demo1"
	PASSWORD     = "demo123"
	NETWORK      = "tcp"
	SERVER       = "127.0.0.1"
	PORT         = 3306
	DATABASE     = "demo"
	MaxLifetime  = 10
	MaxOpenConns = 10
	MaxIdleConns = 10
)

var db *gorm.DB

func init() {
	addr := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

	conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}

	//設定ConnMaxLifetime/MaxIdleConns/MaxOpenConns
	db, err1 := conn.DB()
	if err1 != nil {
		fmt.Println("get db failed:", err)
		return
	}
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)
}

func CreatAccount() {
	fmt.Println("hi")
}

func DeleteAccount() {
	fmt.Println("hi")
}

func ChangePwd() {
	fmt.Println("hi")
}
