package main

//test xorm map
//test glog rotate

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

//{"denyAction":{"atOp":1,"atVal":"default"},"skipAction":null}
type AtCfg struct {
	Name string
	Age  int
}

type User struct {
	Id       int64
	Name     string
	Age      int64
	GroupId  int64
	Action   []AtCfg   `xorm:"action",json:"atCfg,omitempty"`
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

	//Session  transaction
	session := engine.NewSession()
	defer session.Close()
	session.Begin()

	u1 := User{
		Name: "fanpf",
	}
	if _, err := session.Get(&u1); err != nil {
		glog.Error("error:", err)
	} else {
		glog.Info("u:", u1)
	}

	for _, v := range u1.Action {
		glog.Info("name", v.Name)
	}

	//

	session.Commit()
}
