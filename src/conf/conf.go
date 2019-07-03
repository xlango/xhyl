package conf

import (
	"encoding/json"
	"io/ioutil"
	"logger"
	"os"
	"runtime"
	"utils"
)

var GlobalConfig *ConfigInfo

type ConfigInfo struct {
	MysqlUrl          string //mysql连接
	MysqlTbPrefix     string //Mysql表前缀
	MysqlMaxIdleConns int    //Mysql最大空闲连接
	MysqlMaxOpenConns int    //Mysql最大连接

	RedisHost         string //redis ip:port
	RedisPoolSize     int    //redis连接池大小
	RedisReadTimeout  int    //redis 读数据超时
	RedisWriteTimeout int    //redis 写数据超时
	RedisIdleTimeout  int    //空闲连接时长

	MongoHost string //mongodb连接

	KafkaHosts string //kafka连接
}

func InitConfig() {
	//解决windows无法debug启动的问题
	currentDir := utils.GetCurrentExeDir()
	if len(currentDir) > 0 {
		currentDir = currentDir + string(os.PathSeparator)
	}

	//获取主机可用cpu数，配置程序使用cpu核数
	cpuNumber := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNumber)

	//如果配置文件未配置，则使用默认配置
	GlobalConfig = &ConfigInfo{
		MysqlUrl:          "root:123456@tcp(127.0.0.1:3306)/comprehensive?charset=utf8",
		MysqlTbPrefix:     "",
		MysqlMaxIdleConns: 150,
		MysqlMaxOpenConns: 250,
		RedisHost:         "127.0.0.1:6379",
		RedisPoolSize:     1000,
		RedisReadTimeout:  100,
		RedisWriteTimeout: 100,
		RedisIdleTimeout:  60,
		MongoHost:         "127.0.0.1:27017",
		KafkaHosts:        "127.0.0.1:9092",
	}

	//初始化日志
	logger.InitLogger(currentDir)

	//读取config.json配置文件相关配置
	configPath := currentDir + "config.json"
	logger.LogInfo("Config path: " + configPath)
	configContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		logger.LogError(err)
	}
	err = json.Unmarshal(configContent, GlobalConfig)
	if err != nil {
		logger.LogError(err)
	}

	//初始化mysql,创建表
	//InitMysqlDb()
	InitTable()

	//初始化redis
	//InitRedis()

	//初始化Kafka
	InitKafka()
}
