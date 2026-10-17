package charmthemes

import (
	"time"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// BubblesStyles reúne os estilos dos componentes do bubbles derivados de um
// Theme.
//
// Os estilos podem ser aplicados com os métodos SetStyles dos componentes:
//
//	l := list.New(items, list.NewDefaultDelegate(), 80, 24)
//	l.SetStyles(theme.Bubbles().List)
type BubblesStyles struct {
	// List estilos do componente list.
	List list.Styles
	// Table estilos do componente table.
	Table table.Styles
	// TextInput estilos do componente textinput.
	TextInput textinput.Styles
	// TextArea estilos do componente textarea.
	TextArea textarea.Styles
	// Help estilos do componente help.
	Help help.Styles
}

// Bubbles devolve os estilos dos componentes do bubbles derivados do tema.
//
// Assim como Lipgloss, cada chamada constrói os estilos do zero. Chame o
// método uma única vez e reutilize o resultado.
func (t Theme) Bubbles() BubblesStyles {
	return BubblesStyles{
		List:      t.listStyles(),
		Table:     t.tableStyles(),
		TextInput: t.textInputStyles(),
		TextArea:  t.textAreaStyles(),
		Help:      t.helpStyles(),
	}
}

// listStyles devolve os estilos do componente list.
func (t Theme) listStyles() list.Styles {
	return list.Styles{
		TitleBar: lipgloss.NewStyle().Padding(0, 0, 1, 2),
		Title: lipgloss.NewStyle().
			Background(t.Primary).
			Foreground(t.Background).
			Bold(true).
			Padding(0, 1),
		Spinner: lipgloss.NewStyle().Foreground(t.Muted),
		Filter:  t.textInputStyles(),
		DefaultFilterCharacterMatch: lipgloss.NewStyle().
			Foreground(t.Accent).
			Underline(true),
		StatusBar:             lipgloss.NewStyle().Foreground(t.Muted).Padding(0, 0, 1, 2),
		StatusEmpty:           lipgloss.NewStyle().Foreground(t.Muted),
		StatusBarActiveFilter: lipgloss.NewStyle().Foreground(t.Foreground),
		StatusBarFilterCount:  lipgloss.NewStyle().Foreground(t.Muted),
		NoItems:               lipgloss.NewStyle().Foreground(t.Muted),
		PaginationStyle:       lipgloss.NewStyle().PaddingLeft(2),
		HelpStyle:             lipgloss.NewStyle().Padding(1, 0, 0, 2),
		ActivePaginationDot: lipgloss.NewStyle().
			Foreground(t.Accent).
			SetString("•"),
		InactivePaginationDot: lipgloss.NewStyle().
			Foreground(t.Muted).
			SetString("•"),
		ArabicPagination: lipgloss.NewStyle().Foreground(t.Muted),
		DividerDot: lipgloss.NewStyle().
			Foreground(t.Muted).
			SetString(" • "),
	}
}

// tableStyles devolve os estilos do componente table.
func (t Theme) tableStyles() table.Styles {
	return table.Styles{
		Header: lipgloss.NewStyle().
			Foreground(t.Primary).
			Bold(true).
			Padding(0, 1),
		Cell: lipgloss.NewStyle().
			Foreground(t.Foreground).
			Padding(0, 1),
		Selected: lipgloss.NewStyle().
			Foreground(t.Background).
			Background(t.Accent).
			Bold(true),
	}
}

// textInputStyles devolve os estilos do componente textinput.
func (t Theme) textInputStyles() textinput.Styles {
	return textinput.Styles{
		Focused: textinput.StyleState{
			Text:        lipgloss.NewStyle().Foreground(t.Foreground),
			Placeholder: lipgloss.NewStyle().Foreground(t.Muted),
			Suggestion:  lipgloss.NewStyle().Foreground(t.Muted),
			Prompt:      lipgloss.NewStyle().Foreground(t.Primary),
		},
		Blurred: textinput.StyleState{
			Text:        lipgloss.NewStyle().Foreground(t.Muted),
			Placeholder: lipgloss.NewStyle().Foreground(t.Muted),
			Suggestion:  lipgloss.NewStyle().Foreground(t.Muted),
			Prompt:      lipgloss.NewStyle().Foreground(t.Muted),
		},
		Cursor: textinput.CursorStyle{
			Color:      t.Foreground,
			Shape:      tea.CursorBlock,
			Blink:      true,
			BlinkSpeed: 500 * time.Millisecond,
		},
	}
}

// textAreaStyles devolve os estilos do componente textarea.
func (t Theme) textAreaStyles() textarea.Styles {
	return textarea.Styles{
		Focused: textarea.StyleState{
			Base:             lipgloss.NewStyle().Foreground(t.Foreground),
			Text:             lipgloss.NewStyle().Foreground(t.Foreground),
			LineNumber:       lipgloss.NewStyle().Foreground(t.Muted),
			CursorLineNumber: lipgloss.NewStyle().Foreground(t.Primary),
			CursorLine:       lipgloss.NewStyle().Foreground(t.Foreground),
			EndOfBuffer:      lipgloss.NewStyle().Foreground(t.Muted),
			Placeholder:      lipgloss.NewStyle().Foreground(t.Muted),
			Prompt:           lipgloss.NewStyle().Foreground(t.Primary),
		},
		Blurred: textarea.StyleState{
			Base:             lipgloss.NewStyle().Foreground(t.Muted),
			Text:             lipgloss.NewStyle().Foreground(t.Muted),
			LineNumber:       lipgloss.NewStyle().Foreground(t.Muted),
			CursorLineNumber: lipgloss.NewStyle().Foreground(t.Muted),
			CursorLine:       lipgloss.NewStyle().Foreground(t.Muted),
			EndOfBuffer:      lipgloss.NewStyle().Foreground(t.Muted),
			Placeholder:      lipgloss.NewStyle().Foreground(t.Muted),
			Prompt:           lipgloss.NewStyle().Foreground(t.Muted),
		},
		Cursor: textarea.CursorStyle{
			Color:      t.Foreground,
			Shape:      tea.CursorBlock,
			Blink:      true,
			BlinkSpeed: 500 * time.Millisecond,
		},
	}
}

// helpStyles devolve os estilos do componente help.
func (t Theme) helpStyles() help.Styles {
	return help.Styles{
		Ellipsis:       lipgloss.NewStyle().Foreground(t.Muted),
		ShortKey:       lipgloss.NewStyle().Foreground(t.Accent),
		ShortDesc:      lipgloss.NewStyle().Foreground(t.Muted),
		ShortSeparator: lipgloss.NewStyle().Foreground(t.Muted),
		FullKey:        lipgloss.NewStyle().Foreground(t.Accent),
		FullDesc:       lipgloss.NewStyle().Foreground(t.Muted),
		FullSeparator:  lipgloss.NewStyle().Foreground(t.Muted),
	}
}
