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
)

func levelStyle(level Level) lipgloss.Style {
	switch level {
	case DebugLevel:
		return DebugLevelStyle
	default:
		return lipgloss.NewStyle()
	}
}
