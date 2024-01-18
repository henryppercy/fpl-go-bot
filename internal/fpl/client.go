package fpl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/henryppercy/fpl-go-bot/internal/fpl/models"
)

func GetEventStatus() (models.EventStatus, error) {
	url := "https://fantasy.premierleague.com/api/event-status/"

	var eventStatus models.EventStatus
	response, err := http.Get(url)

	if err != nil {
		return models.EventStatus{}, fmt.Errorf("failed to fetch league data %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.EventStatus{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return models.EventStatus{}, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(responseData, &eventStatus)
	if err != nil {
		return models.EventStatus{}, fmt.Errorf("failed to marshal response body: %w", err)
	}

	return eventStatus, nil
}

func GetLeague(leagueId int) (models.LeagueData, error) {
	url := fmt.Sprintf("https://fantasy.premierleague.com/api/leagues-classic/%d/standings/", leagueId)
	response, err := http.Get(url)

	if err != nil {
		return models.LeagueData{}, fmt.Errorf("failed to fetch league data %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.LeagueData{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	var leagueData models.LeagueData

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return models.LeagueData{}, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(responseData, &leagueData)
	if err != nil {
		return models.LeagueData{}, fmt.Errorf("failed to marshal response body: %w", err)
	}

	return leagueData, nil
}

func GetBootstrap() (models.LeagueBootstrap, error) {
	url := "https://fantasy.premierleague.com/api/bootstrap-static/"

	var leagueData models.LeagueBootstrap
	response, err := http.Get(url)

	if err != nil {
		return models.LeagueBootstrap{}, fmt.Errorf("failed to fetch league data %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.LeagueBootstrap{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return models.LeagueBootstrap{}, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(responseData, &leagueData)
	if err != nil {
		return models.LeagueBootstrap{}, fmt.Errorf("failed to marshal response body: %w", err)
	}

	return leagueData, nil
}
