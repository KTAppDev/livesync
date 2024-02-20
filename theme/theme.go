package theme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type MyTheme struct {
	fyne.Theme
}

func NewMyTheme() fyne.Theme {
	return &MyTheme{Theme: theme.DefaultTheme()}
}

func (m *MyTheme) Color(n fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return m.Theme.Color(n, theme.VariantDark)
}
