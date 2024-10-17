package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	Skills      pq.StringArray `json:"skills" gorm:"type:text[]"`
	Location    string         `json:"location"`
	Salary      string         `json:"salary"`
	CompanyID   uint           `json:"company_id"`
	Company     Company        `json:"company" gorm:"foreignKey:CompanyID;constraint:OnDelete:SET NULL"` // Keeps job even if company is deleted
}
