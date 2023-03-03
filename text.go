package log

import (
	"fmt"
	"strings"
)

func (l *logger) textFormatter(keyvals ...interface{}) {
	for i := 0; i < len(keyvals); i += 2 {
		switch keyvals[i] {
		case LevelKey:
			if level, ok := keyvals[i+1].(Level); ok {
				lvl := strings.ToUpper(level.String())

				if !l.notStyles {
					lvl = levelStyle(level).String()
				}

				l.b.WriteString(lvl)
				l.b.WriteByte(' ')
			}
		case MessageKey:
			if msg := keyvals[i+1]; msg != nil {
				m := fmt.Sprint(msg)

				l.b.WriteString(m)
			}
		}
	}
	// Add a newline in log message.
	l.b.WriteByte('\n')
}
