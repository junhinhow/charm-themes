package charmthemes

import (
	"sort"
	"strings"

	"charm.land/lipgloss/v2"
)

// Temas embutidos no pacote.
var (
	Gruvbox    = gruvboxTheme()
	Dracula    = draculaTheme()
	Nord       = nordTheme()
	Catppuccin = catppuccinTheme()
	TokyoNight = tokyoNightTheme()
)

// builtinThemes indexa os temas embutidos pelo nome.
//
// Os ponteiros apontam para as variáveis globais; Themes e Get devolvem cópias
// dos valores, de modo que customizar um tema devolvido não afete o processo
// inteiro.
var builtinThemes = map[string]*Theme{
	"gruvbox":     &Gruvbox,
	"dracula":     &Dracula,
	"nord":        &Nord,
	"catppuccin":  &Catppuccin,
	"tokyo-night": &TokyoNight,
}

// Themes devolve todos os temas embutidos, indexados pelo nome.
//
// Os temas devolvidos são cópias: customizar o valor devolvido não altera os
// temas globais do pacote.
func Themes() map[string]*Theme {
	temas := make(map[string]*Theme, len(builtinThemes))
	for nome, tema := range builtinThemes {
		copia := *tema
		temas[nome] = &copia
	}
	return temas
}

// Get devolve o tema embutido com o nome informado.
//
// A busca ignora maiúsculas, minúsculas e espaços nas pontas. O segundo
// retorno indica se o tema existe. O tema devolvido é uma cópia e pode ser
// modificado com segurança.
func Get(name string) (*Theme, bool) {
	chave := strings.ToLower(strings.TrimSpace(name))
	tema, ok := builtinThemes[chave]
	if !ok {
		return nil, false
	}
	copia := *tema
	return &copia, true
}

// ThemeNames devolve os nomes dos temas embutidos em ordem alfabética.
//
// A ordenação é determinística, ao contrário da iteração sobre o mapa
// devolvido por Themes.
func ThemeNames() []string {
	nomes := make([]string, 0, len(builtinThemes))
	for nome := range builtinThemes {
		nomes = append(nomes, nome)
	}
	sort.Strings(nomes)
	return nomes
}

func gruvboxTheme() Theme {
	return Theme{
		Name:       "gruvbox",
		Background: lipgloss.Color("#282828"),
		Foreground: lipgloss.Color("#ebdbb2"),
		Primary:    lipgloss.Color("#fabd2f"),
		Secondary:  lipgloss.Color("#b8bb26"),
		Accent:     lipgloss.Color("#83a598"),
		Success:    lipgloss.Color("#b8bb26"),
		Warning:    lipgloss.Color("#fe8019"),
		Error:      lipgloss.Color("#fb4934"),
		Info:       lipgloss.Color("#83a598"),
		Muted:      lipgloss.Color("#928374"),
		Syntax: SyntaxColors{
			Keyword:  lipgloss.Color("#fb4934"),
			String:   lipgloss.Color("#b8bb26"),
			Number:   lipgloss.Color("#d3869b"),
			Comment:  lipgloss.Color("#928374"),
			Function: lipgloss.Color("#b8bb26"),
			Type:     lipgloss.Color("#fabd2f"),
			Operator: lipgloss.Color("#fe8019"),
			Constant: lipgloss.Color("#d3869b"),
		},
	}
}

func draculaTheme() Theme {
	return Theme{
		Name:       "dracula",
		Background: lipgloss.Color("#282a36"),
		Foreground: lipgloss.Color("#f8f8f2"),
		Primary:    lipgloss.Color("#bd93f9"),
		Secondary:  lipgloss.Color("#50fa7b"),
		Accent:     lipgloss.Color("#ff79c6"),
		Success:    lipgloss.Color("#50fa7b"),
		Warning:    lipgloss.Color("#ffb86c"),
		Error:      lipgloss.Color("#ff5555"),
		Info:       lipgloss.Color("#8be9fd"),
		Muted:      lipgloss.Color("#6272a4"),
		Syntax: SyntaxColors{
			Keyword:  lipgloss.Color("#ff79c6"),
			String:   lipgloss.Color("#f1fa8c"),
			Number:   lipgloss.Color("#bd93f9"),
			Comment:  lipgloss.Color("#6272a4"),
			Function: lipgloss.Color("#50fa7b"),
			Type:     lipgloss.Color("#8be9fd"),
			Operator: lipgloss.Color("#ff79c6"),
			Constant: lipgloss.Color("#bd93f9"),
		},
	}
}

