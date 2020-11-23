package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/golang/protobuf/proto"
	"log"
	"test/protobuf/pb"
)

type User struct {
	Id     int32
	Name   string `xorm:"name"`
	Age    int32  `xorm:"age"`
	Passwd string `xorm:"passwd"`
}

func GetFromDb() []User {
	//db init
	engine, err := xorm.NewEngine("mysql", "root:MyNewPass@123@(192.168.137.101:3306)/test?charset=utf8mb4")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ping .... OK")
	}

	engine.ShowSQL(true) // 显示SQL的执行, 便于调试分析
	engine.SetTableMapper(core.SnakeMapper{})

	users := make([]User, 0, 10)
	err = engine.Table("user").Where("user.name = ?", "fanpf").Find(&users)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("users:", users)
	}

	return users
}

//从数据库取出， 进行序列化打印
func main() {

	//get user from db
	users := GetFromDb()
	fmt.Println("users:", users)

	//add db users to pb_users
	var pb_users pb.MultiUser
	for i := 0; i < len(users); i++ {
		user1 := &pb.User{
			Id:     *proto.Int32(users[i].Id),
			Name:   *proto.String(users[i].Name),
			Passwd: *proto.String(users[i].Passwd),
		}

		pb_users.Users = append(pb_users.Users, user1)
	}

	// 序列化数据
	data, err := proto.Marshal(&pb_users)
	if err != nil {
		log.Fatalln("Marshal data error: ", err)
	}
	println(pb_users.Users[0].GetName())

	// 对已序列化的数据进行反序列化
	var target pb.MultiUser
	err = proto.Unmarshal(data, &target)
	if err != nil {
		log.Fatalln("Unmarshal data error: ", err)
	}
	println(target.GetUsers()[1].Name)
}
