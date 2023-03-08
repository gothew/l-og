package log

import "github.com/charmbracelet/lipgloss"

var (
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
)

func levelStyle(level Level) lipgloss.Style {
	switch level {
	case DebugLevel:
		return DebugLevelStyle
  case InfoLevel:
    return InfoLevelStyle
	default:
		return lipgloss.NewStyle()
	}
}
