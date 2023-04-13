package log

import (
	"fmt"
	"time"
)

// DefaultTimeFormat is the default time format.
const DefaultTimeFormat = "2006/01/02 15:04:05"

// TimeFunction is a function that returns a time.Time.
type TimeFunction = func() time.Time

// CallerFormatter is the caller formatter.
type CallerFormatter func(string, int, string) string

// NowUTC is a convenient function that returns the current time in UTC
// timezone.
// e.g
//
//	log.SetTimeFunction(log.NowUTC)
func NowUTC() time.Time {
	return time.Now().UTC()
}

// ShowCallerFormatter is a caller formatter that returns the last 2 levels of
// the path and line number.
func ShowCallerFormatter(file string, line int, funcName string) string {
  return fmt.Sprintf("%s:%d", trimCallerPath(file, 2), line)
}

type Options struct {
	// TimeFunction is the time function for the logger. The default is time.Now
	TimeFunction TimeFunction
	// TimeFormat is the time format for the Logger. The defautl is "2006/01/02 15:04:05"
	TimeFormat string
	// Level is the level for the logger.
	Level Level
	// ReporTimestamp is whether the logger should report the timestamp.
	ReportTimestamp bool
	// ReportCaller is whether the logger should report the caller location
	ReportCaller bool
	// CallerFormatter is the caller format for the logger.
	CallerFormatter CallerFormatter
}
