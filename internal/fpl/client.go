package fpl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getEventStatus() (EventStatus, error) {
	url := "https://fantasy.premierleague.com/api/event-status/"

	var eventStatus EventStatus
	response, err := http.Get(url)

	if err != nil {
		return EventStatus{}, fmt.Errorf("failed to fetch league data %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return EventStatus{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return EventStatus{}, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(responseData, &eventStatus)
	if err != nil {
		return EventStatus{}, fmt.Errorf("failed to marshal response body: %w", err)
	}

	return eventStatus, nil
}

func getBootstrap() (LeagueBootstrap, error) {
	url := "https://fantasy.premierleague.com/api/bootstrap-static/"

	var leagueData LeagueBootstrap
	response, err := http.Get(url)

	if err != nil {
		return LeagueBootstrap{}, fmt.Errorf("failed to fetch league data %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return LeagueBootstrap{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return LeagueBootstrap{}, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(responseData, &leagueData)
	if err != nil {
		return LeagueBootstrap{}, fmt.Errorf("failed to marshal response body: %w", err)
	}

	return leagueData, nil
}

func getLeague(leagueId int) (LeagueData, error) {
	url := fmt.Sprintf("https://fantasy.premierleague.com/api/leagues-classic/%d/standings/", leagueId)
	response, err := http.Get(url)

	if err != nil {
		return LeagueData{}, fmt.Errorf("failed to fetch league data %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return LeagueData{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	var leagueData LeagueData

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return LeagueData{}, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(responseData, &leagueData)
	if err != nil {
		return LeagueData{}, fmt.Errorf("failed to marshal response body: %w", err)
	}

	return leagueData, nil
}
