package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var defaultLogger = NewWithOptions(os.Stderr, Options{ReportTimestamp: true})

func New(w io.Writer) *Logger {
	return NewWithOptions(w, Options{})
}

func NewWithOptions(w io.Writer, o Options) *Logger {
	l := &Logger{
		b:     bytes.Buffer{},
		mu:    &sync.RWMutex{},
    helpers: &sync.Map{},
		level: int32(o.Level),
		reportTimestamp: o.ReportTimestamp,
    reportCaller: o.ReportCaller,
    callerFormatter: o.CallerFormatter,
	}

	l.SetOutput(w)
	l.SetLevel(Level(l.level))

	if l.timeFunc == nil {
		l.timeFunc = time.Now
	}

	if l.timeFormat == "" {
		l.timeFormat = DefaultTimeFormat
	}

  if l.callerFormatter == nil {
    l.callerFormatter = ShowCallerFormatter
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
	os.Exit(1)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.log(DebugLevel, fmt.Sprintf(format, args...))
}

func Infof(format string, args ...interface{}) {
	defaultLogger.log(InfoLevel, fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.log(WarnLevel, fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.log(ErrorLevel, fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.log(FatalLevel, fmt.Sprintf(format, args...))
	os.Exit(1)
}
