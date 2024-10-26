package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"content" gorm:"not null"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
	User    User   `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Post    Post   `json:"post" gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
}
