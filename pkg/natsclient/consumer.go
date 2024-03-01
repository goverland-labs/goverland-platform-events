package natsclient

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"

	"github.com/goverland-labs/goverland-platform-events/events"
)

const (
	KiB = 8 * 1024
	MiB = KiB * 1024
	GiB = MiB * 1024
)

const (
	consumerActionAck  = "ack"
	consumerActionNack = "nack"
)

var ErrGroupRequired = errors.New("group is required")

type Consumer[T any] struct {
	sub     *nats.Subscription
	group   string
	subject string
}

type ConsumerOpt func(*nats.ConsumerConfig)

func WithRateLimit(limit uint64) ConsumerOpt {
	return func(cfg *nats.ConsumerConfig) {
		cfg.RateLimit = limit
	}
}

func WithMaxAckPending(count int) ConsumerOpt {
	return func(cfg *nats.ConsumerConfig) {
		cfg.MaxAckPending = count
	}
}

func WithAckWait(wait time.Duration) ConsumerOpt {
	return func(cfg *nats.ConsumerConfig) {
		cfg.AckWait = wait
	}
}

func WithDeliverPolicy(policy nats.DeliverPolicy) ConsumerOpt {
	return func(cfg *nats.ConsumerConfig) {
		cfg.DeliverPolicy = policy
	}
}

func WithAckPolicy(policy nats.AckPolicy) ConsumerOpt {
	return func(cfg *nats.ConsumerConfig) {
		cfg.AckPolicy = policy
	}
}

// NewConsumer creates nats QueueSubscribe with custom handler
// Group must be the name of service or package: core, feed, etc. It's allow handle messages in few consumers.
func NewConsumer[T any](ctx context.Context, conn *nats.Conn, group, subject string, h events.Handler[T], opts ...ConsumerOpt) (*Consumer[T], error) {
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

	cfg := &nats.ConsumerConfig{
		Durable:        consumerName,
		Name:           consumerName,
		DeliverPolicy:  nats.DeliverAllPolicy,
		AckPolicy:      nats.AckExplicitPolicy,
		DeliverSubject: fmt.Sprintf("deliver.%s", consumerName),
		DeliverGroup:   group,
		FilterSubject:  subject,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	if consumer == nil {
		consumer, err = js.AddConsumer(stream.Config.Name, cfg)
		if err != nil {
			return nil, fmt.Errorf("unable to create consumer '%s': %w", group, err)
		}
	}

	// TODO: Think about updating consumer settings if it's different

	subOpts := []nats.SubOpt{
		nats.Durable(consumer.Name),
		nats.ManualAck(),
		nats.DeliverAll(),
		nats.Context(ctx),
		nats.AckWait(cfg.AckWait),
	}

	if cfg.MaxAckPending > 0 {
		subOpts = append(subOpts, nats.MaxAckPending(cfg.MaxAckPending))
	}

	subscription, err := js.QueueSubscribe(subject, group, func(msg *nats.Msg) {
		var (
			start  = time.Now()
			action = consumerActionAck
		)

		defer func() {
			CollectConsumerMetric(subject, action, err, time.Since(start).Seconds())
		}()

		err = h.RawHandler()(msg.Data)
		if err != nil {
			action = consumerActionNack
			// todo: think about nak with delay and timeouts
			if err = msg.NakWithDelay(time.Second); err != nil {
				log.Error().Err(fmt.Errorf("[%s/%s] nack err: %w", group, subject, err))
			}

			return
		}

		if err = msg.AckSync(); err != nil {
			log.Error().Err(fmt.Errorf("[%s/%s] ack err: %w", group, subject, err))
			return
		}
	}, subOpts...)
	if err != nil {
		return nil, fmt.Errorf("queue subscriibe: %w", err)
	}

	return &Consumer[T]{
		sub:     subscription,
		group:   group,
		subject: subject,
	}, nil
}

func (c *Consumer[T]) Close() error {
	if err := c.sub.Drain(); err != nil {
		return fmt.Errorf("drain [%s/%s]: %w", c.subject, c.group, err)
	}

	return nil
}

func buildConsumerName(group, subject string) string {
	return strings.Replace(fmt.Sprintf("consumer_%s_%s", group, subject), ".", "_", -1)
}
