package service

import (
	"os"
	"strconv"
	"time"

	"github.com/henryppercy/fpl-go-bot/internal/discord"
	"github.com/henryppercy/fpl-go-bot/internal/fpl"
	"github.com/henryppercy/fpl-go-bot/internal/logger"
	"github.com/henryppercy/fpl-go-bot/internal/scheduler"
)

var messageSentToday bool
var noFootballYesterday bool

func ScheduleFplJobs() {
	scheduler := scheduler.NewCronScheduler()

	_, err := scheduler.ScheduleTask("0 16 * * *", resetMessageSentBoolean)
	if err != nil {
		logger.ErrorLogger.Println("failed to message reset task")
	}

	_, err = scheduler.ScheduleTask("0 11-15 * * *", checkAndSendLeagueUpdate)
	if err != nil {
		logger.ErrorLogger.Println("failed to schedule league update task")
	}

	_, err = scheduler.ScheduleTask("0 9 * * *", checkAndSendDeadlineReminder)
	if err != nil {
		logger.ErrorLogger.Println("failed to schedule deadline update task")
	}

	scheduler.StartScheduler()
}

func checkAndSendDeadlineReminder() {
	lb, err := fpl.GetBootstrap()
	if err != nil {
		logger.ErrorLogger.Println("unable to retrieve league bootstrap data")
		return
	}

	event, exists := lb.GetNextEvent()

	if exists && event.IsDeadlineDay() {
		err := discord.DispatchMessage(event.FormatDeadlineMessage())
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
	}

	logger.InfoLogger.Println("not deadline day")
}

func resetMessageSentBoolean() {
	logger.InfoLogger.Println("resetting message sent boolean")

	messageSentToday = false
}

func checkAndSendLeagueUpdate() {
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

	es, err := fpl.GetEventStatus()
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
}

func dispatchLeagueMessage() error {
	id, _ := strconv.Atoi(os.Getenv("LEAGUE_ID"))

	league, err := fpl.GetLeague(id)
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
