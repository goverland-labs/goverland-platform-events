package natsclient

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/goverland-labs/platform-events/events"
)

func TestUnitNewConsumer(t *testing.T) {
	t.Run("run one consumer", func(t *testing.T) {
		s, nc, err := runNatsServer(testPort)
		require.NoError(t, err)
		defer func() {
			s.Shutdown()
			nc.Close()
		}()

		c, err := NewConsumer(context.Background(), nc, "group", "subject", nil)
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

		c1, err := NewConsumer(context.Background(), nc, "group", "subject", nil)
		require.NoError(t, err)
		require.NotNil(t, c1)

		c2, err := NewConsumer(context.Background(), nc, "group", "subject", nil)
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
		var handler events.ProposalHandler = func(payload events.ProposalPayload) error {
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
			_ = pl.PublishJSON(context.Background(), &events.ProposalPayload{ID: "id-1"})
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
		var handler events.ProposalHandler = func(payload events.ProposalPayload) error {
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
			_ = pl.PublishJSON(context.Background(), &events.ProposalPayload{ID: "id-1"})
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
		var handler events.ProposalHandler = func(payload events.ProposalPayload) error {
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
			_ = pl.PublishJSON(context.Background(), &events.ProposalPayload{ID: "id-1"})
		}

		<-time.After(time.Millisecond * 100)
		assert.Equal(t, msgs*2, int(cnt))
	})

	t.Run("handle only new messages", func(t *testing.T) {
		s, nc, err := runNatsServer(testPort)
		require.NoError(t, err)
		defer func() {
			s.Shutdown()
			nc.Close()
		}()

		subject := "consume.msg.new"
		pl, err := NewProducer(nc, subject)
		require.NoError(t, err)
		msgs := 500
		for i := 0; i < msgs; i++ {
			_ = pl.PublishJSON(context.Background(), &events.ProposalPayload{ID: "id-1"})
		}

		for i := 0; i < msgs; i++ {
			_ = pl.PublishJSON(context.Background(), &events.ProposalPayload{ID: "id-1"})
		}

		<-time.After(time.Millisecond * 100)

		var cnt int64 = 0
		var handler events.ProposalHandler = func(payload events.ProposalPayload) error {
			atomic.AddInt64(&cnt, 1)
			return nil
		}

		group := fmt.Sprintf("group-%d", rand.Int())
		_, err = NewConsumer(context.Background(), nc, group, subject, handler)
		require.NoError(t, err)

		for i := 0; i < msgs; i++ {
			_ = pl.PublishJSON(context.Background(), &events.ProposalPayload{ID: "id-1"})
		}

		<-time.After(time.Millisecond * 100)
		assert.Equal(t, msgs, int(cnt))
	})

	t.Run("handle only new messages", func(t *testing.T) {
		s, nc, err := runNatsServer(testPort)
		require.NoError(t, err)
		defer func() {
			s.Shutdown()
			nc.Close()
		}()

		subject := "consume.msg.error.handler"
		var cnt int64 = 0
		var handler events.ProposalHandler = func(payload events.ProposalPayload) error {
			atomic.AddInt64(&cnt, 1)
			if cnt < 3 {
				return errors.New("unexpected error")
			}

			return nil
		}

		group := fmt.Sprintf("group-%d", rand.Int())
		_, err = NewConsumer(context.Background(), nc, group, subject, handler)
		require.NoError(t, err)

		pl, err := NewProducer(nc, subject)
		require.NoError(t, err)
		_ = pl.PublishJSON(context.Background(), &events.ProposalPayload{ID: "id-1"})

		<-time.After(time.Millisecond * 100)

		// checks how many attempts consumer processed
		assert.Equal(t, 3, int(cnt))
	})
}
