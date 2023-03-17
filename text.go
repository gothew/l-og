package log

import (
	"fmt"
	"strings"
	"time"
)

func (l *Logger) textFormatter(keyvals ...interface{}) {
	for i := 0; i < len(keyvals); i += 2 {
		switch keyvals[i] {
    case TimestampKey:
      if t, ok := keyvals[i+1].(time.Time); ok {
        ts := t.Format(l.timeFormat)
        
        if !l.notStyles {
          ts = TimestampStyle.Render(ts)
        }

        l.b.WriteString(ts)
        l.b.WriteByte(' ')
      }
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
