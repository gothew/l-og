package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

var (
	ErrMissingValue = fmt.Errorf("missing value")
)

var _ Logger = &logger{}

// logger is a logger that implements Logger
type logger struct {
	w         io.Writer
	b         bytes.Buffer
	mu        *sync.RWMutex
	formatter Formatter
	keyvals   []interface{}
}

// New return new logger
func New() Logger {
	l := &logger{
		b:  bytes.Buffer{},
		mu: &sync.RWMutex{},
	}

	if l.w == nil {
		l.w = os.Stderr
	}

	l.SetOutput(l.w)

	return l
}

func (l *logger) log(msg interface{}, keyvals ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	defer l.b.Reset()

	var kvs []interface{}

	if msg != nil {
		m := fmt.Sprint(msg)
		kvs = append(kvs, msgKey, m)
	}

	// append logger fields
	kvs = append(kvs, l.keyvals...)
	if len(l.keyvals)%2 != 0 {
		kvs = append(kvs, ErrMissingValue)
	}

	switch l.formatter {
	case LogftmFormatter:
    // FIXME: use in default
    l.textFormatter(kvs...)
	default:
	}

  _, _ = l.w.Write(l.b.Bytes())
}

func (l *logger) SetOutput(w io.Writer) {
	l.w = w
}

func (l *logger) Debug(msg interface{}, keyvals ...interface{}) {
	l.log(msg, keyvals...)
}
