package natsclient

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"

	"github.com/goverland-labs/platform-events/events"
)

const rateLimit = 1 * 8 * 1024 * 1024 // 3MiB TODO: Move it to options

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
func NewConsumer(ctx context.Context, conn *nats.Conn, group, subject string, h EventHandler, maxAckPending ...int) (*Consumer, error) {
	// TODO: Passing MaxAckPending as an option
	maxPending := 0
	if len(maxAckPending) > 0 {
		maxPending = maxAckPending[0]
	}

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

	stream, err := getOrCreateStream(js, subject)
	if err != nil {
		return nil, err
	}

	consumerName := buildConsumerName(group, subject)

	var consumer *nats.ConsumerInfo
	for c := range js.Consumers(stream.Config.Name) {
		if c.Name == consumerName {
			consumer = c
		}
	}

	if consumer == nil {
		consumer, err = js.AddConsumer(stream.Config.Name, &nats.ConsumerConfig{
			Durable:        consumerName,
			Name:           consumerName,
			DeliverPolicy:  nats.DeliverAllPolicy,
			AckPolicy:      nats.AckExplicitPolicy,
			AckWait:        time.Minute,
			DeliverSubject: fmt.Sprintf("deliver.%s", consumerName),
			DeliverGroup:   group,
			FilterSubject:  subject,
			MaxAckPending:  maxPending,
			RateLimit:      rateLimit,
		})
		if err != nil {
			return nil, fmt.Errorf("unable to create consumer '%s': %w", group, err)
		}
	}

	opts := []nats.SubOpt{
		nats.Durable(consumerName),
		nats.ManualAck(),
		nats.DeliverAll(),
		nats.Context(ctx),
		nats.AckWait(time.Minute),
	}

	if maxPending > 0 {
		opts = append(opts, nats.MaxAckPending(maxPending))
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
	}, opts...)
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
	if err := c.sub.Drain(); err != nil {
		return fmt.Errorf("drain [%s/%s]: %w", c.subject, c.group, err)
	}

	return nil
}

func buildConsumerName(group, subject string) string {
	return strings.Replace(fmt.Sprintf("consumer_%s_%s", group, subject), ".", "_", -1)
}
