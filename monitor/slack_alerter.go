package monitor

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/log"
)

const reponseBuffer = 2048

type slackAlerter struct {
	url    string
	client *http.Client
}

func formatMessage(header, body string) string {
	headerSection := map[string]interface{}{
		"type": "header",
		"text": map[string]interface{}{
			"type":  "plain_text",
			"text":  header,
			"emoji": true,
		},
	}

	bodySection := map[string]interface{}{
		"type": "section",
		"text": map[string]interface{}{
			"type":  "plain_text",
			"text":  body,
			"emoji": true,
		},
	}

	messageBlock := map[string]interface{}{
		"blocks": []interface{}{
			headerSection,
			bodySection,
		},
	}

	message, _ := json.Marshal(messageBlock)
	return string(message)
}

func (alerter *slackAlerter) Alert(header, body string) {
	message := formatMessage(header, body)

	request, err := http.NewRequest("POST", alerter.url, strings.NewReader(message))
	if err != nil {
		log.Error("Failed to send Slack alert", "err", err)
		return
	}

	response, err := alerter.client.Do(request)
	if err != nil {
		log.Error("Failed to send HTTP request", "err", err)
		return
	}

	if response.StatusCode >= 400 {
		responseBody := make([]byte, reponseBuffer)
		response.Body.Read(responseBody)
		log.Error("Error response from server", "status", response.StatusCode, "body", responseBody)
		return
	}
}

func NewSlackAlert() *slackAlerter {
	slackUrl := os.Getenv("SLACK_WEBHOOK_URL")
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("invalid redirect")
		},
	}

	return &slackAlerter{
		url:    slackUrl,
		client: client,
	}
}
