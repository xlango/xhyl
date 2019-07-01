package conf

import (
	"gopkg.in/mgo.v2"
)

type MongoClient struct {
	Database   string //数据库
	Collection string //集合
}

func (this *MongoClient) Insert(v interface{}) bool {
	mongo, err := mgo.Dial(GlobalConfig.MongoHost) // 建立连接

	defer mongo.Close()

	if err != nil {
		return false
	}

	client := mongo.DB(this.Database).C(this.Collection) //选择数据库和集合

	//插入数据
	cErr := client.Insert(v)

	if cErr != nil {
		return false
	}
	return true
}
