package charmthemes

import (
	"fmt"
	"image/color"
)

// colorToHex converte uma cor para a representação hexadecimal "#rrggbb".
//
// Devolve string vazia quando a cor não está definida.
func colorToHex(c color.Color) string {
	if c == nil {
		return ""
	}
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
}

// hexPtr devolve um ponteiro para a representação hexadecimal da cor, ou nil
// quando a cor não está definida.
func hexPtr(c color.Color) *string {
	hex := colorToHex(c)
	if hex == "" {
		return nil
	}
	return &hex
}

// boolPtr devolve um ponteiro para o valor booleano informado.
func boolPtr(b bool) *bool {
	return &b
}

// uintPtr devolve um ponteiro para o valor sem sinal informado.
func uintPtr(u uint) *uint {
	return &u
}

// stringPtr devolve um ponteiro para o texto informado.
func stringPtr(s string) *string {
	return &s
}
