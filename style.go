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
)

func levelStyle(level Level) lipgloss.Style {
	switch level {
	case DebugLevel:
		return DebugLevelStyle
	case InfoLevel:
		return InfoLevelStyle
  case WarnLevel:
    return WarnLevelStyle
	default:
		return lipgloss.NewStyle()
	}
}
