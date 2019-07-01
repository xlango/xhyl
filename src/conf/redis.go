package conf

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var Client *redis.Client

type RedisClient struct {
	DB int //redis库编号
}

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:         GlobalConfig.RedisHost,
		PoolSize:     GlobalConfig.RedisPoolSize,
		ReadTimeout:  time.Millisecond * time.Duration(GlobalConfig.RedisReadTimeout),
		WriteTimeout: time.Millisecond * time.Duration(GlobalConfig.RedisWriteTimeout),
		IdleTimeout:  time.Second * time.Duration(GlobalConfig.RedisIdleTimeout),
	})

	_, err := Client.Ping().Result()
	if err != nil {
		panic("init redis error")
	} else {
		fmt.Println("init redis ok")
	}
}

func (this *RedisClient) Get(key string) (string, bool) {
	Client.Options().DB = this.DB
	r, err := Client.Get(key).Result()
	if err != nil {
		return "", false
	}
	return r, true
}

func (this *RedisClient) SetExpTime(key string, val interface{}, expTime int32) {
	Client.Options().DB = this.DB
	Client.Set(key, val, time.Duration(expTime)*time.Second)
}

func (this *RedisClient) Set(key string, val interface{}) {

}
