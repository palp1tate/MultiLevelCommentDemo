package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname string `gorm:"not null;index;varchar(20)"` // 昵称
	Password string `gorm:"not null"`                   // 密码
	Avatar   string `gorm:"not null;"`                  // 头像
}
