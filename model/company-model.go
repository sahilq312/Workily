package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Logo     string `json:"logo"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"-" gorm:"not null"`
	Address  string `json:"address,omitempty"`
	Jobs     []Job  `json:"jobs" gorm:"foreignKey:CompanyID;constraint:OnDelete:CASCADE"`
}
