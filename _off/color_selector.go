package _off

type ColorSelector struct {
	NumColors   int
	Selected    int
	SelectedMax int
}

var ColorSelectorFullRGB = ColorSelector{
	NumColors:   imagelib.NumColorsRGBA,
	Selected:    0,
	SelectedMax: imaging.NumColorsRGB,
}
