package fpl

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/henryppercy/fpl-go-bot/internal/discord"
	"github.com/henryppercy/fpl-go-bot/internal/utils"
	"github.com/robfig/cron/v3"
)

var messageSentToday bool

func IntiCron() {
	c := cron.New()

	c.AddFunc("0 16 * * *", func() {
		utils.LogPrettyTime()
		fmt.Print("Resetting message sent.")

        messageSentToday = false
    })

	c.AddFunc(("0 11-15 * * *"), func() {
		fmt.Println()
		utils.LogPrettyTime()

		if (messageSentToday) {
			fmt.Print("Message already sent today.\n")
			return
		} else {
			fmt.Print("Message has not been sent yet today.\n")
		}

		es, err := GetEventStatus()
		if err != nil {
			log.Fatal("Error getting event status")
		}

		ba := es.BonusAdded(time.Now().AddDate(0, 0, -1))
		if (ba) {
			fmt.Print("Bonus points have been added!.\n")

			dispatchLeagueMessage()
			messageSentToday = true
		} else {
			fmt.Print("Bonus points not added yet.\n")
		}
	})
	
	c.Start()
}

func dispatchLeagueMessage() {
	id, _ := strconv.Atoi(os.Getenv("LEAGUE_ID"))

	league, err := GetLeague(id)
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
}
