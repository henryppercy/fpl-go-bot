package service

import (
	"encoding/json"
	"fmt"
	"io"
	"github.com/henryppercy/fpl-go-bot/pkg/fpl/model"
)

func MarshalLeagueData(body io.ReadCloser) (model.LeagueData, error) {
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
