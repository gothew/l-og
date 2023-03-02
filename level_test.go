package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultLevel(t *testing.T) {
	var level Level
	assert.Equal(t, InfoLevel, level)
}

func TestParseLevel(t *testing.T) {
	testCases := []struct {
		name     string
		level    string
		expLevel Level
	}{
		{
			name:     "Parse debug",
			level:    "debug",
			expLevel: DebugLevel,
		},
		{
			name:     "Parse info",
			level:    "info",
			expLevel: InfoLevel,
		},
	}

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      assert.Equal(t, tc.expLevel, ParseLevel(tc.level))
    })
  }
}