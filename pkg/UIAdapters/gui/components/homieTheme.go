package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

// Fyne Theme Implementation
type HsTheme struct {
}
func (m *HsTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		if v == theme.VariantLight {
			return &color.NRGBA{0xcf, 0xd8, 0xdc, 0xff}
		}
		return &color.NRGBA{0x45, 0x5A, 0x64, 0xff}
	case theme.ColorNameFocus:
		return &color.NRGBA{0xff, 0xc1, 0x07, 0xff}
	case theme.ColorNameForeground:
		return &color.NRGBA{44, 69, 232, 255}
	}

	return theme.DefaultTheme().Color(n, v)
}

func (t *HsTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

func (t *HsTheme) Font(fyne.TextStyle) fyne.Resource {
	return theme.DefaultTextMonospaceFont()
}

func (t *HsTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}
// end theme implementation

