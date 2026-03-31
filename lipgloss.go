package charmthemes

import "charm.land/lipgloss/v2"

// LipglossStyles provides pre-built lipgloss styles derived from a Theme.
type LipglossStyles struct {
	// Base text styles
	Base    lipgloss.Style
	Bold    lipgloss.Style
	Italic  lipgloss.Style
	Faint   lipgloss.Style

	// Semantic styles
	Success lipgloss.Style
	Warning lipgloss.Style
	Error   lipgloss.Style
	Info    lipgloss.Style

	// UI component styles
	Title     lipgloss.Style
	Subtitle  lipgloss.Style
	Highlight lipgloss.Style
	Border    lipgloss.Style
	Code      lipgloss.Style
}

// Lipgloss returns a set of lipgloss styles derived from the theme.
func (t Theme) Lipgloss() LipglossStyles {
	return LipglossStyles{
		Base:   lipgloss.NewStyle().Foreground(t.Foreground),
		Bold:   lipgloss.NewStyle().Foreground(t.Foreground).Bold(true),
		Italic: lipgloss.NewStyle().Foreground(t.Foreground).Italic(true),
		Faint:  lipgloss.NewStyle().Foreground(t.Muted),

		Success: lipgloss.NewStyle().Foreground(t.Success),
		Warning: lipgloss.NewStyle().Foreground(t.Warning),
		Error:   lipgloss.NewStyle().Foreground(t.Error),
		Info:    lipgloss.NewStyle().Foreground(t.Info),

		Title: lipgloss.NewStyle().
			Foreground(t.Primary).
			Bold(true),

		Subtitle: lipgloss.NewStyle().
			Foreground(t.Secondary),

		Highlight: lipgloss.NewStyle().
			Foreground(t.Accent).
			Bold(true),

		Border: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(t.Muted).
			Padding(1, 2),

		Code: lipgloss.NewStyle().
			Foreground(t.Syntax.String).
			Background(t.Background),
	}
}
