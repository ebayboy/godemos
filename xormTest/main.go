package main

import (
	"time"

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
	CreateAt int64     `xorm:"created"`
	DeleteAt time.Time `xorm:"deleted"`
	UpdateAt int64     `xorm:"updated"`
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

	//Insert
	rows, err2 := engine.Insert(&u1)
	if err2 != nil {
		glog.Error("Error:", err2.Error())
	}
	glog.Info("rows:", rows)

	//Delete
	var u2 User
	rows, err = engine.Id(8).Delete(&u2)
	if err != nil {
		glog.Error("Error:", err.Error())
	}
	glog.Info("del rows:", rows)

	//Get
	u3 := User{
		Name: "888",
	}
	var found bool
	found, err = engine.Get(&u3)
	if err != nil {
		glog.Error("Error:", err.Error())
	}
	if !found {
		glog.Warning("not found user!")
	} else {
		glog.Info("found user:", u3)
	}

	//Update
	u3.Age = 999
	rows, err = engine.Id(u3.Id).Update(&u3)
	if err != nil {
		glog.Error("error:", err.Error())
	}
	if rows > 0 {
		glog.Infof("update user %s ok!", u3.Name)
	} else {
		glog.Warningf("update user %s failed!", u3.Name)
	}
}
