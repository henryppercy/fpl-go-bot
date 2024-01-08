package fpl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func BootstrapApp() (LeagueBootstrap, error) {
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

func GetLeague(leagueId int) (LeagueData, error) {
	url := fmt.Sprintf("https://fantasy.premierleague.com/api/leagues-classic/%d/standings/", leagueId)
	response, err := http.Get(url)

	if err != nil {
		return LeagueData{}, fmt.Errorf("failed to fetch league data %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return LeagueData{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	return marshalLeagueData(response.Body)
}

func marshalLeagueData(body io.ReadCloser) (LeagueData, error) {
	var leagueData LeagueData

	responseData, err := io.ReadAll(body)
	if err != nil {
		return LeagueData{}, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(responseData, &leagueData)
	if err != nil {
		return LeagueData{}, fmt.Errorf("failed to marshal response body: %w", err)
	}

	return leagueData, nil
}
