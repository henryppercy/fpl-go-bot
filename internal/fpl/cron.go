package fpl

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/henryppercy/fpl-go-bot/internal/discord"
	"github.com/henryppercy/fpl-go-bot/internal/logger"
	"github.com/robfig/cron/v3"
)

var messageSentToday bool

func IntiCron() {
	c := cron.New()

	c.AddFunc("0 16 * * *", func() {
		logger.InfoLogger.Println("resetting message sent boolean")

		messageSentToday = false
	})

	c.AddFunc(("0 11-15 * * *"), func() {
		if messageSentToday {
			logger.InfoLogger.Println("message has already been sent today")
			return
		} else {
			logger.InfoLogger.Println("message has not been sent today")
		}

		es, err := GetEventStatus()
		if err != nil {
			logger.ErrorLogger.Println("failed to retrieve event status data from /event-status")
			return
		}

		ba := es.BonusAdded(time.Now().AddDate(0, 0, -1))
		if ba {
			logger.InfoLogger.Println("bonus points have been applied")

			err := dispatchLeagueMessage()
			if err != nil {
				return
			}
			messageSentToday = true
		} else {
			logger.InfoLogger.Println("bonus points have not been applied")
		}
	})

	c.Start()
	logger.InfoLogger.Println("cron initiated")
}

func dispatchLeagueMessage() error {
	id, _ := strconv.Atoi(os.Getenv("LEAGUE_ID"))

	league, err := GetLeague(id)
	if err != nil {
		logger.ErrorLogger.Printf("error fetching league data for ID %d: %v\n", id, err)
		return err
	}

	channelId := os.Getenv("CHANNEL_ID")

	var res *http.Response
	res, err = discord.Send(channelId, league.String())
	if err != nil {
		logger.ErrorLogger.Printf("error sending discord message: %v\n", err)
		return err
	}

	if res != nil {
		logger.InfoLogger.Printf("discord message sent successfully with code: %d\n", res.StatusCode)
	}

	return nil
}
