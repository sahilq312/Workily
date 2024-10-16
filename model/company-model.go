package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"not null"`
	Logo    string `json:"logo"`
	OwnerID uint   `json:"owner_id" gorm:"not null"`
	Owner   User   `json:"owner" gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE"`  // Cascade deletion
	Jobs    []Job  `json:"jobs" gorm:"foreignKey:CompanyID;constraint:OnDelete:CASCADE"` // Jobs related to the company
}
