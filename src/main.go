package main

import (
	"conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Like struct {
	ID int `gorm:"primary_key"`
	//Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua string `gorm:"type:varchar(256);"`
	//Title     string `gorm:"type:varchar(128);index:title_idx"`
	//Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}
type DD struct {
	Db *gorm.DB
}

func (this *DD) aa() {
	this.Db, _ = gorm.Open("mysql", "root:wolaoda.sheigandong@tcp(39.108.147.36:3306)/comprehensive?charset=utf8")
	defer this.Db.Close()
	this.Db.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tb_" + defaultTableName
	}

	a := Like{}
	first := this.Db.Where("ip = ?", "111111").First(&a)
	fmt.Println(first.Value)
}
func main() {
	//初始化配置
	conf.InitConfig()
	msdb := conf.GetMysqlDb(Like{})
	defer msdb.Close()

	like := &Like{
		Ua: "222",
	}
	msdb.NewRecord(like) // => 主键为空返回`true`
	msdb.Create(&like)
	record := msdb.NewRecord(like)
	fmt.Println(record)

	////插入数据
	//userdao := dao.UserDao{
	//	User: &model.User{
	//		//Id:       1,
	//		//Username: "333",
	//		Password: "33333"},
	//}
	////userdao.InsertUser()
	//
	//users := userdao.GetUserByIf()
	//fmt.Println(users)

	////查询id为1的数据
	//p1 := model.User{}
	//msdb.Id(1).Get(&p1)
	//fmt.Println(p1)
	////查询name为yuanye的数据
	//p2 := model.User{}
	//msdb.Where("username = ?", "xx").Get(&p2)
	//fmt.Println(p2)

	//msdb := conf.GetMysqlDb()
	//if !msdb.HasTable(&Like{}) {
	//	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
	//		panic(err)
	//	}
	//}
	//a := Like{Ip: "22222"}

	//msdb.Create(&a)

	//likes := make([]Like, 0)
	//db.Find(&likes)

}
