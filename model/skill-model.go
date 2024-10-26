package model

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name        string       `json:"name" gorm:"unique;not null"`
	Users       []User       `gorm:"many2many:user_skills"`
	Jobs        []Job        `gorm:"many2many:job_skills"`
	Experiences []Experience `gorm:"many2many:experience_skills"`
}
