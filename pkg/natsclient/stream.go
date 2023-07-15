package natsclient

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nats-io/nats.go"
)

func getOrCreateStream(js nats.JetStreamContext, subject string) (*nats.StreamInfo, error) {
	streamName := buildStreamName(subject)
	s, err := js.StreamInfo(streamName)
	if err == nil {
		return s, nil
	}

	if err != nil && !errors.Is(err, nats.ErrStreamNotFound) {
		return nil, fmt.Errorf("get stream info [%s]: %v", streamName, err)
	}

	s, err = js.AddStream(&nats.StreamConfig{
		Name:      streamName,
		Subjects:  []string{subject},
		Retention: nats.LimitsPolicy,
		Discard:   nats.DiscardOld,
		Storage:   nats.FileStorage,
		MaxAge:    StreamDefaultMaxAge,
	})

	if err != nil {
		return nil, fmt.Errorf("add stream: %w", err)
	}

	return s, nil
}

func buildStreamName(subject string) string {
	return strings.Replace(fmt.Sprintf("str_%s", subject), ".", "_", -1)
}
