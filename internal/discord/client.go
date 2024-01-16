package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/henryppercy/fpl-go-bot/internal/logger"
)

func Send(id, content string) (*http.Response, error) {
	token := os.Getenv("BOT_TOKEN")
	url := fmt.Sprintf("%s/channels/%s/messages", os.Getenv("DISCORD_URL"), id)

	body := Body{
		Content: content,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	return res, nil
}

func DispatchMessage(msg string) error {
	channelId := os.Getenv("CHANNEL_ID")

	var res *http.Response
	res, err := Send(channelId, msg)
	if err != nil {
		logger.ErrorLogger.Printf("error sending discord message: %v\n", err)
		return err
	}
	
	if res != nil {
		logger.InfoLogger.Printf("discord message sent successfully with code: %d\n", res.StatusCode)
	}
	
	return nil
}

func DispatchInitMessage() error {
	err := DispatchMessage("*🤖 FPL Go Bot has been initialised*")

	if err != nil {
		return err
	}

	return nil
}

func DispatchUpdatedMessage() error {
	err := DispatchMessage("*🤖 FPL Go Bot has been updated*")

	if err != nil {
		return err
	}

	return nil
}
