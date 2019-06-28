package main

import (
	"conf"
	"dao"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"model"
	"strconv"
)

func main() {
	//初始化配置
	conf.InitConfig()

	//ch := make(chan string, 10)
	//go write(ch)
	//go insert(ch)
	//
	//time.Sleep(time.Second * 20)

	user := &model.User{
		//Id:       1,
		//Username: "234",
		//Password: "234",
	}
	userdao := &dao.UserDao{
		User: user,
	}
	ids := []int{4, 5}
	userdao.DeleteUser(ids)
	//userdao.UpdateUser()
	//userdao.InsertUser()
	//users := userdao.GetUserByIf(2, 2)
	//fmt.Println(users)
}

func write(ch chan string) {
	for i := 0; i < 20; i++ {
		ch <- strconv.Itoa(i)
	}
}

func insert(ch chan string) {
	for {
		var y string
		y = <-ch
		msdb := conf.GetMysqlDb()
		defer msdb.Close()
		fmt.Println(y)
		like := &model.Like{
			Ua: y,
		}
		msdb.Create(&like)
		record := msdb.NewRecord(like)
		fmt.Println(record)
	}
}
