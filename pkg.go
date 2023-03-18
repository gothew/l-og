package log

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

var defaultLogger = NewWithOptions(os.Stderr, Options{ReporTimestamp: true}) 

func New(w io.Writer) *Logger {
  return NewWithOptions(w, Options{})
}

func NewWithOptions(w io.Writer, o Options) *Logger {
  l := &Logger{
    b: bytes.Buffer{},
    mu: &sync.RWMutex{},
    level: int32(o.Level),

    reportTimestamp: o.ReporTimestamp,
  }

  l.SetOutput(w)
  l.SetLevel(Level(l.level))

  if l.timeFunc == nil {
    l.timeFunc = time.Now
  }

  if l.timeFormat == "" {
    l.timeFormat = DefaultTimeFormat
  }

  return l
}

func SetOutput(w io.Writer) {
	defaultLogger.SetOutput(w)
}

func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}

func SetFormatter(f Formatter) {
  defaultLogger.SetFormatter(f)
}

func GetLevel() Level {
	return defaultLogger.GetLevel()
}

func Debug(msg interface{}, keyvals ...interface{}) {
	defaultLogger.log(DebugLevel, msg, keyvals...)
}

func Info(msg interface{}, keyvals ...interface{}) {
	defaultLogger.log(InfoLevel, msg, keyvals...)
}

func Warn(msg interface{}, keyvals ...interface{}) {
	defaultLogger.log(WarnLevel, msg, keyvals...)
}

func Error(msg interface{}, keyvals ...interface{}) {
	defaultLogger.log(ErrorLevel, msg, keyvals...)
}

func Fatal(msg interface{}, keyvals ...interface{}) {
	defaultLogger.log(FatalLevel, msg, keyvals...)
}
