package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
