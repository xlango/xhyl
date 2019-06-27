package model

type User struct {
	Id       int64  `gorm:"primary_key"`
	Username string `gorm:"type:varchar(20);not null;"`
	Password string `gorm:"type:varchar(256);"`
}
