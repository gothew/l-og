package log

import "github.com/charmbracelet/lipgloss"

var (
	TimestampStyle = lipgloss.NewStyle()

	DebugLevelStyle = lipgloss.NewStyle().
			SetString("DEBUG").
			Bold(true).
			MaxWidth(4).
			Foreground(lipgloss.AdaptiveColor{
			Light: "63",
			Dark:  "63",
		})

	InfoLevelStyle = lipgloss.NewStyle().
			SetString("INFO").
			Bold(true).
			MaxWidth(4).
			Foreground(lipgloss.AdaptiveColor{
			Light: "39",
			Dark:  "86",
		})

	WarnLevelStyle = lipgloss.NewStyle().
			SetString("WARN").
			Bold(true).
			MaxWidth(4).
			Foreground(lipgloss.AdaptiveColor{
			Light: "208",
			Dark:  "192",
		})

	ErrorLevelStyle = lipgloss.NewStyle().
			SetString("ERROR").
			Bold(true).
			MaxWidth(4).
			Foreground(lipgloss.AdaptiveColor{
			Light: "203",
			Dark:  "204",
		})

  FatalLevelStyle = lipgloss.NewStyle().
			SetString("FATAL").
			Bold(true).
			MaxWidth(4).
			Foreground(lipgloss.AdaptiveColor{
			Light: "133",
			Dark:  "134",
		})
)

func levelStyle(level Level) lipgloss.Style {
	switch level {
	case DebugLevel:
		return DebugLevelStyle
	case InfoLevel:
		return InfoLevelStyle
  case WarnLevel:
    return WarnLevelStyle
  case ErrorLevel:
    return ErrorLevelStyle
  case FatalLevel:
    return FatalLevelStyle
	default:
		return lipgloss.NewStyle()
	}
}
