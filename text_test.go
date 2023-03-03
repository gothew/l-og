package log

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := New()

	cases := []struct {
		name     string
		expected string
		msg      string
		kvs      []interface{}
		f        func(msg interface{}, kvs ...interface{})
	}{
		{
			name:     "ignore message",
			expected: "",
			msg:      "this is a debug message",
			kvs:      nil,
			f:        logger.Debug,
		},
	}

	for _, c := range cases {
		buf.Reset()
		t.Run(c.name, func(t *testing.T) {
			c.f(c.msg, c.kvs...)
			assert.Equal(t, c.expected, buf.String())
		})
	}
}
