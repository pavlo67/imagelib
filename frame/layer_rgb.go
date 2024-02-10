package frame

import (
	"image"

	"github.com/pavlo67/common/common/imagelib"
	"github.com/pavlo67/common/common/imagelib/pix"
)

type ValueRGBA [imagelib.NumColorsRGBA]pix.Value

var _ imagelib.Described = &LayerRGBA{}

type LayerRGBA struct {
	image.RGBA
	imagelib.Settings
	// From LayerOuter
}

func (lyrRGB LayerRGBA) Description() imagelib.Settings {
	return lyrRGB.Settings
}
