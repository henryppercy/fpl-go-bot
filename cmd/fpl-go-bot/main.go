package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/henryppercy/fpl-go-bot/internal/discord"
	"github.com/henryppercy/fpl-go-bot/internal/service"
	"github.com/henryppercy/fpl-go-bot/internal/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.ErrorLogger.Fatal("error loading .env file")
	}

	fmt.Println("Application starting...")
	service.ScheduleFplJobs()
	discord.DispatchUpdatedMessage()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	logger.WarningLogger.Println("shutting down the application")
	log.Println("Shutting down the application")
}
