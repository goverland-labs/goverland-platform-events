package natsclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/nats-io/nats.go"
)

var (
	ErrSubjectRequired = errors.New("subject is required")
	ErrInvalidChars    = errors.New("subject contains spaces")
)

type (
	Producer struct {
		js      nats.JetStreamContext
		subject string
	}
)

// NewProducer ...
// fixme: add description
func NewProducer(conn *nats.Conn, subject string) (*Producer, error) {
	if subject == "" {
		return nil, ErrSubjectRequired
	}

	if strings.ContainsAny(subject, " ") {
		return nil, ErrInvalidChars
	}

	// todo: allow provide options for creating JetStream
	js, err := conn.JetStream()
	if err != nil {
		return nil, fmt.Errorf("prepare jet stream: %w", err)
	}

	if err = getOrCreateStream(js, subject); err != nil {
		return nil, fmt.Errorf("prepare stream: %w", err)
	}

	return &Producer{
		js:      js,
		subject: subject,
	}, nil
}

func (p *Producer) PublishData(ctx context.Context, data []byte) error {
	return p.publish(ctx, data)
}

func (p *Producer) PublishJSON(ctx context.Context, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("marshal json: %v", err)
	}

	return p.publish(ctx, data)
}

func (p *Producer) publish(ctx context.Context, data []byte) error {
	if _, err := p.js.Publish(p.subject, data, nats.Context(ctx)); err != nil {
		return fmt.Errorf("publish data: %w", err)
	}

	return nil
}
