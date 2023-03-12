package log

func WithTimestamp() LoggerOptions {
  return func(l *logger) {
    l.timestamp = true
  }
}
