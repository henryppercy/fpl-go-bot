package model

import (
	"fmt"
	"strings"
	"time"
)

type League struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

type PlayerStanding struct {
	ID         int    `json:"id"`
	EventTotal int    `json:"event_total"`
	PlayerName string `json:"player_name"`
	Rank       int    `json:"rank"`
	LastRank   int    `json:"last_rank"`
	RankSort   int    `json:"rank_sort"`
	Total      int    `json:"total"`
	Entry      int    `json:"entry"`
	EntryName  string `json:"entry_name"`
}

type Standings struct {
	HasNext bool             `json:"has_next"`
	Page    int              `json:"page"`
	Results []PlayerStanding `json:"results"`
}

type LeagueData struct {
	League      League    `json:"league"`
	Standings   Standings `json:"standings"`
	LastUpdated time.Time `json:"last_updated_data"`
}

func (ld LeagueData) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("🏆 *%s*\n\n", ld.League.Name))

	for _, s := range ld.Standings.Results {
		movementEmoji := getMovementEmoji(s.Rank, s.LastRank)
		medalEmoji := getMedalEmoji(s.Rank)

		line := fmt.Sprintf("%d. %s: %d %s", s.Rank, s.EntryName, s.Total, movementEmoji)

		if medalEmoji != "" {
			line = medalEmoji + " " + line
		}

		sb.WriteString(line + "\n")
	}

	return sb.String()
}

func getMovementEmoji(currentPos, prevPos int) string {
	switch {
	case currentPos < prevPos:
		return "🔼"
	case currentPos > prevPos:
		return "🔽"
	default:
		return "➖"
	}
}

func getMedalEmoji(pos int) string {
	switch {
	case pos == 1:
		return "🥇"
	case pos == 2:
		return "🥈"
	case pos == 3:
		return "🥉"
	default:
		return ""
	}
}
