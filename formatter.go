package log

type Formatter uint8

const (
	// LogftmFormatter is a format that formats log messages as logfmt.
	LogftmFormatter = iota
  //
  JSONFormatter
)

var (
  // TimestampKey is the key for the timestamp.
  TimestampKey = "ts"
  // MessageKey is the key for message
  MessageKey = "msg"
  // LevelKey is the key for level
  LevelKey = "lvl"
  // CallerKey is the key for caller
  CallerKey = "caller"
)
