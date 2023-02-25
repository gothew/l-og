package log

import "fmt"

func (l *logger) textFormatter(keyvals ...interface{}) {
	for i := 0; i < len(keyvals); i += 2 {
		switch keyvals[i] {
		case msgKey:
			if msg := keyvals[i+1]; msg != nil {
				m := fmt.Sprint(msg)

				l.b.WriteString(m)
			}
		}
	}
}
