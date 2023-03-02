package log

import "strings"

// Level is a logging level.
type Level int32

const (
	// DebugLevel is the debug level.
	DebugLevel Level = iota - 1
	// InfoLevel is the info level.
	InfoLevel

  // noLevel is used with log.Print.
  noLevel
)

// String returns the string representation of the Level.
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	default:
		return ""
	}
}

// ParseLevel converts level in string to Level type.
func ParseLevel(level string) Level {
	switch strings.ToLower(level) {
	case DebugLevel.String():
		return DebugLevel
	case InfoLevel.String():
		return InfoLevel
	default:
		return InfoLevel
	}
}
