package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USERNAME = "demo1"
	PASSWORD = "demo123"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "demo"
)

var db *gorm.DB

func init() {
	var err error
	addr := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

	db, err = gorm.Open(mysql.Open(addr), &gorm.Config{})

	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}

	if db.Error != nil {
		fmt.Printf("database error %v", db.Error)
	}
}
