package mandarinfcard

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

const (
	slackUrl = "https://hooks.slack.com/services/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"
)

func NotiSend(ctx context.Context, url string, msg string) error {
	type Payload struct {
		Text string `json:"text"`
	}

	data := Payload{
		Text: msg,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", slackUrl, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return nil
}
