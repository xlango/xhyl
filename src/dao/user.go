package dao

import (
	"conf"
	"model"
)

type UserDao struct {
	User *model.User
}

func (this *UserDao) InsertUser() {
	msdb := conf.GetMysqlDb()
	defer msdb.Close()

	msdb.Create(this.User)
}

func (this *UserDao) GetUserByIf(pageSize int, pageIndex int) []model.User {
	msdb := conf.GetMysqlDb()
	defer msdb.Close()
	users := make([]model.User, 0)
	msdb.Model(model.User{}).Where(this.User).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&users)
	return users
}

func (this *UserDao) UpdateUser() {
	msdb := conf.GetMysqlDb()
	defer msdb.Close()
	msdb.Model(this.User).Updates(this.User)
}

func (this *UserDao) DeleteUser(ids []int) {
	msdb := conf.GetMysqlDb()
	defer msdb.Close()
	for _, id := range ids {
		msdb.Where("id = ?", id).Delete(model.User{})
	}
}
