package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	//Image   string `json:"image"`
}
