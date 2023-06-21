package natsclient

import (
	"context"
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	mu        sync.Mutex
	conn      *nats.Conn
	producers map[string]*Producer
}

func NewPublisher(nc *nats.Conn) (*Publisher, error) {
	p := &Publisher{
		conn:      nc,
		producers: make(map[string]*Producer, 0),
	}

	return p, nil
}

func (p *Publisher) PublishJSON(ctx context.Context, subject string, obj any) error {
	pr, err := p.getProducer(subject)
	if err != nil {
		return fmt.Errorf("get producer: %w", err)
	}

	return pr.PublishJSON(ctx, obj)
}

func (p *Publisher) getProducer(subject string) (*Producer, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	pr, ok := p.producers[subject]
	if ok {
		return pr, nil
	}

	pr, err := NewProducer(p.conn, subject)
	if err != nil {
		return nil, fmt.Errorf("create producer: %w", err)
	}

	p.producers[subject] = pr

	return pr, nil
}
