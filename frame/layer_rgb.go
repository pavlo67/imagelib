package frame

import (
	"image"

	"github.com/pavlo67/imagelib/imagelib"
	"github.com/pavlo67/imagelib/imagelib/pix"
)

type ValueRGBA [imagelib.NumColorsRGBA]pix.Value

var _ imagelib.Described = &LayerRGBA{}

type LayerRGBA struct {
	image.RGBA
	imagelib.Settings
}

func (lyrRGB LayerRGBA) Description() imagelib.Settings {
	return lyrRGB.Settings
}
