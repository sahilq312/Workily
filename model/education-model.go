package model

import (
	"time"

	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	School    string    `json:"school" gorm:"not null"`
	Degree    string    `json:"degree" gorm:"not null"`
	Field     string    `json:"field"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // Each education belongs to a user
}
