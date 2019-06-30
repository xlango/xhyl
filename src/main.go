package main

import (
	"conf"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"model"
	"strconv"
	"time"
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
		Id:       1,
		Username: "234",
		Password: "234",
	}

	jsons, _ := json.Marshal(user)
	set(strconv.FormatInt(user.Id, 10), string(jsons), -1)
	s, b := get(strconv.FormatInt(user.Id, 10))
	fmt.Println(s, b)
	//userdao := &dao.UserDao{
	//	User: user,
	//}
	//ids := []int{4, 5}
	//userdao.DeleteUser(ids)
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

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:         "39.108.147.36:6379",
		PoolSize:     1000,
		ReadTimeout:  time.Millisecond * time.Duration(100),
		WriteTimeout: time.Millisecond * time.Duration(100),
		IdleTimeout:  time.Second * time.Duration(60),
	})

	_, err := Client.Ping().Result()
	if err != nil {
		panic("init redis error")
	} else {
		fmt.Println("init redis ok")
	}
}

func get(key string) (string, bool) {
	r, err := Client.Get(key).Result()
	if err != nil {
		return "", false
	}

	return r, true
}

func set(key string, val interface{}, expTime int32) {
	Client.Set(key, val, time.Duration(expTime)*time.Second)
}
