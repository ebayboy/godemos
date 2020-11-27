package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"time"
)

type User struct {
	Id      int
	Name    string    `xorm:"name"`
	Age     int       `xorm:"age"`
	Passwd  string    `xorm:"passwd"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {

	//数据库名称:数据库连接密码@(数据库地址:3306)/数据库实例名称?charset=utf8
	engine, err := xorm.NewEngine("mysql", "root:MyNewPass@123@(192.168.137.101:3306)/test?charset=utf8mb4")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ping .... OK")
	}
	defer engine.Close()

	engine.ShowSQL(true) // 显示SQL的执行, 便于调试分析
	engine.SetTableMapper(core.SnakeMapper{})

	//create session
	session := engine.NewSession()
	defer session.Close()

	//select
	results, err := session.QueryString("select * from user")
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("results:", results)
	}

	//select ...  where
	u2 := new(User)
	res2, err := session.Where("name=?", "fanpf").Get(u2)
	if err != nil {
		fmt.Println("Err:", err)
	} else {
		fmt.Println("res2 where:", res2)
		fmt.Println("u2 :", u2)
	}

	//insert
	u1 := new(User)
	u1.Name = "fanpf"
	u1.Age = 38
	u1.Passwd = "Passwd123"
	u1.Created = time.Now()
	u1.Updated = time.Now()

	affected, err := session.Insert(u1)
	fmt.Println("affected:", affected)

	//update
	u3 := User{Age: 40}
	affected, err = session.Where("name=?", "fanpf").Update(&u3)
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("affected:", affected)
	}

	//delete
	u4 := User{Name: "fanpf1"}
	affected, err = session.Delete(&u4)
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("affected:", affected)
	}

	//find return []
	var users []User
	err = session.Table("user").Where("user.name = ?", "fanpf").Limit(10, 0).Find(&users)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("users:", users)
	}

	//find all
	allusers := make([]User, 0, 10)
	err = session.Table("user").Find(&allusers)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("allusers size:", len(allusers))
	}
}
