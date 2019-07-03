package conf

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

var url []string

type KafkaClient struct {
	Topic string
}

func InitKafka() {
	url = strings.Split(GlobalConfig.KafkaHosts, ",")
}

func (this *KafkaClient) KafkaConsumer(msgChan chan string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	// 根据给定的代理地址和配置创建一个消费者
	consumer, err := sarama.NewConsumer(url, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	//Partitions(topic):该方法返回了该topic的所有分区id
	partitionList, err := consumer.Partitions(this.Topic)
	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		//ConsumePartition方法根据主题，分区和给定的偏移量创建创建了相应的分区消费者
		//如果该分区消费者已经消费了该信息将会返回error
		//sarama.OffsetNewest:表明了为最新消息
		pc, err := consumer.ConsumePartition(this.Topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			//Messages()该方法返回一个消费消息类型的只读通道，由代理产生
			for msg := range pc.Messages() {
				//fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				setMsg(msgChan, string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
}

func setMsg(msgs chan string, msg string) {
	msgs <- msg
}

//同步消息模式
func (this *KafkaClient) SyncProducer(key string, value string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(url, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	defer p.Close()
	topic := this.Topic

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(value),
	}
	part, offset, err := p.SendMessage(msg)
	if err != nil {
		log.Printf("send message(%s) err=%s \n", value, err)
	} else {
		fmt.Fprintf(os.Stdout, value+"发送成功，partition=%d, offset=%d \n", part, offset)
	}
	//time.Sleep(2 * time.Second)
}
