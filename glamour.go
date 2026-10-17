package charmthemes

import (
	"charm.land/glamour/v2"
	glamouransi "charm.land/glamour/v2/ansi"
)

// Glamour devolve a configuração de estilos do glamour derivada do tema.
//
// O resultado pode ser usado diretamente com o glamour:
//
//	renderer, err := glamour.NewTermRenderer(glamour.WithStyles(theme.Glamour()))
//
// As cores de realce de sintaxe do tema são aplicadas ao Chroma dos blocos de
// código.
func (t Theme) Glamour() glamouransi.StyleConfig {
	return glamouransi.StyleConfig{
		Document: glamouransi.StyleBlock{
			StylePrimitive: glamouransi.StylePrimitive{
				BlockPrefix: "\n",
				BlockSuffix: "\n",
				Color:       hexPtr(t.Foreground),
			},
			Margin: uintPtr(2),
		},
		BlockQuote: glamouransi.StyleBlock{
			Indent:      uintPtr(1),
			IndentToken: stringPtr("│ "),
		},
		Paragraph: glamouransi.StyleBlock{},
		List: glamouransi.StyleList{
			LevelIndent: 2,
		},
		Heading: glamouransi.StyleBlock{
			StylePrimitive: glamouransi.StylePrimitive{
				BlockSuffix: "\n",
				Color:       hexPtr(t.Primary),
				Bold:        boolPtr(true),
			},
		},
		H1: glamouransi.StyleBlock{
			StylePrimitive: glamouransi.StylePrimitive{
				Prefix:          " ",
				Suffix:          " ",
				Color:           hexPtr(t.Primary),
				BackgroundColor: hexPtr(t.Background),
				Bold:            boolPtr(true),
			},
		},
		H2: glamouransi.StyleBlock{
			StylePrimitive: glamouransi.StylePrimitive{
				Color: hexPtr(t.Primary),
				Bold:  boolPtr(true),
			},
		},
		H3: glamouransi.StyleBlock{
			StylePrimitive: glamouransi.StylePrimitive{
				Color: hexPtr(t.Secondary),
				Bold:  boolPtr(true),
			},
		},
		H4: glamouransi.StyleBlock{
			StylePrimitive: glamouransi.StylePrimitive{
				Color: hexPtr(t.Accent),
				Bold:  boolPtr(true),
			},
		},
		H5: glamouransi.StyleBlock{
			StylePrimitive: glamouransi.StylePrimitive{
				Color: hexPtr(t.Muted),
				Bold:  boolPtr(true),
			},
		},
		H6: glamouransi.StyleBlock{
			StylePrimitive: glamouransi.StylePrimitive{
				Color: hexPtr(t.Muted),
			},
		},
		Text: glamouransi.StylePrimitive{
			Color: hexPtr(t.Foreground),
		},
		Strikethrough: glamouransi.StylePrimitive{
			CrossedOut: boolPtr(true),
		},
		Emph: glamouransi.StylePrimitive{
			Italic: boolPtr(true),
		},
		Strong: glamouransi.StylePrimitive{
			Bold: boolPtr(true),
		},
		HorizontalRule: glamouransi.StylePrimitive{
			Color:  hexPtr(t.Muted),
			Format: "\n--------\n",
		},
		Item: glamouransi.StylePrimitive{
			BlockPrefix: "• ",
		},
		Enumeration: glamouransi.StylePrimitive{
			BlockPrefix: ". ",
		},
		Task: glamouransi.StyleTask{
			Ticked:   "[✓] ",
			Unticked: "[ ] ",
		},
		Link: glamouransi.StylePrimitive{
			Color:     hexPtr(t.Info),
			Underline: boolPtr(true),
		},
		LinkText: glamouransi.StylePrimitive{
			Color: hexPtr(t.Accent),
			Bold:  boolPtr(true),
		},
		Image: glamouransi.StylePrimitive{
			Color:     hexPtr(t.Accent),
			Underline: boolPtr(true),
		},
		ImageText: glamouransi.StylePrimitive{
			Color:  hexPtr(t.Muted),
			Format: "Image: {{.text}} ",
		},
		Code: glamouransi.StyleBlock{
			StylePrimitive: glamouransi.StylePrimitive{
				Prefix:          "`",
				Suffix:          "`",
				Color:           hexPtr(t.Syntax.String),
				BackgroundColor: hexPtr(t.Background),
			},
		},
		CodeBlock: glamouransi.StyleCodeBlock{
			StyleBlock: glamouransi.StyleBlock{
				StylePrimitive: glamouransi.StylePrimitive{
					Color: hexPtr(t.Foreground),
				},
				Margin: uintPtr(2),
			},
			Chroma: t.chroma(),
		},
		Table:          glamouransi.StyleTable{},
		DefinitionList: glamouransi.StyleBlock{},
		DefinitionTerm: glamouransi.StylePrimitive{
			Color: hexPtr(t.Primary),
			Bold:  boolPtr(true),
		},
		DefinitionDescription: glamouransi.StylePrimitive{
			BlockPrefix: "\n→ ",
		},
		HTMLBlock: glamouransi.StyleBlock{},
		HTMLSpan:  glamouransi.StyleBlock{},
	}
}

