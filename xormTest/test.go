package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"xorm.io/core"
)

var engine *xorm.Engine

type Groupp struct {
	Id   int64
	Name string
}

type User struct {
	Id       int64
	Name     string
	Age      int64
	GroupId  int64
	CreateAt int64     `xorm:"created"`
	DeleteAt time.Time `xorm:"deleted"`
	UpdateAt int64     `xorm:"updated"`
}

type UserGroup struct {
	User   `xorm:"extends"`
	Groupp `xorm:"extends"`
}

func (UserGroup) TableName() string {
	return "user"
}

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@(10.0.2.15:3306)/test?charset=utf8")
	if err != nil {
		glog.Error("Error:", err.Error())
		return
	}
	defer engine.Close()

	engine.SetColumnMapper(core.SnakeMapper{})

	u := User{}
	tbl := engine.TableInfo(&u)
	glog.Info("tbl:", tbl)

	tbName := engine.TableName(&u)
	glog.Info("name:", tbName)

	dbMetas, err1 := engine.DBMetas()
	if err1 != nil {
		glog.Error("Error:", err1.Error())
		return
	}
	glog.Info("dbMetas:", dbMetas)

	for k, v := range dbMetas {
		glog.Info(k, ":", v)
	}

	u1 := User{
		Name: "rose",
		Age:  33,
	}

	//Session  transaction
	session := engine.NewSession()
	defer session.Close()
	session.Begin()

	//inner join
	ugs := make([]UserGroup, 0)
	session.Join("INNER", "groupp", "groupp.id = user.group_id").Find(&ugs)
	glog.Info("ugs size:", len(ugs))

	session.Commit()
}
