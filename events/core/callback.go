package core

import (
	"encoding/json"

	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	SubjectCallback = "core.callback"
)

type CallbackPayload struct {
	WebhookURL string          `json:"webhook_url"`
	Body       json.RawMessage `json:"body"`
}

type CallbackHandler = events.Handler[CallbackPayload]
