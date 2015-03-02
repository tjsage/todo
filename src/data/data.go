package data

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/coopernurse/gorp"
	"fmt"
	"log"
	"github.com/bradfitz/gomemcache/memcache"
	"config"
)

var dbMap *gorp.DbMap

func init() {
	config := config.GetConfig()


	db, _ := sql.Open("mysql", config.DbConnectionString)
	dbMap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	dbMap.AddTableWithName(Task{}, "Tasks").SetKeys(true, "TaskId")
}

func GetDbMap() *gorp.DbMap {
	return dbMap
}