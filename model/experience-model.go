package model

import (
	"time"

	"gorm.io/gorm"
)

type Experience struct {
	gorm.Model
	Title       string    `json:"title" gorm:"not null"`
	Company     string    `json:"company" gorm:"not null"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	UserID      uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // Experience belongs to a user
	Skills      []string  `json:"skills" gorm:"type:text[]"`
}
