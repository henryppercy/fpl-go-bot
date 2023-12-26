package main

import (
	"fmt"
	"github.com/henryppercy/fpl-go-bot/pkg/fpl/client"
	whatsAppClient "github.com/henryppercy/fpl-go-bot/pkg/notify/client"
	"log"
	"net/http"
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

		var res *http.Response
		res, err = whatsAppClient.Send("", data.String())
		if err != nil {
			log.Printf("Error fetching sending whatsapp message: %v\n", err)
		}

		if res != nil {
			fmt.Printf("WhatsApp sent successfully with code: %d\n", res.StatusCode)
		}

		fmt.Printf("%+v\n", data)
	}
}
