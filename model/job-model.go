package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Skills      []string `json:"skills"`
	Location    string   `json:"location"`
	Salary      string   `json:"salary"`
	CompanyID   uint     `json:"company_id"`
	Company     Company  `json:"company"`
}
