package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	MomentId int      `gorm:"not null;index"`
	Moment   Moment   `gorm:"foreignKey:MomentId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId   int      `gorm:"not null;index"`
	User     User     `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ParentId *int     `gorm:"index;default:NULL"`
	Parent   *Comment `gorm:"foreignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content  string   `gorm:"not null;size:1024"`
}

type CommentTree struct {
	CommentId int                    `json:"cid"`
	Content   string                 `json:"content"`
	Author    map[string]interface{} `json:"author"`
	CreatedAt string                 `json:"createdAt"`
	Children  []*CommentTree         `json:"children"`
}
