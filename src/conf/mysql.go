package conf

import (
	"github.com/jinzhu/gorm"
)

func GetMysqlDb(v interface{}) (db *gorm.DB) {
	var err error
	//连接串
	db, err = gorm.Open("mysql", GlobalConfig.MysqlUrl)
	//defer db.Close()
	if err != nil {
		panic(err)
	}
	//设置最大空闲连接数和最大连接
	db.DB().SetMaxIdleConns(GlobalConfig.MysqlMaxIdleConns)
	db.DB().SetMaxOpenConns(GlobalConfig.MysqlMaxOpenConns)
	//true:不使用结构体名称的复数形式映射生成表名
	db.SingularTable(true)
	//设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return GlobalConfig.MysqlTbPrefix + defaultTableName
	}
	//判断表是否存在，不存在则创建
	if !db.HasTable(v) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(v).Error; err != nil {
			panic(err)
		}
	}
	return db
}
