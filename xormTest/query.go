package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"xorm.io/core"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@(10.0.2.15:3306)/cloudwaf_api_test?charset=utf8")
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

	//sql querystring

	sql := "select name from dn_risk_policy where waf_instance_id= ? and domain = ?"
	res, err := session.QueryString(sql, "waf-ins_cn-north-1_7a172aa9704a", "test.com")
	if err != nil {
		session.Rollback()
		glog.Fatal("Error:", err.Error())
	}

	glog.Info("res:", res)

	session.Commit()
}
