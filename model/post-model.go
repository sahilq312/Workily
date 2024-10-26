package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string    `json:"title" gorm:"not null"`
	Content  string    `json:"content" gorm:"not null"`
	UserID   uint      `json:"user_id"`
	User     User      `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Likes    []Like    `json:"likes" gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
}
