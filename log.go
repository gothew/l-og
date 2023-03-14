package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrMissingValue = fmt.Errorf("missing value")
)

type LoggerOptions = func(*logger)

var _ Logger = &logger{}

// logger is a logger that implements Logger
type logger struct {
	w  io.Writer
	b  bytes.Buffer
	mu *sync.RWMutex

	level      int32
	prefix     string
	timeFunc   TimeFunction
	timeFormat string
	formatter  Formatter
	keyvals    []interface{}

	timestamp bool
	notStyles bool
}

// New return new logger
func New(opts ...LoggerOptions) Logger {
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

	if l.timeFunc == nil {
		l.timeFunc = time.Now
	}

	if l.timeFormat == "" {
		l.timeFormat = DefaultTimeFormat
	}

	return l
}

func (l *logger) log(level Level, msg interface{}, keyvals ...interface{}) {

	// check if the level is allowed
	if atomic.LoadInt32(&l.level) > int32(level) {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	defer l.b.Reset()

	var kvs []interface{}
	if !l.timestamp {
		kvs = append(kvs, TimestampKey, l.timeFunc())
	}

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

// SetReportTimestamp sets whether the timestamp should be reported.
func (l *logger) SetReportTimestamp(report bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timestamp = report
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

// SetTimeFormat sets the time format.
func (l *logger) SetTimeFormat(format string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timeFormat = format
}

// SetTimeFunction sets the time function.
func (l *logger) SetTimeFunction(f TimeFunction) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timeFunc = f
}

func (l *logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.w = w

	if !isTerminal(w) {
		// This only affects the TextFormatter
		l.notStyles = true
	}
}

func (l *logger) Debug(msg interface{}, keyvals ...interface{}) {
	l.log(DebugLevel, msg, keyvals...)
}

func (l *logger) Info(msg interface{}, keyvals ...interface{}) {
	l.log(InfoLevel, msg, keyvals...)
}

func (l *logger) Warn(msg interface{}, kevals ...interface{}) {
  l.log(WarnLevel, msg, kevals...)
}
