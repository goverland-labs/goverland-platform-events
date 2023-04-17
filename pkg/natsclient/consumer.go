package natsclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"

	"github.com/goverland-labs/platform-events/events"
)

var (
	ErrGroupRequired = errors.New("group is required")
)

type EventHandler interface {
	RawHandler() events.RawMessageHandler
}

type Consumer struct {
	sub     *nats.Subscription
	group   string
	subject string
}

// NewConsumer creates nats QueueSubscribe with custom handler
// Group must be the name of service or package: core, feed, etc. It's allow handle messages in few consumers.
func NewConsumer(ctx context.Context, conn *nats.Conn, group, subject string, h EventHandler) (*Consumer, error) {
	if group == "" {
		return nil, ErrGroupRequired
	}

	if subject == "" {
		return nil, ErrSubjectRequired
	}

	js, err := conn.JetStream()
	if err != nil {
		return nil, err
	}

	err = getOrCreateStream(js, subject)
	if err != nil {
		return nil, err
	}

	subscription, err := js.QueueSubscribe(subject, group, func(msg *nats.Msg) {
		err := h.RawHandler()(msg.Data)
		if err != nil {
			// todo: think about nak with delay and timeouts
			err = msg.Nak()
			if err != nil {
				log.Error().Err(fmt.Errorf("[%s/%s]nack err: %w", group, subject, err))
				return
			}
		}

		if err := msg.Ack(); err != nil {
			log.Error().Err(fmt.Errorf("[%s/%s]nack err: %w", group, subject, err))
			return
		}
	}, nats.ManualAck(), nats.DeliverNew(), nats.Context(ctx))
	if err != nil {
		return nil, fmt.Errorf("queue subscriibe: %w", err)
	}

	return &Consumer{
		sub:     subscription,
		group:   group,
		subject: subject,
	}, nil
}

func (c *Consumer) Close() error {
	if err := c.sub.Unsubscribe(); err != nil {
		return fmt.Errorf("unsubscribe [%s/%s]: %w", c.subject, c.group, err)
	}

	if err := c.sub.Drain(); err != nil {
		return fmt.Errorf("drain [%s/%s]: %w", c.subject, c.group, err)
	}

	return nil
}
