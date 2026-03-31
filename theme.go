// Package charmthemes provides a unified theming system for charmbracelet
// libraries. It defines color palettes that can be applied consistently across
// lipgloss styles, glamour markdown rendering, and bubbles components.
//
// Instead of configuring colors separately for each library, define a Theme
// once and derive all styles from it.
package charmthemes

import "image/color"

// Theme defines a complete color palette that can be applied across all
// charmbracelet libraries. Each theme provides semantic color roles that
// map to different UI contexts.
type Theme struct {
	Name string

	// Base colors
	Background color.Color
	Foreground color.Color

	// Accent colors for headings, links, emphasis
	Primary   color.Color
	Secondary color.Color
	Accent    color.Color

	// Semantic colors
	Success color.Color
	Warning color.Color
	Error   color.Color
	Info    color.Color
	Muted   color.Color

	// Syntax highlighting (for code blocks)
	Syntax SyntaxColors
}

// SyntaxColors defines colors for code syntax highlighting.
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
