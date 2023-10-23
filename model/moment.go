package model

import "gorm.io/gorm"

type Moment struct {
	gorm.Model
	UserId  int    `gorm:"not null;index"`
	User    User   `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content string `gorm:"size:2048"`
}
