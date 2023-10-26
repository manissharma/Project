package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	data, err := gorm.Open("mysql", "sharma:Sharma002@tcp(127.0.0.1:8090)/simple?charset=utf8&parseTime=True&loc=local")
	if err != nil {
		panic(err)
	}
	db = data
}

func GetDB() *gorm.DB {
	return db
}
