package dao

import (
	"model"
)

type UserDao struct {
	User *model.User
}

func (this *UserDao) InsertUser() {
	//msdb := conf.GetMysqlDb()
}

//func (this *UserDao) GetUserByIf() []model.User {
//	msdb := conf.GetMysqlDb()
//	//同步数据库结构
//	msdb.Sync2(new(model.User))
//
//	users := make([]model.User, 0)
//	//sql := "select * from tb_user where 1=1"
//	//if this.User.Id != 0 {
//	//	sql = fmt.Sprintf("%s %s", sql, " and id=?")
//	//}
//	////if this.User.Username != "" || len(this.User.Username) != 0 {
//	////	sql = fmt.Sprintf("%s %s", sql, " and username =?")
//	////}
//	//results, _ := msdb.Query(sql, this.User.Id)
//	//for _, i := range results {
//	//	bytes := i["id"]
//	//	user := model.User{}
//	//	user.Username = string(i["username"])
//	//	user.Password = string(i["password"])
//	//	users = append(users, user)
//	//	fmt.Println(users, BytesToInt64(bytes))
//	//}
//	msdb.Where(&this.User).Find(&users)
//	return users
//}
