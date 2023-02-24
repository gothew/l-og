package log

import (
	"io"
	"time"
)

// DefaultTimeFormat is the default time format.
const DefaultTimeFormat = "2006/01/02 15:04:05"

// TimeFunction is a function that returns a time.Time
type TimeFunction = func() time.Time


// NowUTC is a convenient function that returns the current time in UTC
// timezone.
// e.g
//
//    log.SetTimeFunction(log.NowUTC)
func NowUTC() time.Time {
  return time.Now().UTC()
}

// Logger is a interface for logging.
type Logger interface {
  // SetOutput sets the output destination.
  SetOutput(w io.Writer)

  // Debug logs a debug message.
  Debug(msg interface{}, keyval ...interface{})
}
