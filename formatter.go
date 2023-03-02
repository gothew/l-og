package log

type Formatter uint8

const (
	// LogftmFormatter is a format that formats log messages as logfmt.
	LogftmFormatter = iota
)

const (
	msgKey = "msg"
)

var (
  // MessageKey is the key for message
  MessageKey = "msg"
  // LevelKey is the key for level
  LevelKey = "lvl"
)
