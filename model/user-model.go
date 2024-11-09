package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string        `json:"name" gorm:"not null"`
	Email        string        `json:"email" gorm:"unique;not null"`
	Password     string        `json:"-"`
	Experience   []Experience  `json:"experience" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Posts        []Post        `json:"posts" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Skills       []Skill       `json:"skills" gorm:"many2many:user_skills"`
	Education    []Education   `json:"education" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Likes        []Like        `json:"likes" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Comments     []Comment     `json:"comments" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Applications []Application `json:"applications" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
