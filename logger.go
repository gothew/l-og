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

type LoggerOptions = func(*Logger)

// Logger is a logger that implements Logger.
type Logger struct {
	w  io.Writer
	b  bytes.Buffer
	mu *sync.RWMutex

  isDiscard uint32

	level      int32
	timeFunc   TimeFunction
	timeFormat string
	formatter  Formatter
	keyvals    []interface{}

	reportTimestamp bool
	notStyles bool
}

func (l *Logger) log(level Level, msg interface{}, keyvals ...interface{}) {

	// check if the level is allowed
	if atomic.LoadInt32(&l.level) > int32(level) {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	defer l.b.Reset()

	var kvs []interface{}
	if l.reportTimestamp {
		kvs = append(kvs, TimestampKey, l.timeFunc())
	}

	if level != noLevel {
		kvs = append(kvs, LevelKey, level)
	}

	if msg != nil {
		m := fmt.Sprint(msg)
		kvs = append(kvs, MessageKey, m)
	}

	// append Logger fields
	kvs = append(kvs, l.keyvals...)
	if len(l.keyvals)%2 != 0 {
		kvs = append(kvs, ErrMissingValue)
	}

	switch l.formatter {
	case LogftmFormatter:
		l.textFormatter(kvs...)
  case JSONFormatter:
    l.jsonFormatter(kvs...)
	default:
		l.textFormatter(kvs...)
	}

	_, _ = l.w.Write(l.b.Bytes())
}

// SetReportTimestamp sets whether the timestamp should be reported.
func (l *Logger) SetReportTimestamp(report bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.reportTimestamp = report
}

// SetLevel sets the current level.
func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	atomic.StoreInt32(&l.level, int32(level))
}

// GetLevel returns the current level.
func (l *Logger) GetLevel() Level {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return Level(l.level)
}

// SetTimeFormat sets the time format.
func (l *Logger) SetTimeFormat(format string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timeFormat = format
}

// SetTimeFunction sets the time function.
func (l *Logger) SetTimeFunction(f TimeFunction) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timeFunc = f
}

// SetFormatter sets the format log.
func (l *Logger) SetFormatter(f Formatter) {
  l.mu.Lock()
  defer l.mu.Unlock()
  l.formatter = f
}

func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()

  if w == nil {
    w = os.Stderr
  }
	l.w = w
  var isDiscard uint32
  if w == io.Discard {
    isDiscard = 1
  }
  atomic.StoreUint32(&l.isDiscard, isDiscard)
}

func (l *Logger) Debug(msg interface{}, keyvals ...interface{}) {
	l.log(DebugLevel, msg, keyvals...)
}

func (l *Logger) Info(msg interface{}, keyvals ...interface{}) {
	l.log(InfoLevel, msg, keyvals...)
}

func (l *Logger) Warn(msg interface{}, keyvals ...interface{}) {
	l.log(WarnLevel, msg, keyvals...)
}

func (l *Logger) Error(msg interface{}, keyvals ...interface{}) {
	l.log(ErrorLevel, msg, keyvals...)
}

func (l *Logger) Fatal(msg interface{}, keyvals ...interface{}) {
	l.log(FatalLevel, msg, keyvals...)
  os.Exit(1)
}
