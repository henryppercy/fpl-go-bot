package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/henryppercy/fpl-go-bot/internal/fpl"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fpl.IntiCron()

	sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
    log.Println("Shutting down the application")
}
