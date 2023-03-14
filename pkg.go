package log

import (
	"io"
)

var defaultLogger = New(WithTimestamp()).(*logger)

func SetOutput(w io.Writer) {
	defaultLogger.SetOutput(w)
}

func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
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
