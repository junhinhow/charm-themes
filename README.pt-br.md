# charm-themes

Sistema de temas unificado para as bibliotecas [charmbracelet](https://github.com/charmbracelet).

Defina uma paleta de cores uma vez, use em todos os lugares — lipgloss, glamour e bubbles.

> **[Read in English](README.md)**

## O Problema

Cada biblioteca charmbracelet tem seu próprio sistema de estilos:
- **lipgloss** usa `lipgloss.Style` com cores inline
- **glamour** usa uma `StyleConfig` com JSON ou structs Go
- **bubbles** cada componente tem seu próprio struct `Styles`

Ao construir um app TUI, você acaba definindo as mesmas cores 3+ vezes em formatos diferentes. Mudar o tema significa atualizar cada definição de estilo.

## A Solução

```go
import (
    "fmt"

    "charm.land/lipgloss/v2"

    themes "github.com/junhinhow/charm-themes"
)

// Escolha um tema
theme := themes.Gruvbox

// Use com lipgloss — chame uma única vez e reutilize o resultado
styles := theme.Lipgloss()
fmt.Println(styles.Title.Render("Olá"))
fmt.Println(styles.Error.Render("Algo deu errado"))
fmt.Println(styles.Border.Render("Conteúdo em caixa"))

// Ou use as cores diretamente em estilos customizados
myStyle := lipgloss.NewStyle().
    Foreground(theme.Primary).
    Background(theme.Background)
```

## Instalação

```bash
go get github.com/junhinhow/charm-themes
```

## Temas Incluídos

| Tema | Background | Primary | Accent |
|------|-----------|---------|--------|
| **Gruvbox** | `#282828` | `#fabd2f` | `#83a598` |
| **Dracula** | `#282a36` | `#bd93f9` | `#ff79c6` |
| **Nord** | `#2e3440` | `#88c0d0` | `#b48ead` |
| **Catppuccin** | `#1e1e2e` | `#cba6f7` | `#f5c2e7` |
| **Tokyo Night** | `#1a1b26` | `#7aa2f7` | `#bb9af7` |

## Busca de Temas

```go
// Busca por nome (ignora maiúsculas e minúsculas)
theme, ok := themes.Get("tokyo-night")
if !ok {
    // tema inexistente
}

// Todos os temas, indexados pelo nome
all := themes.Themes()

// Nomes em ordem determinística (alfabética)
names := themes.ThemeNames()
```

`Themes()` e `Get()` devolvem **cópias** dos temas embutidos, portanto customizar o valor
retornado nunca corrompe os temas globais do pacote (`themes.Nord`, `themes.Gruvbox`, …).

## Estrutura do Tema

```go
type Theme struct {
    Name       string
    Background color.Color  // fundo principal
    Foreground color.Color  // texto principal
    Primary    color.Color  // títulos, headings
    Secondary  color.Color  // subtítulos, info secundária
    Accent     color.Color  // destaques, ênfase
    Success    color.Color  // verde/positivo
    Warning    color.Color  // amarelo/cuidado
    Error      color.Color  // vermelho/negativo
    Info       color.Color  // azul/informativo
    Muted      color.Color  // cinza/desabilitado
    Syntax     SyntaxColors // cores de syntax highlighting
}
```

## Validação

Um tema montado à mão pode estar sem cores. `Validate()` informa todos os papéis ausentes
em vez de renderizar silenciosamente sem cor:

```go
if err := theme.Validate(); err != nil {
    return err
}

// As cores de sintaxe também podem ser validadas separadamente
if err := theme.Syntax.Validate(); err != nil {
    return err
}
```

## Estilos Lipgloss Pré-Construídos

```go
styles := theme.Lipgloss()

styles.Base        // texto normal
styles.Bold        // negrito
styles.Italic      // itálico
styles.Faint       // desbotado

styles.Success     // mensagem de sucesso
styles.Warning     // aviso
styles.Error       // erro
styles.Info        // informação

styles.Title       // título (primary, bold)
styles.Subtitle    // subtítulo (secondary)
styles.Highlight   // destaque (accent, bold)
styles.Border      // caixa com borda arredondada
styles.Code        // código inline
```

> `Lipgloss()` constrói 13 valores de `lipgloss.Style` a cada chamada. Chame **uma vez**
> na inicialização e reutilize o resultado — chamá-la dentro do laço de renderização
> desperdiça alocações.

## Integração com o Glamour

`Glamour()` devolve uma `glamour/ansi.StyleConfig` derivada do tema, incluindo uma
configuração do Chroma para os blocos de código:

```go
import "charm.land/glamour/v2"

renderer, err := glamour.NewTermRenderer(glamour.WithStyles(theme.Glamour()))
if err != nil {
    return err
}

out, err := renderer.Render("# Olá\n")
```

O método de conveniência `GlamourRenderer()` faz os dois de uma vez e aceita opções
adicionais:

```go
renderer, err := theme.GlamourRenderer(glamour.WithWordWrap(100))
```

## Integração com o Bubbles

`Bubbles()` devolve os estilos dos componentes mais comuns do bubbles, prontos para serem
aplicados aos componentes:

```go
styles := theme.Bubbles()

styles.List       // componente list
styles.Table      // componente table
styles.TextInput  // componente textinput
styles.TextArea   // componente textarea
styles.Help       // componente help

// O list expõe os estilos como um campo público...
l := list.New(items, list.NewDefaultDelegate(), 80, 24)
l.Styles = styles.List

// ...enquanto os demais componentes usam SetStyles.
t := table.New(table.WithColumns(columns))
t.SetStyles(styles.Table)
```

> Assim como `Lipgloss()`, `Bubbles()` constrói todos os estilos do zero a cada chamada.
> Chame **uma vez** e reutilize o resultado.

## Roadmap

- [x] Integração com lipgloss
- [x] Integração com glamour (StyleConfig + Chroma)
- [x] Integração com bubbles (list, table, textinput, textarea, help)
- [ ] Integração com huh (forms)
- [ ] Detecção automática de fundo claro/escuro
- [ ] Carregamento de temas customizados via JSON/YAML

## Licença

MIT
