package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null"`
	Logo    string `json:"logo"`
	OwnerID uint   `json:"owner_id" gorm:"not null"`
	Owner   User   `json:"owner" gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE"`
	Jobs    []Job  `json:"jobs" gorm:"foreignKey:CompanyID;constraint:OnDelete:CASCADE"`
}
