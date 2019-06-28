package model

import "time"

type User struct {
	Id       int64  `gorm:"primary_key"`
	Username string `gorm:"type:varchar(20);not null;"`
	Password string `gorm:"type:varchar(256);"`
}

type Like struct {
	ID int `gorm:"primary_key"`
	//Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua string `gorm:"type:varchar(256);"`
	//Title     string `gorm:"type:varchar(128);index:title_idx"`
	//Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}
