package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Logo    string `json:"logo"`
	OwnerID uint   `json:"owner_id"`
	Owner   User   `json:"owner"`
	Jobs    []Job  `json:"jobs"`
}
