package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"xorm.io/core"
)

var engine *xorm.Engine

type User struct {
	Id       int64
	Name     string
	Age      int64
	CreateAt int64 `xorm:"created"`
}

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@(10.0.2.15:3306)/test?charset=utf8")
	if err != nil {
		glog.Error("Error:", err.Error())
	}
	engine.SetColumnMapper(core.SnakeMapper{})

	u := User{}
	tbl := engine.TableInfo(&u)
	glog.Info("tbl:", tbl)

	tbName := engine.TableName(&u)
	glog.Info("name:", tbName)

	dbMetas, err1 := engine.DBMetas()
	if err1 != nil {
		glog.Error("Error:", err1.Error())
	}
	glog.Info("dbMetas:", dbMetas)

	for k, v := range dbMetas {
		glog.Info(k, ":", v)
	}

	u1 := User{
		Name: "rose",
		Age:  33,
	}

	rows, err2 := engine.Insert(&u1)
	if err2 != nil {
		glog.Error("Error:", err2.Error())
	}
	glog.Info("rows:", rows)
}
