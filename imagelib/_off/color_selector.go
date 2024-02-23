package _off

import "github.com/pavlo67/imagelib/imagelib"

type ColorSelector struct {
	NumColors   int
	Selected    int
	SelectedMax int
}

var ColorSelectorFullRGB = ColorSelector{
	NumColors:   imagelib.NumColorsRGBA,
	Selected:    0,
	SelectedMax: imagelib.NumColorsRGB,
}