func nordTheme() Theme {
	return Theme{
		Name:       "nord",
		Background: lipgloss.Color("#2e3440"),
		Foreground: lipgloss.Color("#d8dee9"),
		Primary:    lipgloss.Color("#88c0d0"),
		Secondary:  lipgloss.Color("#a3be8c"),
		Accent:     lipgloss.Color("#b48ead"),
		Success:    lipgloss.Color("#a3be8c"),
		Warning:    lipgloss.Color("#ebcb8b"),
		Error:      lipgloss.Color("#bf616a"),
		Info:       lipgloss.Color("#5e81ac"),
		Muted:      lipgloss.Color("#4c566a"),
		Syntax: SyntaxColors{
			Keyword:  lipgloss.Color("#81a1c1"),
			String:   lipgloss.Color("#a3be8c"),
			Number:   lipgloss.Color("#b48ead"),
			Comment:  lipgloss.Color("#616e88"),
			Function: lipgloss.Color("#88c0d0"),
			Type:     lipgloss.Color("#8fbcbb"),
			Operator: lipgloss.Color("#81a1c1"),
			Constant: lipgloss.Color("#b48ead"),
		},
	}
}

func catppuccinTheme() Theme {
	return Theme{
		Name:       "catppuccin",
		Background: lipgloss.Color("#1e1e2e"),
		Foreground: lipgloss.Color("#cdd6f4"),
		Primary:    lipgloss.Color("#cba6f7"),
		Secondary:  lipgloss.Color("#a6e3a1"),
		Accent:     lipgloss.Color("#f5c2e7"),
		Success:    lipgloss.Color("#a6e3a1"),
		Warning:    lipgloss.Color("#f9e2af"),
		Error:      lipgloss.Color("#f38ba8"),
		Info:       lipgloss.Color("#89b4fa"),
		Muted:      lipgloss.Color("#585b70"),
		Syntax: SyntaxColors{
			Keyword:  lipgloss.Color("#cba6f7"),
			String:   lipgloss.Color("#a6e3a1"),
			Number:   lipgloss.Color("#fab387"),
			Comment:  lipgloss.Color("#585b70"),
			Function: lipgloss.Color("#89b4fa"),
			Type:     lipgloss.Color("#f9e2af"),
			Operator: lipgloss.Color("#89dceb"),
			Constant: lipgloss.Color("#fab387"),
		},
	}
}

func tokyoNightTheme() Theme {
	return Theme{
		Name:       "tokyo-night",
		Background: lipgloss.Color("#1a1b26"),
		Foreground: lipgloss.Color("#a9b1d6"),
		Primary:    lipgloss.Color("#7aa2f7"),
		Secondary:  lipgloss.Color("#9ece6a"),
		Accent:     lipgloss.Color("#bb9af7"),
		Success:    lipgloss.Color("#9ece6a"),
		Warning:    lipgloss.Color("#e0af68"),
		Error:      lipgloss.Color("#f7768e"),
		Info:       lipgloss.Color("#7dcfff"),
		Muted:      lipgloss.Color("#565f89"),
		Syntax: SyntaxColors{
			Keyword:  lipgloss.Color("#bb9af7"),
			String:   lipgloss.Color("#9ece6a"),
			Number:   lipgloss.Color("#ff9e64"),
			Comment:  lipgloss.Color("#565f89"),
			Function: lipgloss.Color("#7aa2f7"),
			Type:     lipgloss.Color("#2ac3de"),
			Operator: lipgloss.Color("#89ddff"),
			Constant: lipgloss.Color("#ff9e64"),
		},
	}
}
