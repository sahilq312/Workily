package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariale() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading the env file")
	}
}
