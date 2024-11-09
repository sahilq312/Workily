package model

import (
	"time"

	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	UserID    uint      `json:"user_id" gorm:"not null"`
	JobID     uint      `json:"job_id" gorm:"not null"`
	Status    string    `json:"status" gorm:"default:'Pending'"`
	AppliedAt time.Time `json:"applied_at" gorm:"autoCreateTime"`
	User      User      `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Job       Job       `json:"job" gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE"`
}
