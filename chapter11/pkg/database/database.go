package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
)

var EngineMySQLGORM *gorm.DB

func EngineGORMInit() {
	db, err := gorm.Open("mysql", "root:admin123@/votes?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("CONNECT DB FAIL: ", err.Error())
		return
	}
	db.LogMode(true)
	EngineMySQLGORM = db
}

var EngineMySQLXORM *xorm.Engine

func EngineXORMInit() {
	db, err := xorm.NewEngine("mysql", "root:admin123@/votes?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("CONNECT DB FAIL :", err.Error())
		return
	}
	db.ShowSQL(true)
	EngineMySQLXORM = db
}
