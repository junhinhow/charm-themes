package charmthemes

import "charm.land/lipgloss/v2"

// LipglossStyles reúne os estilos de lipgloss derivados de um Theme.
type LipglossStyles struct {
	// Estilos de texto.
	Base   lipgloss.Style
	Bold   lipgloss.Style
	Italic lipgloss.Style
	Faint  lipgloss.Style

	// Estilos semânticos.
	Success lipgloss.Style
	Warning lipgloss.Style
	Error   lipgloss.Style
	Info    lipgloss.Style

	// Estilos de componentes de interface.
	Title     lipgloss.Style
	Subtitle  lipgloss.Style
	Highlight lipgloss.Style
	Border    lipgloss.Style
	Code      lipgloss.Style
}

// Lipgloss devolve um conjunto de estilos de lipgloss derivados do tema.
//
// Cada chamada constrói 13 estilos de lipgloss. Chame o método uma única vez
// e reutilize o resultado, em vez de chamá-lo dentro do laço de renderização:
//
//	styles := theme.Lipgloss() // uma vez, na inicialização
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
