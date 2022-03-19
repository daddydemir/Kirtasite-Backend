package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConnect() {

	dsn := "root:@tcp(127.0.0.1:3306)/kirtasiye_test?parseTime=true"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
