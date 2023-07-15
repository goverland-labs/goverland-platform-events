package natsclient

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/goverland-labs/platform-events/events/aggregator"
)

func TestUnitNewConsumer(t *testing.T) {
	t.Run("run one consumer", func(t *testing.T) {
		s, nc, err := runNatsServer(testPort)
		require.NoError(t, err)
		defer func() {
			s.Shutdown()
			nc.Close()
		}()

		c, err := NewConsumer[interface{}](context.Background(), nc, "group", "subject", nil)
		require.NoError(t, err)
		require.NotNil(t, c)
	})

	t.Run("run consumers group", func(t *testing.T) {
		s, nc, err := runNatsServer(testPort)
		require.NoError(t, err)
		defer func() {
			s.Shutdown()
			nc.Close()
		}()

		c1, err := NewConsumer[interface{}](context.Background(), nc, "group", "subject", nil)
		require.NoError(t, err)
		require.NotNil(t, c1)

		c2, err := NewConsumer[interface{}](context.Background(), nc, "group", "subject", nil)
		require.NoError(t, err)
		require.NotNil(t, c2)
	})
}

func TestUnitConsumeMsg(t *testing.T) {
	t.Run("run one consumer", func(t *testing.T) {
		s, nc, err := runNatsServer(testPort)
		require.NoError(t, err)
		defer func() {
			s.Shutdown()
			nc.Close()
		}()

		cnt := 0
		var handler aggregator.ProposalHandler = func(payload aggregator.ProposalPayload) error {
			cnt++
			return nil
		}

		subject := "consume.msg.one"
		_, err = NewConsumer(context.Background(), nc, "group-1", subject, handler)
		require.NoError(t, err)

		pl, err := NewProducer(nc, subject)
		require.NoError(t, err)

		msgs := 5
		for i := 0; i < msgs; i++ {
			_ = pl.PublishJSON(context.Background(), &aggregator.ProposalPayload{ID: "id-1"})
		}

		<-time.After(time.Millisecond * 100)
		assert.Equal(t, msgs, cnt)
	})

	t.Run("run consumer groups", func(t *testing.T) {
		s, nc, err := runNatsServer(testPort)
		require.NoError(t, err)
		defer func() {
			s.Shutdown()
			nc.Close()
		}()

		var cnt int64 = 0
		var handler aggregator.ProposalHandler = func(payload aggregator.ProposalPayload) error {
			atomic.AddInt64(&cnt, 1)
			return nil
		}

		subject := "consume.msg.similar.name"
		_, err = NewConsumer(context.Background(), nc, "group-2", subject, handler)
		require.NoError(t, err)

		_, err = NewConsumer(context.Background(), nc, "group-2", subject, handler)
		require.NoError(t, err)

		pl, err := NewProducer(nc, subject)
		require.NoError(t, err)

		msgs := 500
		for i := 0; i < msgs; i++ {
			_ = pl.PublishJSON(context.Background(), &aggregator.ProposalPayload{ID: "id-1"})
		}

		<-time.After(time.Millisecond * 100)
		assert.Equal(t, msgs, int(cnt))
	})

	t.Run("run two different consumer groups", func(t *testing.T) {
		s, nc, err := runNatsServer(testPort)
		require.NoError(t, err)
		defer func() {
			s.Shutdown()
			nc.Close()
		}()

		var cnt int64 = 0
		var handler aggregator.ProposalHandler = func(payload aggregator.ProposalPayload) error {
			atomic.AddInt64(&cnt, 1)
			return nil
		}

		subject := "consume.msg.different.names"
		_, err = NewConsumer(context.Background(), nc, "group-3", subject, handler)
		require.NoError(t, err)

		_, err = NewConsumer(context.Background(), nc, "group-4", subject, handler)
		require.NoError(t, err)

		pl, err := NewProducer(nc, subject)
		require.NoError(t, err)

		msgs := 500
		for i := 0; i < msgs; i++ {
			_ = pl.PublishJSON(context.Background(), &aggregator.ProposalPayload{ID: "id-1"})
		}

		<-time.After(time.Millisecond * 100)
		assert.Equal(t, msgs*2, int(cnt))
	})
}
