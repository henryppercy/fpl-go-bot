package api

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/henryppercy/fpl-go-bot/internal/fpl"
	"github.com/henryppercy/fpl-go-bot/internal/logger"
)

func (s *Server) handleGetDeadline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	lb, err := fpl.GetBootstrap()
	if err != nil {
		logger.ErrorLogger.Println("unable to retrieve league bootstrap data")
		return
	}

	event, exists := lb.GetNextEvent()

	if exists {
		date := event.DeadlineTime.Format("Monday 02 Jan, 15:04")
		err := json.NewEncoder(w).Encode(date)
		if err != nil {
			logger.ErrorLogger.Println("error returning deadline json")
		}
	} else {
		err := json.NewEncoder(w).Encode("deadline not available")
		if err != nil {
			logger.ErrorLogger.Println("error returning deadline json")
		}
	}
}

func (s *Server) handleGetLeague(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(os.Getenv("LEAGUE_ID"))

	league, err := fpl.GetLeague(id)
	if err != nil {
		logger.ErrorLogger.Printf("error fetching league data for ID %d: %v\n", id, err)
	}
	
	err = json.NewEncoder(w).Encode(league)
	if err != nil {
		logger.ErrorLogger.Println("error returning league json")
	}
}
