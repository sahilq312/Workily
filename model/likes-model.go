package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID uint `json:"user_id"`
	PostID uint `json:"post_id"`
	User   User `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Post   Post `json:"post" gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
}
