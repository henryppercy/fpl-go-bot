package models

import (
	"time"

	"github.com/henryppercy/fpl-go-bot/internal/logger"
	"github.com/henryppercy/fpl-go-bot/internal/date"
)

type status struct {
	BonusAdded bool   `json:"bonus_added"`
	Date       string `json:"date"`
	Event      int    `json:"event"`
	Points     string `json:"points"`
}

type EventStatus struct {
	Status  []status `json:"status"`
	Leagues string   `json:"leagues"`
}

func (es EventStatus) DateInCurrentEvent(d time.Time) bool {
	for _, status := range es.Status {
		eventDate, err := time.Parse("2006-01-02", status.Date)
		if err != nil {
			logger.ErrorLogger.Panicln("error parsing date:", err)
			continue
		}

		if date.SameDate(eventDate, d) {
			return true
		}
	}
	return false
}

func (es EventStatus) BonusAdded(d time.Time) bool {
	if !es.DateInCurrentEvent(d) {
		return false
	}

	for _, status := range es.Status {
		eventDate, err := time.Parse("2006-01-02", status.Date)
		if err != nil {
			logger.ErrorLogger.Panicln("error parsing date:", err)
			continue
		}

		if date.SameDate(eventDate, d) && status.BonusAdded {
			return true
		}
	}
	return false
}
