package log

import (
	"io"
)

var defaultLogger = New().(*logger)

func SetOutput(w io.Writer) {
  defaultLogger.SetOutput(w)
}

func Debug(msg interface{}, keyvals ...interface{}) {
  defaultLogger.log(msg, keyvals...)
}
