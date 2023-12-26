package client

import (
	"github.com/henryppercy/fpl-go-bot/internal/utils"
	"fmt"
	"os"
)

func GetLeague() {
	data, err := os.ReadFile("tmp/league_test.json")
	utils.Check(err)
	fmt.Println(string(data))
}
