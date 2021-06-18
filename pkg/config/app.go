package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("upay:12345678/go_bookstore?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil{
		panic(err)
	}
	db = d
}

func getDB() *gorm.DB {
	return db
}