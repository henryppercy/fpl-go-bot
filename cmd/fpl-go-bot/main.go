package main

import (
	"fmt"
	"github.com/henryppercy/fpl-go-bot/internal/fpl"
	"github.com/henryppercy/fpl-go-bot/internal/notify"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	leagueIds := os.Args[1:]

	for _, idStr := range leagueIds {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("Invalid league ID '%s': %v\n", idStr, err)
			continue
		}

		league, err := fpl.GetLeague(id)
		if err != nil {
			log.Printf("Error fetching league data for ID %d: %v\n", id, err)
			continue
		}

		whatsAppMsgId := os.Getenv("WHATSAPP_MSG_ID")

		var res *http.Response
		res, err = notify.Send(whatsAppMsgId, league.String())
		if err != nil {
			log.Printf("Error fetching sending whatsapp message: %v\n", err)
		}

		if res != nil {
			fmt.Printf("WhatsApp sent successfully with code: %d\n", res.StatusCode)
		}

		fmt.Printf("%+v\n", league.String())
	}
}