// GlamourRenderer devolve um glamour.TermRenderer configurado com o tema.
//
// As opções passadas pelo chamador são aplicadas depois das opções do tema,
// portanto podem sobrescrevê-las:
//
//	renderer, err := theme.GlamourRenderer(glamour.WithWordWrap(100))
func (t Theme) GlamourRenderer(options ...glamour.TermRendererOption) (*glamour.TermRenderer, error) {
	opcoes := make([]glamour.TermRendererOption, 0, len(options)+1)
	opcoes = append(opcoes, glamour.WithStyles(t.Glamour()))
	opcoes = append(opcoes, options...)
	return glamour.NewTermRenderer(opcoes...)
}

// chroma devolve a configuração de realce de sintaxe do glamour derivada das
// cores de sintaxe do tema.
func (t Theme) chroma() *glamouransi.Chroma {
	return &glamouransi.Chroma{
		Text: glamouransi.StylePrimitive{
			Color: hexPtr(t.Foreground),
		},
		Error: glamouransi.StylePrimitive{
			Color: hexPtr(t.Error),
		},
		Comment: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Comment),
		},
		CommentPreproc: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Operator),
		},
		Keyword: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Keyword),
		},
		KeywordReserved: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Keyword),
		},
		KeywordNamespace: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Type),
		},
		KeywordType: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Type),
		},
		Operator: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Operator),
		},
		Punctuation: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Constant),
		},
		Name: glamouransi.StylePrimitive{
			Color: hexPtr(t.Foreground),
		},
		NameBuiltin: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Function),
		},
		NameTag: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Type),
		},
		NameAttribute: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Constant),
		},
		NameClass: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Type),
			Bold:  boolPtr(true),
		},
		NameConstant: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Constant),
		},
		NameDecorator: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Function),
		},
		NameException: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Keyword),
		},
		NameFunction: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Function),
		},
		NameOther: glamouransi.StylePrimitive{
			Color: hexPtr(t.Foreground),
		},
		Literal: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Constant),
		},
		LiteralNumber: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Number),
		},
		LiteralDate: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Constant),
		},
		LiteralString: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.String),
		},
		LiteralStringEscape: glamouransi.StylePrimitive{
			Color: hexPtr(t.Syntax.Operator),
		},
		GenericDeleted: glamouransi.StylePrimitive{
			Color: hexPtr(t.Error),
		},
		GenericEmph: glamouransi.StylePrimitive{
			Italic: boolPtr(true),
		},
		GenericInserted: glamouransi.StylePrimitive{
			Color: hexPtr(t.Success),
		},
		GenericStrong: glamouransi.StylePrimitive{
			Bold: boolPtr(true),
		},
		GenericSubheading: glamouransi.StylePrimitive{
			Color: hexPtr(t.Secondary),
			Bold:  boolPtr(true),
		},
		Background: glamouransi.StylePrimitive{
			BackgroundColor: hexPtr(t.Background),
		},
	}
}
