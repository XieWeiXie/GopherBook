package database_v1

import (
	"fmt"

	"xorm.io/core"

	"github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
)

var BeeQuickDatabase *xorm.Engine

var (
	drivenName string
	db         string
	port       string
	password   string
	user       string
	dsn        string
)

func init() {
	drivenName = "mysql"
	if drivenName == "mysql2" {
		config := mysql.NewConfig()
		config = &mysql.Config{
			User:   user,
			Passwd: password,
			DBName: db,
			Addr:   port,
		}
		dsn = config.FormatDSN()
	} else {
		//dsn = fmt.Sprintf("%s:%s@/%s?chartset=utf8&parseTime=true&loc=Local", user, password,db)
		dsn = fmt.Sprintf("root:admin123@/beequick_dev?charset=utf8&parseTime=true&loc=Local")
	}

}

func DataBaseInit() {
	var err error
	BeeQuickDatabase, err = xorm.NewEngine(drivenName, dsn)
	if err != nil {
		panic(err)
		return
	}
	BeeQuickDatabase.ShowSQL(true)
	BeeQuickDatabase.Logger()
	BeeQuickDatabase.Charset("utf8")
	BeeQuickDatabase.SetMapper(core.GonicMapper{})
	BeeQuickDatabase.SetTableMapper(core.SameMapper{})

}
