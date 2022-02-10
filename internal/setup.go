package internal

import (
	"github.com/joho/godotenv"
	"log"
)

func Setup() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Could not load env file.")
	}
}
