package natsclient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

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
			s, nc, err := runNatsServer(testPort)
			require.NoError(t, err)
			defer func() {
				s.Shutdown()
				nc.Close()
			}()

			_, err = NewProducer(nc, tc.subject)
			if tc.err == nil {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tc.err.Error())
			}
		})
	}

	t.Run("allow few producers with same subject", func(t *testing.T) {
		s, nc, err := runNatsServer(testPort)
		require.NoError(t, err)
		defer func() {
			s.Shutdown()
			nc.Close()
		}()

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
	s, nc, err := runNatsServer(testPort)
	require.NoError(t, err)
	defer func() {
		s.Shutdown()
		nc.Close()
	}()

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
