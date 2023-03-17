package log

import (
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

type Options struct {
  // TimeFunction is the time function for the logger. The default is time.Now
  TimeFunction TimeFunction
  // TimeFormat is the time format for the Logger. The defautl is "2006/01/02 15:04:05"
  TimeFormat string
  // Level is the level for the logger.
  Level Level
  // ReporTimestamp is whether the logger should report the timestamp.
  ReporTimestamp bool
}
