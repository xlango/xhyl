package conf

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var Client *redis.Client

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

func Get(key string) (string, bool) {
	r, err := Client.Get(key).Result()
	if err != nil {
		return "", false
	}
	return r, true
}

func SetExpTime(key string, val interface{}, expTime int32) {
	Client.Set(key, val, time.Duration(expTime)*time.Second)
}

func Set(key string, val interface{}) {

}
