package charmthemes_test

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/table"
	"charm.land/glamour/v2"
	"charm.land/lipgloss/v2"

	themes "github.com/junhinhow/charm-themes"
)

// Exemplo de uso com lipgloss.
func ExampleTheme_Lipgloss() {
	theme := themes.Gruvbox

	// Chame uma única vez e reutilize o resultado.
	styles := theme.Lipgloss()

	fmt.Println(styles.Title.Render("Hello"))
	fmt.Println(styles.Error.Render("Something went wrong"))
	fmt.Println(styles.Border.Render("Boxed content"))

	// Ou use as cores diretamente em estilos customizados.
	custom := lipgloss.NewStyle().
		Foreground(theme.Primary).
		Background(theme.Background)

	fmt.Println(custom.Render("Custom"))
}

// Exemplo de lookup de tema por nome.
func ExampleGet() {
	theme, ok := themes.Get("nord")
	if !ok {
		panic("tema não encontrado")
	}
	fmt.Println(theme.Name)
}

// Exemplo de renderização de markdown com o glamour.
func ExampleTheme_Glamour() {
	theme := themes.Nord

	renderer, err := theme.GlamourRenderer(glamour.WithWordWrap(60))
	if err != nil {
		panic(err)
	}

	saida, err := renderer.Render("# Título\n\nTexto de exemplo.\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(saida)
}

// Exemplo de estilos dos componentes do bubbles.
func ExampleTheme_Bubbles() {
	theme := themes.Catppuccin

	styles := theme.Bubbles()
	fmt.Println(styles.Table.Header.Render("ID"))
	fmt.Println(styles.List.Title.Render("Itens"))

	l := list.New(nil, list.NewDefaultDelegate(), 80, 24)
	l.Styles = styles.List
	fmt.Println(l.View())

	t := table.New(table.WithColumns([]table.Column{{Title: "ID", Width: 10}}))
	t.SetStyles(styles.Table)
}
