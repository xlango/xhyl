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
	}

	//初始化日志
	logger.InitLogger(currentDir)

	//初始化mysql,创建表
	//InitMysqlDb()
	InitTable()

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

}
