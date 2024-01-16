package fpl

import (
	"os"
	"strconv"
	"time"

	"github.com/henryppercy/fpl-go-bot/internal/discord"
	"github.com/henryppercy/fpl-go-bot/internal/logger"
	"github.com/robfig/cron/v3"
)

var messageSentToday bool
var noFootballYesterday bool

func IntiCron() {
	c := cron.New()

	c.AddFunc("0 16 * * *", func() {
		logger.InfoLogger.Println("resetting message sent boolean")

		messageSentToday = false
	})

	c.AddFunc(("0 11-15 * * *"), func() {
		if noFootballYesterday {
			logger.InfoLogger.Println("no football played yesterday")
			return
		}

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

		fy := es.DateInCurrentEvent(time.Now().AddDate(0, 0, -1))

		if !fy {
			logger.InfoLogger.Println("no football played yesterday")
			noFootballYesterday = true
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

	err = discord.DispatchMessage(league.String())
	if err != nil {
		return err
	}

	return nil
}
