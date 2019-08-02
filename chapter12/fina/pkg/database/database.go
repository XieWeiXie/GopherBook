package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/xormplus/core"
)

var MySQL *xorm.Engine

var (
	dbMySQL    = "fifa"
	dbUser     = "root"
	dbPassword = "admin123"
)

func MySQLInit() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPassword, dbMySQL))
	if err != nil {
		panic(fmt.Sprintf("CONNECT ENGINE BY XORM FAIL %s", err.Error()))
	}
	engine.SetTableMapper(core.SameMapper{})
	engine.SetMapper(core.GonicMapper{})
	MySQL = engine
	MySQL.ShowSQL(true)
	return MySQL

}

var POSTGRES *xorm.Engine
var (
	dbPGSQL      = "fifa"
	dbPGUser     = "postgres"
	dbPGPassword = "admin123"
)

func PostgreSQLInit() *xorm.Engine {
	return POSTGRES
}
