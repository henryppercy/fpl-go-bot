package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/henryppercy/fpl-go-bot/pkg/notify/model"
	"net/http"
)

func Send(id, body string) (*http.Response, error) {
	url := "https://gate.whapi.cloud/messages/text"

	message := model.Message{
		To:   id,
		Body: body,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", "Bearer ")

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
