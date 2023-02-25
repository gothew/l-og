package log

type Formatter uint8

const (
  // LogftmFormatter is a format that formats log messages as logfmt.
  LogftmFormatter = iota
)

const (
  msgKey = "msg"
)
