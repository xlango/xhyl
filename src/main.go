package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"xhyl/conf"
	"xhyl/consul"
)

const URL = "39.108.147.36:27017" //mongodb的地址

func main() {
	//初始化配置
	conf.InitConfig()

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		servers, _ := consul.GetNodeServerInfo(conf.GlobalConfig.ConsulRegisterName, conf.GlobalConfig.ConsulRegisterTags)
		bytes, _ := json.Marshal(servers)
		writer.Write(bytes)
	})
	consul.RegisterServer()

	//consul.GetNodeServerInfo(conf.GlobalConfig.ConsulRegisterName, conf.GlobalConfig.ConsulRegisterTags)
}
