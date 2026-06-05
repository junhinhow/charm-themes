# charm-themes

Unified theming system for [charmbracelet](https://github.com/charmbracelet) libraries.

Define a color palette once, use it everywhere — lipgloss, glamour, bubbles.

> **[Leia em Portugues](README.pt-br.md)**

## The Problem

Each charmbracelet library has its own styling system:
- **lipgloss** uses `lipgloss.Style` with inline color definitions
- **glamour** uses `StyleConfig` with JSON or Go structs
- **bubbles** components each have their own `Styles` struct

When building a TUI app, you end up defining the same colors 3+ times in different formats. Changing your theme means updating every single style definition.

## The Solution

```go
import themes "github.com/junhinhow/charm-themes"

// Pick a theme
theme := themes.Gruvbox

// Use with lipgloss
styles := theme.Lipgloss()
fmt.Println(styles.Title.Render("Hello"))
fmt.Println(styles.Error.Render("Something went wrong"))
fmt.Println(styles.Border.Render("Boxed content"))

// Use semantic colors directly
myStyle := lipgloss.NewStyle().
    Foreground(theme.Primary).
    Background(theme.Background)
```

## Built-in Themes

| Theme | Background | Primary | Accent |
|-------|-----------|---------|--------|
| **Gruvbox** | `#282828` | `#fabd2f` | `#83a598` |
| **Dracula** | `#282a36` | `#bd93f9` | `#ff79c6` |
| **Nord** | `#2e3440` | `#88c0d0` | `#b48ead` |
| **Catppuccin** | `#1e1e2e` | `#cba6f7` | `#f5c2e7` |
| **Tokyo Night** | `#1a1b26` | `#7aa2f7` | `#bb9af7` |

## Theme Structure

```go
type Theme struct {
    Name       string
    Background color.Color
    Foreground color.Color
    Primary    color.Color  // headings, titles
    Secondary  color.Color  // subtitles, secondary info
    Accent     color.Color  // highlights, emphasis
    Success    color.Color  // green/positive
    Warning    color.Color  // yellow/caution
    Error      color.Color  // red/negative
    Info       color.Color  // blue/informational
    Muted      color.Color  // gray/disabled
    Syntax     SyntaxColors // code highlighting
}
```

## License

MIT
