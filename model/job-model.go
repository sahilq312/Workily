package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	ID          uint     `json:"id" gorm:"primaryKey"`
	Title       string   `json:"title" gorm:"not null"`
	Description string   `json:"description"`
	Skills      []string `json:"skills" gorm:"type:text[]"`
	Location    string   `json:"location"`
	Salary      string   `json:"salary"`
	CompanyID   uint     `json:"company_id"`
	Company     Company  `json:"company" gorm:"foreignKey:CompanyID;constraint:OnDelete:SET NULL"` // Keeps job even if company is deleted
}
