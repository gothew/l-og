package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

var (
	ErrMissingValue = fmt.Errorf("missing value")
)

var _ Logger = &logger{}

// logger is a logger that implements Logger
type logger struct {
	w  io.Writer
	b  bytes.Buffer
	mu *sync.RWMutex

	level     int32
	formatter Formatter
	keyvals   []interface{}
}

// New return new logger
func New() Logger {
	l := &logger{
		b:     bytes.Buffer{},
		mu:    &sync.RWMutex{},
		level: int32(InfoLevel),
	}

	if l.w == nil {
		l.w = os.Stderr
	}

	l.SetOutput(l.w)
	l.SetLevel(Level(l.level))

	return l
}

func (l *logger) log(level Level, msg interface{}, keyvals ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	defer l.b.Reset()

	var kvs []interface{}

	if level != noLevel {
		kvs = append(kvs, LevelKey, level)
	}

	if msg != nil {
		m := fmt.Sprint(msg)
		kvs = append(kvs, MessageKey, m)
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

// SetLevel sets the current level.
func (l *logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	atomic.StoreInt32(&l.level, int32(level))
}

// GetLevel returns the current level.
func (l *logger) GetLevel() Level {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return Level(l.level)
}

func (l *logger) SetOutput(w io.Writer) {
	l.w = w
}

func (l *logger) Debug(msg interface{}, keyvals ...interface{}) {
	l.log(DebugLevel, msg, keyvals...)
}
