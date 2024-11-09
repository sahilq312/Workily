package main

import (
	"github.com/sahilq312/workly/initializer"
	"github.com/sahilq312/workly/model"
)

func init() {
	initializer.LoadEnvVariale()
	initializer.ConnectPostgresDatabase()
}

func main() {
	initializer.DB.AutoMigrate(
		&model.User{},
		&model.Experience{},
		&model.Education{},
		&model.Post{},
		&model.Company{},
		&model.Job{},
		&model.Skill{},
		&model.UserFollow{},
		&model.Like{},
		&model.Comment{},
		&model.Application{},
	)
}
