package client

import (
	"encoding/json"
	"fmt"
	"github.com/henryppercy/fpl-go-bot/pkg/fpl/model"
	"io"
	"net/http"
)

func GetLeague(leagueId int) (model.LeagueData, error) {
	url := fmt.Sprintf("https://fantasy.premierleague.com/api/leagues-classic/%d/standings/", leagueId)
	response, err := http.Get(url)

	if err != nil {
		return model.LeagueData{}, fmt.Errorf("failed to fetch league data %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return model.LeagueData{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	return marshalLeagueData(response.Body)
}

func marshalLeagueData(body io.ReadCloser) (model.LeagueData, error) {
	var leagueData model.LeagueData

	responseData, err := io.ReadAll(body)
	if err != nil {
		return model.LeagueData{}, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(responseData, &leagueData)
	if err != nil {
		return model.LeagueData{}, fmt.Errorf("failed to marshal response body: %w", err)
	}

	return leagueData, nil
}
