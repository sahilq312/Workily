package model

import "gorm.io/gorm"

type UserFollow struct {
	gorm.Model
	FollowerID uint `json:"follower_id" gorm:"not null"`
	FollowedID uint `json:"followed_id" gorm:"not null"`
	Follower   User `json:"follower" gorm:"foreignKey:FollowerID;constraint:OnDelete:CASCADE"`
	Followed   User `json:"followed" gorm:"foreignKey:FollowedID;constraint:OnDelete:CASCADE"`
}
