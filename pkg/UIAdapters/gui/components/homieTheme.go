package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	cc "github.com/skoona/homie-service/pkg/utils"
	"image/color"
)

// Fyne Theme Implementation
type homieTheme struct {
}

func NewHomieTheme() fyne.Theme {
	return &homieTheme{}
}

func (m *homieTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		return cc.Purple
	case theme.ColorNameButton, theme.ColorNameDisabled:
		return color.Black
	case theme.ColorNamePlaceHolder, theme.ColorNameScrollBar:
		return cc.Grey
	case theme.ColorNamePrimary, theme.ColorNameHover, theme.ColorNameFocus:
		return cc.Orange
	case theme.ColorNameShadow:
		return &color.RGBA{R: 0xcc, G: 0xcc, B: 0xcc, A: 0xcc}
	default:
		return color.White
	}

	return theme.DefaultTheme().Color(n, v)
}

func (t *homieTheme) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNamePadding:
		return 10
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNameScrollBar:
		return 10
	case theme.SizeNameScrollBarSmall:
		return 5
	case theme.SizeNameText:
		return 12
	default:
		return 0
	}
}

func (t *homieTheme) Font(fyne.TextStyle) fyne.Resource {
	return theme.DefaultTextMonospaceFont()
}

func (t *homieTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}
// end theme implementation

