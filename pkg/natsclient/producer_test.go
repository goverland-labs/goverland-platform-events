package natsclient

import (
	"context"
	"fmt"
	"testing"

	"github.com/nats-io/nats-server/v2/server"
	natsserver "github.com/nats-io/nats-server/v2/test"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/require"
)

const testPort = 8369

func RunNatsServer(port int) *server.Server {
	opts := natsserver.DefaultTestOptions
	opts.Port = port
	opts.JetStream = true

	return natsserver.RunServer(&opts)
}

func TestUnitNewProducer(t *testing.T) {
	for name, tc := range map[string]struct {
		subject string
		err     error
	}{
		"correct init": {
			subject: "subject",
			err:     nil,
		},
		"empty subject": {
			subject: "",
			err:     ErrSubjectRequired,
		},
		"subject contains spaces": {
			subject: "test name",
			err:     ErrInvalidChars,
		},
	} {
		t.Run(name, func(t *testing.T) {
			s := RunNatsServer(testPort)
			defer s.Shutdown()

			sUrl := fmt.Sprintf("nats://127.0.0.1:%d", testPort)
			nc, err := nats.Connect(sUrl)
			require.NoError(t, err)
			defer nc.Close()

			_, err = NewProducer(nc, tc.subject)
			if tc.err == nil {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tc.err.Error())
			}
		})
	}

	t.Run("allow few producers with same subject", func(t *testing.T) {
		s := RunNatsServer(testPort)
		defer s.Shutdown()

		sUrl := fmt.Sprintf("nats://127.0.0.1:%d", testPort)
		nc, err := nats.Connect(sUrl)
		require.NoError(t, err)
		defer nc.Close()

		subject := "test.subject"
		_, err = NewProducer(nc, subject)
		require.NoError(t, err)
		_, err = NewProducer(nc, subject)
		require.NoError(t, err)
		_, err = NewProducer(nc, subject)
		require.NoError(t, err)
	})
}

func TestUnitPublish(t *testing.T) {
	s := RunNatsServer(testPort)
	defer s.Shutdown()

	sUrl := fmt.Sprintf("nats://127.0.0.1:%d", testPort)
	nc, err := nats.Connect(sUrl)
	require.NoError(t, err)
	defer nc.Close()

	subject := "test.subject"
	pr, err := NewProducer(nc, subject)
	require.NoError(t, err)

	t.Run("publish string", func(t *testing.T) {
		err = pr.PublishData(context.Background(), []byte("string"))
		require.NoError(t, err)
	})

	t.Run("publish object", func(t *testing.T) {
		type TestData struct {
			Data string `json:"data"`
		}

		err = pr.PublishJSON(context.Background(), TestData{Data: "information"})
		require.NoError(t, err)
	})
}
