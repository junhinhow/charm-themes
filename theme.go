// Package charmthemes fornece um sistema de temas unificado para as
// bibliotecas charmbracelet.
//
// Um Theme define a paleta de cores uma única vez e deriva os estilos de cada
// biblioteca a partir dela:
//
//   - Theme.Lipgloss() devolve os estilos do lipgloss.
//   - Theme.Glamour() devolve a configuração de estilos do glamour.
//   - Theme.Bubbles() devolve os estilos dos componentes do bubbles.
//
// Em vez de configurar as cores separadamente em cada biblioteca, defina um
// Theme uma vez e derive todos os estilos a partir dele.
package charmthemes

import (
	"fmt"
	"image/color"
	"strings"
)

// Theme define uma paleta de cores completa, aplicável a todas as
// bibliotecas charmbracelet. Cada tema oferece papéis semânticos de cor que
// se mapeiam para diferentes contextos de interface.
type Theme struct {
	// Name nome do tema.
	Name string

	// Cores base.
	Background color.Color
	Foreground color.Color

	// Cores de destaque para títulos, links e enfases.
	Primary   color.Color
	Secondary color.Color
	Accent    color.Color

	// Cores semânticas.
	Success color.Color
	Warning color.Color
	Error   color.Color
	Info    color.Color
	Muted   color.Color

	// Syntax cores de realce de sintaxe (para blocos de código).
	Syntax SyntaxColors
}

// papelColor associa o nome de um papel semântico à cor correspondente.
type papelColor struct {
	nome string
	cor  color.Color
}

// papeis devolve os papéis de cor do tema em ordem determinística.
func (t Theme) papeis() []papelColor {
	return []papelColor{
		{"Background", t.Background},
		{"Foreground", t.Foreground},
		{"Primary", t.Primary},
		{"Secondary", t.Secondary},
		{"Accent", t.Accent},
		{"Success", t.Success},
		{"Warning", t.Warning},
		{"Error", t.Error},
		{"Info", t.Info},
		{"Muted", t.Muted},
	}
}

// Validate verifica se o tema está completo, isto é, se todas as cores
// declaradas e o nome estão definidos.
//
// Devolve um erro listando os papéis ausentes. Um tema montado à mão sem
// alguma cor deve ser validado antes do uso: sem Validate, a renderização
// simplesmente não teria cor nenhuma e nenhum erro seria emitido.
func (t Theme) Validate() error {
	var ausentes []string

	if strings.TrimSpace(t.Name) == "" {
		ausentes = append(ausentes, "Name")
	}
	for _, papel := range t.papeis() {
		if papel.cor == nil {
			ausentes = append(ausentes, papel.nome)
		}
	}
	ausentes = append(ausentes, t.Syntax.ausentes()...)

	if len(ausentes) > 0 {
		return fmt.Errorf("tema %q incompleto: cores ausentes: %s", t.Name, strings.Join(ausentes, ", "))
	}
	return nil
}

// SyntaxColors define as cores de realce de sintaxe de código.
type SyntaxColors struct {
	Keyword  color.Color
	String   color.Color
	Number   color.Color
	Comment  color.Color
	Function color.Color
	Type     color.Color
	Operator color.Color
	Constant color.Color
}

// Validate verifica se todas as cores de realce de sintaxe estão definidas.
func (s SyntaxColors) Validate() error {
	ausentes := s.ausentes()
	if len(ausentes) > 0 {
		return fmt.Errorf("cores de sintaxe ausentes: %s", strings.Join(ausentes, ", "))
	}
	return nil
}

// ausentes devolve os nomes das cores de realce de sintaxe não definidas, em
// ordem determinística.
func (s SyntaxColors) ausentes() []string {
	cores := []papelColor{
		{"Syntax.Keyword", s.Keyword},
		{"Syntax.String", s.String},
		{"Syntax.Number", s.Number},
		{"Syntax.Comment", s.Comment},
		{"Syntax.Function", s.Function},
		{"Syntax.Type", s.Type},
		{"Syntax.Operator", s.Operator},
		{"Syntax.Constant", s.Constant},
	}

	var ausentes []string
	for _, c := range cores {
		if c.cor == nil {
			ausentes = append(ausentes, c.nome)
		}
	}
	return ausentes
}
