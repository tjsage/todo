package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tjsage/todo/services"
	//"fmt"
	//"log"
	//"github.com/bradfitz/gomemcache/memcache"
	"log"
)

var db *gorm.DB

func init() {
	dbPointer, err := gorm.Open("mysql", services.GetConfig().DbConnectionString)

	if err != nil {
		log.Println(err)
		return
	}

	db = &dbPointer

	db.AutoMigrate(&Task{})

}
