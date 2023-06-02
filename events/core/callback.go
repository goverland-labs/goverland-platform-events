package core

import (
	"encoding/json"

	"github.com/goverland-labs/platform-events/events"
)

const (
	SubjectCallback = "core.callback"
)

type CallbackPayload struct {
	WebhookURL string          `json:"webhook_url"`
	Body       json.RawMessage `json:"body"`
}

type CallbackHandler func(payload CallbackPayload) error

func (h CallbackHandler) RawHandler() events.RawMessageHandler {
	return func(raw []byte) error {
		var d CallbackPayload
		if err := json.Unmarshal(raw, &d); err != nil {
			return err
		}

		return h(d)
	}
}
