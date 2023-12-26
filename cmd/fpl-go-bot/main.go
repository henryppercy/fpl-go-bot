package main

import (
	"fmt"
	"github.com/henryppercy/fpl-go-bot/pkg/fpl/client"
	"log"
	"os"
	"strconv"
)

func main() {
	leagueIds := os.Args[1:]

	for _, idStr := range leagueIds {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("Invalid league ID '%s': %v\n", idStr, err)
			continue
		}

		data, err := client.GetLeague(id)
		if err != nil {
			log.Printf("Error fetching league data for ID %d: %v\n", id, err)
			continue
		}

		fmt.Printf("%+v\n", data)
	}
}
