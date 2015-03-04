package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tjsage/todo/services"
	//"fmt"
	//"log"
	//"github.com/bradfitz/gomemcache/memcache"
	"log"
)

var dbMap *gorp.DbMap

func init() {
	config := services.GetConfig()
	log.Println("Connection String: ", config.DbConnectionString)
	db, err := sql.Open("mysql", config.DbConnectionString)

	if err != nil {
		log.Println(err)
	}
	dbMap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	dbMap.AddTableWithName(Task{}, "Tasks").SetKeys(true, "TaskId")
	log.Println(dbMap)
}

func GetDbMap() *gorp.DbMap {
	return dbMap
}
