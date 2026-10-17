package charmthemes

import (
	"image/color"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

// TestThemesNaoCorrompeGlobais garante que Themes() devolve cópias e não
// ponteiros para as variáveis globais do pacote.
func TestThemesNaoCorrompeGlobais(t *testing.T) {
	original := Nord.Primary

	temas := Themes()
	nord, ok := temas["nord"]
	if !ok {
		t.Fatal("Themes() deveria conter o tema nord")
	}

	nord.Primary = lipgloss.Color("#ff0000")

	if Nord.Primary != original {
		t.Errorf("customizar o tema devolvido corrompeu o global: esperado %v, obteve %v", original, Nord.Primary)
	}
	if temas["nord"].Primary == Nord.Primary {
		t.Error("Themes() deveria devolver uma cópia, não o ponteiro para o global")
	}

	// O global deve continuar intacto para quem usa diretamente.
	if got := colorToHex(Nord.Primary); got != colorToHex(original) {
		t.Errorf("global Nord.Primary alterado: esperado %s, obteve %s", colorToHex(original), got)
	}
}

func TestGetDevolveCopia(t *testing.T) {
	original := Nord.Primary

	tema, ok := Get("nord")
	if !ok {
		t.Fatal("Get(\"nord\") deveria encontrar o tema")
	}
	if tema.Name != "nord" {
		t.Errorf("nome esperado 'nord', obteve '%s'", tema.Name)
	}

	tema.Primary = lipgloss.Color("#00ff00")

	if Nord.Primary != original {
		t.Error("customizar o tema devolvido por Get corrompeu o global")
	}
}

func TestGetIgnoraCaixaEEspacos(t *testing.T) {
	for _, nome := range []string{"nord", "NORD", "Nord", "  nord  "} {
		tema, ok := Get(nome)
		if !ok {
			t.Errorf("Get(%q) deveria encontrar o tema", nome)
			continue
		}
		if tema.Name != "nord" {
			t.Errorf("Get(%q) devolveu o tema '%s'", nome, tema.Name)
		}
	}
}

func TestGetTemaInexistente(t *testing.T) {
	tema, ok := Get("nao-existe")
	if ok {
		t.Error("Get deveria informar que o tema não existe")
	}
	if tema != nil {
		t.Error("Get deveria devolver nil para um tema inexistente")
	}
}

func TestThemeNamesOrdemDeterministica(t *testing.T) {
	esperado := []string{"catppuccin", "dracula", "gruvbox", "nord", "tokyo-night"}

	for range 5 {
		nomes := ThemeNames()
		if len(nomes) != len(esperado) {
			t.Fatalf("esperava %d nomes, obteve %d", len(esperado), len(nomes))
		}
		for i, nome := range esperado {
			if nomes[i] != nome {
				t.Errorf("posição %d: esperado '%s', obteve '%s'", i, nome, nomes[i])
			}
		}
	}
}

func TestThemeNamesCobreTodosOsTemas(t *testing.T) {
	temas := Themes()
	for _, nome := range ThemeNames() {
		if _, ok := temas[nome]; !ok {
			t.Errorf("ThemeNames() informou '%s', que não existe em Themes()", nome)
		}
	}
	if len(ThemeNames()) != len(temas) {
		t.Errorf("ThemeNames() devolveu %d nomes, Themes() devolveu %d temas", len(ThemeNames()), len(temas))
	}
}

// TestThemesSincronizadosComGlobais garante que o mapa e as variáveis globais
// carregam os mesmos temas, com os nomes batendo com as chaves.
func TestThemesSincronizadosComGlobais(t *testing.T) {
	temas := Themes()
	globais := map[string]Theme{
		"gruvbox":     Gruvbox,
		"dracula":     Dracula,
		"nord":        Nord,
		"catppuccin":  Catppuccin,
		"tokyo-night": TokyoNight,
	}

	if len(temas) != len(globais) {
		t.Fatalf("Themes() devolveu %d temas, esperado %d", len(temas), len(globais))
	}

	for nome, global := range globais {
		tema, ok := temas[nome]
		if !ok {
			t.Errorf("Themes() não contém o tema '%s'", nome)
			continue
		}
		if tema.Name != nome {
			t.Errorf("tema '%s' tem Name '%s'", nome, tema.Name)
		}
		if tema.Name != global.Name {
			t.Errorf("tema '%s': nome global '%s' difere do mapa '%s'", nome, global.Name, tema.Name)
		}
		for _, papel := range []struct {
			nome   string
			mapa   color.Color
			global color.Color
		}{
			{"Background", tema.Background, global.Background},
			{"Foreground", tema.Foreground, global.Foreground},
			{"Primary", tema.Primary, global.Primary},
			{"Secondary", tema.Secondary, global.Secondary},
			{"Accent", tema.Accent, global.Accent},
			{"Success", tema.Success, global.Success},
			{"Warning", tema.Warning, global.Warning},
			{"Error", tema.Error, global.Error},
			{"Info", tema.Info, global.Info},
			{"Muted", tema.Muted, global.Muted},
		} {
			if colorToHex(papel.mapa) != colorToHex(papel.global) {
				t.Errorf("tema '%s': cor %s difere (mapa %s, global %s)",
					nome, papel.nome, colorToHex(papel.mapa), colorToHex(papel.global))
			}
		}
	}
}

func TestValidateAceitaTemasEmbutidos(t *testing.T) {
	for nome, tema := range Themes() {
		if err := tema.Validate(); err != nil {
			t.Errorf("tema '%s' deveria ser válido, obteve %v", nome, err)
		}
	}
}

// TestValidateDetectaCoresAusentes garante que um tema montado à mão sem cores
// é rejeitado.
func TestValidateDetectaCoresAusentes(t *testing.T) {
	incompleto := Theme{
		Name:       "incompleto",
		Foreground: lipgloss.Color("#ffffff"),
	}

	err := incompleto.Validate()
	if err == nil {
		t.Fatal("Validate deveria reprovar um tema sem cores")
	}
	for _, esperado := range []string{"Background", "Primary", "Muted"} {
		if !strings.Contains(err.Error(), esperado) {
			t.Errorf("erro deveria mencionar '%s', obteve: %v", esperado, err)
		}
	}
}

func TestValidateDetectaNomeAusente(t *testing.T) {
	tema := Theme{
		Background: lipgloss.Color("#000000"),
		Foreground: lipgloss.Color("#ffffff"),
		Primary:    lipgloss.Color("#ff0000"),
		Secondary:  lipgloss.Color("#00ff00"),
		Accent:     lipgloss.Color("#0000ff"),
		Success:    lipgloss.Color("#00ff00"),
		Warning:    lipgloss.Color("#ffff00"),
		Error:      lipgloss.Color("#ff0000"),
		Info:       lipgloss.Color("#0000ff"),
		Muted:      lipgloss.Color("#888888"),
		Syntax:     Nord.Syntax,
	}

	if err := tema.Validate(); err == nil {
		t.Fatal("Validate deveria reprovar um tema sem nome")
	} else if !strings.Contains(err.Error(), "Name") {
		t.Errorf("erro deveria mencionar 'Name', obteve: %v", err)
	}
}

func TestSyntaxColorsValidate(t *testing.T) {
	if err := Nord.Syntax.Validate(); err != nil {
		t.Errorf("cores de sintaxe do tema nord deveriam ser válidas, obteve %v", err)
	}

	incompleto := SyntaxColors{Keyword: lipgloss.Color("#ff0000")}
	err := incompleto.Validate()
	if err == nil {
		t.Fatal("Validate deveria reprovar cores de sintaxe incompletas")
	}
	if !strings.Contains(err.Error(), "Syntax.String") {
		t.Errorf("erro deveria mencionar 'Syntax.String', obteve: %v", err)
	}
}

// TestLipglossUsaAsCoresDoTema garante que os estilos derivados carregam as
// cores do tema.
func TestLipglossUsaAsCoresDoTema(t *testing.T) {
	styles := Nord.Lipgloss()

	if got := colorToHex(styles.Title.GetForeground()); got != colorToHex(Nord.Primary) {
		t.Errorf("Title: cor esperada %s, obteve %s", colorToHex(Nord.Primary), got)
	}
	if got := colorToHex(styles.Subtitle.GetForeground()); got != colorToHex(Nord.Secondary) {
		t.Errorf("Subtitle: cor esperada %s, obteve %s", colorToHex(Nord.Secondary), got)
	}
	if got := colorToHex(styles.Highlight.GetForeground()); got != colorToHex(Nord.Accent) {
		t.Errorf("Highlight: cor esperada %s, obteve %s", colorToHex(Nord.Accent), got)
	}
	if got := colorToHex(styles.Error.GetForeground()); got != colorToHex(Nord.Error) {
		t.Errorf("Error: cor esperada %s, obteve %s", colorToHex(Nord.Error), got)
	}
	if got := colorToHex(styles.Success.GetForeground()); got != colorToHex(Nord.Success) {
		t.Errorf("Success: cor esperada %s, obteve %s", colorToHex(Nord.Success), got)
	}
	if got := colorToHex(styles.Faint.GetForeground()); got != colorToHex(Nord.Muted) {
		t.Errorf("Faint: cor esperada %s, obteve %s", colorToHex(Nord.Muted), got)
	}
	if got := colorToHex(styles.Code.GetForeground()); got != colorToHex(Nord.Syntax.String) {
		t.Errorf("Code: cor esperada %s, obteve %s", colorToHex(Nord.Syntax.String), got)
	}
}

// TestLipglossEDeterministico garante que duas chamadas produzem o mesmo
// resultado, para que a renderização seja estável.
func TestLipglossEDeterministico(t *testing.T) {
	primeira := Nord.Lipgloss().Title.Render("Título")
	segunda := Nord.Lipgloss().Title.Render("Título")
	if primeira != segunda {
		t.Errorf("Lipgloss() deveria ser determinístico:\n%q\n%q", primeira, segunda)
	}
}

func TestGlamourUsaAsCoresDoTema(t *testing.T) {
	cfg := Nord.Glamour()

	if cfg.Heading.Color == nil {
		t.Fatal("Heading deveria ter cor definida")
	}
	if got := *cfg.Heading.Color; got != colorToHex(Nord.Primary) {
		t.Errorf("Heading: cor esperada %s, obteve %s", colorToHex(Nord.Primary), got)
	}
	if cfg.Text.Color == nil {
		t.Fatal("Text deveria ter cor definida")
	}
	if got := *cfg.Text.Color; got != colorToHex(Nord.Foreground) {
		t.Errorf("Text: cor esperada %s, obteve %s", colorToHex(Nord.Foreground), got)
	}
	if cfg.CodeBlock.Chroma == nil {
		t.Fatal("CodeBlock.Chroma deveria estar configurado")
	}
	if cfg.CodeBlock.Chroma.Keyword.Color == nil {
		t.Fatal("Chroma.Keyword deveria ter cor definida")
	}
	if got := *cfg.CodeBlock.Chroma.Keyword.Color; got != colorToHex(Nord.Syntax.Keyword) {
		t.Errorf("Chroma.Keyword: cor esperada %s, obteve %s", colorToHex(Nord.Syntax.Keyword), got)
	}
}

func TestGlamourRendererRenderizaMarkdown(t *testing.T) {
	renderer, err := Nord.GlamourRenderer()
	if err != nil {
		t.Fatalf("GlamourRenderer deveria criar o renderer, obteve %v", err)
	}

	saida, err := renderer.Render("# Título\n\nTexto de exemplo.\n")
	if err != nil {
		t.Fatalf("renderização do markdown falhou: %v", err)
	}
	if !strings.Contains(saida, "Título") {
		t.Errorf("a saída deveria conter o texto renderizado, obteve:\n%s", saida)
	}
}

func TestBubblesUsaAsCoresDoTema(t *testing.T) {
	styles := Nord.Bubbles()

	if got := colorToHex(styles.Table.Header.GetForeground()); got != colorToHex(Nord.Primary) {
		t.Errorf("Table.Header: cor esperada %s, obteve %s", colorToHex(Nord.Primary), got)
	}
	if got := colorToHex(styles.Table.Cell.GetForeground()); got != colorToHex(Nord.Foreground) {
		t.Errorf("Table.Cell: cor esperada %s, obteve %s", colorToHex(Nord.Foreground), got)
	}
	if got := colorToHex(styles.TextInput.Focused.Prompt.GetForeground()); got != colorToHex(Nord.Primary) {
		t.Errorf("TextInput.Focused.Prompt: cor esperada %s, obteve %s", colorToHex(Nord.Primary), got)
	}
	if got := colorToHex(styles.List.Title.GetForeground()); got != colorToHex(Nord.Background) {
		t.Errorf("List.Title: cor esperada %s, obteve %s", colorToHex(Nord.Background), got)
	}
	if got := colorToHex(styles.Help.ShortKey.GetForeground()); got != colorToHex(Nord.Accent) {
		t.Errorf("Help.ShortKey: cor esperada %s, obteve %s", colorToHex(Nord.Accent), got)
	}
	if styles.TextInput.Cursor.Color == nil {
		t.Error("o cursor do TextInput deveria ter cor definida")
	}
}

func TestBubblesEDeterministico(t *testing.T) {
	primeira := Nord.Bubbles().Table.Header.Render("ID")
	segunda := Nord.Bubbles().Table.Header.Render("ID")
	if primeira != segunda {
		t.Errorf("Bubbles() deveria ser determinístico:\n%q\n%q", primeira, segunda)
	}
}

func TestColorToHex(t *testing.T) {
	if got := colorToHex(lipgloss.Color("#1a2b3c")); got != "#1a2b3c" {
		t.Errorf("hex esperado '#1a2b3c', obteve '%s'", got)
	}
	if got := colorToHex(nil); got != "" {
		t.Errorf("hex de cor nula deveria ser vazio, obteve '%s'", got)
	}
}
