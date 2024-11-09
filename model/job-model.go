package model

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title        string        `json:"title" gorm:"not null"`
	Description  string        `json:"description"`
	Skills       []Skill       `json:"skills" gorm:"many2many:job_skills"`
	Location     string        `json:"location"`
	Salary       string        `json:"salary"`
	CompanyID    uint          `json:"company_id"`
	Company      Company       `json:"company" gorm:"foreignKey:CompanyID;constraint:OnDelete:SET NULL"`
	Applications []Application `json:"applications" gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE"`
}
