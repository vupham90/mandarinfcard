package mandarinfcard

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	respPayload, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http_status=%d, resp=%s", resp.StatusCode, string(respPayload))
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return nil
}
