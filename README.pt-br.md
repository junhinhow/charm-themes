# charm-themes

Sistema de temas unificado para as bibliotecas [charmbracelet](https://github.com/charmbracelet).

Defina uma paleta de cores uma vez, use em todos os lugares — lipgloss, glamour, bubbles.

> **[Read in English](README.md)**

## O Problema

Cada biblioteca charmbracelet tem seu proprio sistema de estilos:
- **lipgloss** usa `lipgloss.Style` com cores inline
- **glamour** usa `StyleConfig` com JSON ou structs Go
- **bubbles** cada componente tem seu proprio struct `Styles`

Ao construir um app TUI, voce acaba definindo as mesmas cores 3+ vezes em formatos diferentes. Mudar o tema significa atualizar cada definicao de estilo.

## A Solucao

```go
import themes "github.com/junhinhow/charm-themes"

// Escolha um tema
theme := themes.Gruvbox

// Use com lipgloss — estilos semanticos prontos
styles := theme.Lipgloss()
fmt.Println(styles.Title.Render("Ola"))
fmt.Println(styles.Error.Render("Algo deu errado"))
fmt.Println(styles.Border.Render("Conteudo em caixa"))

// Ou use as cores diretamente em estilos customizados
myStyle := lipgloss.NewStyle().
    Foreground(theme.Primary).
    Background(theme.Background)
```

## Instalacao

```bash
go get github.com/junhinhow/charm-themes
```

## Temas Inclusos

| Tema | Background | Primary | Accent |
|------|-----------|---------|--------|
| **Gruvbox** | `#282828` | `#fabd2f` | `#83a598` |
| **Dracula** | `#282a36` | `#bd93f9` | `#ff79c6` |
| **Nord** | `#2e3440` | `#88c0d0` | `#b48ead` |
| **Catppuccin** | `#1e1e2e` | `#cba6f7` | `#f5c2e7` |
| **Tokyo Night** | `#1a1b26` | `#7aa2f7` | `#bb9af7` |

## Estrutura do Tema

```go
type Theme struct {
    Name       string
    Background color.Color  // fundo principal
    Foreground color.Color  // texto principal
    Primary    color.Color  // titulos, headings
    Secondary  color.Color  // subtitulos, info secundaria
    Accent     color.Color  // destaques, enfase
    Success    color.Color  // verde/positivo
    Warning    color.Color  // amarelo/cuidado
    Error      color.Color  // vermelho/negativo
    Info       color.Color  // azul/informativo
    Muted      color.Color  // cinza/desabilitado
    Syntax     SyntaxColors // cores de syntax highlighting
}
```

## Estilos Lipgloss Pre-Construidos

```go
styles := theme.Lipgloss()

styles.Base        // texto normal
styles.Bold        // negrito
styles.Italic      // italico
styles.Faint       // desbotado

styles.Success     // mensagem de sucesso
styles.Warning     // aviso
styles.Error       // erro
styles.Info        // informacao

styles.Title       // titulo (primary, bold)
styles.Subtitle    // subtitulo (secondary)
styles.Highlight   // destaque (accent, bold)
styles.Border      // caixa com borda arredondada
styles.Code        // codigo inline
```

## Roadmap

- [x] Integracao com lipgloss
- [ ] Integracao com glamour (StyleConfig + Chroma)
- [ ] Integracao com bubbles (table, list, textinput)
- [ ] Integracao com huh (forms)
- [ ] Deteccao automatica de fundo claro/escuro
- [ ] Carregamento de temas customizados via JSON/YAML

## Licenca

MIT
