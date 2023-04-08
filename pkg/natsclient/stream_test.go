package natsclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitBuildStreamName(t *testing.T) {
	for name, tc := range map[string]struct {
		subject  string
		expected string
	}{
		"correct": {
			subject:  "test",
			expected: "str_test",
		},
		"replace dots to underscores": {
			subject:  "test.sub1",
			expected: "str_test_sub1",
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := buildStreamName(tc.subject)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
