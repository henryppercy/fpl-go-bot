package main

import (
	"fmt"
	"github.com/henryppercy/fpl-go-bot/internal/discord"
	"github.com/henryppercy/fpl-go-bot/internal/fpl"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	es, err := fpl.GetEventStatus()
	if err != nil {
		log.Fatal("Error getting event status")
	}

	ba := es.BonusAdded(time.Now().AddDate(0, 0, -1))
	if ba {
		fmt.Print("Bonus points added!\n")

		id, _ := strconv.Atoi(os.Getenv("LEAGUE_ID"))

		league, err := fpl.GetLeague(id)
		if err != nil {
			log.Printf("Error fetching league data for ID %d: %v\n", id, err)
		}

		channelId := os.Getenv("CHANNEL_ID")

		var res *http.Response
		res, err = discord.Send(channelId, league.String())
		if err != nil {
			log.Printf("Error sending discord message: %v\n", err)
		}

		if res != nil {
			fmt.Printf("Discord message sent successfully with code: %d\n", res.StatusCode)
		}
	} else {
		fmt.Print("Event not yet full complete.\n")
	}
}
