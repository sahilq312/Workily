package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	/* education string
	skill string
	expreince []string
	post string
	connection []string
	profile string
	chats []string */
}
