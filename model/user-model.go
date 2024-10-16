package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string       `json:"name" gorm:"not null"`
	Email      string       `json:"email" gorm:"unique;not null"`
	Password   string       `json:"-"`
	Experience []Experience `json:"experience" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Posts      []Post       `json:"posts" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Skills     []string     `json:"skills" gorm:"type:text[]"`
	Education  []Education  `json:"education" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Followers  []User       `json:"followers" gorm:"many2many:user_followers;constraint:OnDelete:CASCADE"`
	Following  []User       `json:"following" gorm:"many2many:user_following;constraint:OnDelete:CASCADE"`
}
