package main

import (
	"conf"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"model"
	"os"
	"strconv"
	"time"
)

var Address = []string{"192.168.10.108:9092"}
var RedisClient *conf.RedisClient
var MongoClient *conf.MongoClient

const URL = "39.108.147.36:27017" //mongodb的地址

func main() {
	//初始化配置
	conf.InitConfig()

	//ch := make(chan string, 10)
	//go write(ch)
	//go insert(ch)
	//
	//time.Sleep(time.Second * 20)

	//user := &model.User{
	//	Id: 3,
	//	//Username: "333333",
	//	//Password: "333333",
	//}
	//
	//jsons, _ := json.Marshal(user)
	//RedisClient.SetExpTime(strconv.FormatInt(user.Id, 10), string(jsons), -1)
	//s, b := RedisClient.Get(strconv.FormatInt(user.Id, 10))
	//fmt.Println(s, b)
	//userdao := &dao.UserDao{
	//	User: user,
	//}
	//ids := []int{4, 5}
	//userdao.DeleteUser(ids)
	//userdao.UpdateUser()
	//userdao.InsertUser()
	//users := userdao.GetUserByIf(2, 2)
	//fmt.Println(users)

	//rs := MongoClient.FindOne(map[string]interface{}{"username": "222"})
	//rs := MongoClient.FindAll(nil, model.User{}, 1, 2)
	//MongoClient.Update(map[string]interface{}{"id": 1}, map[string]interface{}{"username": "111111", "password": "111111"})
	//MongoClient.Delete(map[string]interface{}{"id": 2})
	//fmt.Println(rs)

	syncProducer(Address)
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

	MongoClient = &conf.MongoClient{
		Database:   "mydb_tutorial",
		Collection: "t_student",
	}
}

//同步消息模式
func syncProducer(address []string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	defer p.Close()
	topic := "test"
	srcValue := "sync: this is a message. index=%d"
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf(srcValue, i)
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(value),
		}
		part, offset, err := p.SendMessage(msg)
		if err != nil {
			log.Printf("send message(%s) err=%s \n", value, err)
		} else {
			fmt.Fprintf(os.Stdout, value+"发送成功，partition=%d, offset=%d \n", part, offset)
		}
		time.Sleep(2 * time.Second)
	}
}

func SaramaProducer() {

	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
	config.Version = sarama.V0_10_0_1

	fmt.Println("start make producer")
	//使用配置,新建一个异步生产者
	producer, e := sarama.NewAsyncProducer([]string{"182.61.9.153:6667", "182.61.9.154:6667", "182.61.9.155:6667"}, config)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer producer.AsyncClose()

	//循环判断哪个通道发送过来数据.
	fmt.Println("start goroutine")
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case <-p.Successes():
				//fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
			case fail := <-p.Errors():
				fmt.Println("err: ", fail.Err)
			}
		}
	}(producer)

	var value string
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		time11 := time.Now()
		value = "this is a message 0606 " + time11.Format("15:04:05")

		// 发送的消息,主题。
		// 注意：这里的msg必须得是新构建的变量，不然你会发现发送过去的消息内容都是一样的，因为批次发送消息的关系。
		msg := &sarama.ProducerMessage{
			Topic: "0606_test",
		}

		//将字符串转化为字节数组
		msg.Value = sarama.ByteEncoder(value)
		//fmt.Println(value)

		//使用通道发送
		producer.Input() <- msg
	}
}
