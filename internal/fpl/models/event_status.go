package models

import (
	"fmt"
	"time"
	"github.com/henryppercy/fpl-go-bot/internal/utils"
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

func (es EventStatus) DateInCurrentEvent(date time.Time) bool {
	for _, status := range es.Status {
		eventDate, err := time.Parse("2006-01-02", status.Date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}

		if utils.SameDate(eventDate, date) {
			return true
		}
	}
	return false
}

func (es EventStatus) BonusAdded(date time.Time) bool {
	if !es.DateInCurrentEvent(date) {
		return false
	}

	for _, status := range es.Status {
		eventDate, err := time.Parse("2006-01-02", status.Date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}

		if utils.SameDate(eventDate, date) && status.BonusAdded {
			return true
		}
	}
	return false
}
