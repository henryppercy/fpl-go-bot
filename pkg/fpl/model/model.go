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

	sb.WriteString(fmt.Sprintf("League ID: %d\n", ld.League.ID))
	sb.WriteString(fmt.Sprintf("League Name: %s\n", ld.League.Name))
	sb.WriteString(fmt.Sprintf("Created: %v\n", ld.League.Created))
	sb.WriteString("Standings:\n")

	for _, s := range ld.Standings.Results {
		sb.WriteString(fmt.Sprintf(" - %s: Pos %d (Prev Pos: %d, Total: %d)\n", s.EntryName, s.Rank, s.LastRank, s.Total))
	}

	return sb.String()
}
