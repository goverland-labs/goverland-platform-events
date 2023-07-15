package events

import (
	"encoding/json"
)

type RawHandler func([]byte) error

type Handler[T any] func(payload T) error

func (h Handler[T]) RawHandler() RawHandler {
	return func(raw []byte) error {
		var payload T

		if err := json.Unmarshal(raw, &payload); err != nil {
			return err
		}

		return h(payload)
	}
}
