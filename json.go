package log

import "encoding/json"

func (l *Logger) jsonFormatter(keyvals ...interface{}) {
	m := make(map[string]interface{}, len(keyvals)%2)
	for i := 0; i < len(keyvals); i += 2 {
		switch keyvals[i] {
		case LevelKey:
			if level, ok := keyvals[i+1].(Level); ok {
				m[LevelKey] = level.String()
			}
		}
	}

	e := json.NewEncoder(&l.b)
	e.SetEscapeHTML(false)
	_ = e.Encode(m)
}
