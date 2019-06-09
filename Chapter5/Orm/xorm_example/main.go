package main

import (
	"GopherBook/Chapter5/Orm/xorm_example/model"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:admin123@/xorm_example?charset=utf8")
	if err != nil {
		log.Println(err)
		panic("mysql connect fail")
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "wechat_")
	engine.SetTableMapper(tbMapper)
	engine.Logger().SetLevel(core.LOG_WARNING)
	dropTable()
	syncTable()

}
func dropTable() {
	engine.DropTables(&model.Receipt{})
}
func syncTable() {
	engine.Sync2(new(model.Person))
}

func main() {
	tables, err := engine.DBMetas()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range tables {
		fmt.Println(i)
	}
}
