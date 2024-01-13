package main

import (
	"github.com/henryppercy/fpl-go-bot/internal/fpl"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fpl.IntiCron()
}
