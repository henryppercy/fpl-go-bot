package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/henryppercy/fpl-go-bot/internal/api"
	"github.com/henryppercy/fpl-go-bot/internal/discord"
	"github.com/henryppercy/fpl-go-bot/internal/logger"
	"github.com/henryppercy/fpl-go-bot/internal/service"
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

	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()

	server := api.NewServer(*listenAddr)
	fmt.Println("server running on port: ", *listenAddr)
	log.Fatal(server.Start())
}
