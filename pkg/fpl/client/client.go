package client

import (
	"fmt"
	"github.com/henryppercy/fpl-go-bot/pkg/fpl/model"
	"github.com/henryppercy/fpl-go-bot/pkg/fpl/service"
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

	return service.MarshalLeagueData(response.Body)
}
