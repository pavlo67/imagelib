package convolution

import (
	"fmt"
	"github.com/pavlo67/common/common/imagelib"
	"github.com/pavlo67/imagelib/coloring"
	frame2 "github.com/pavlo67/imagelib/frame"
	"github.com/pavlo67/imagelib/layers"
	"github.com/pavlo67/imagelib/pix"
	"image"

	"github.com/pavlo67/common/common"
)

var _ Mask = &colorRangeMask{}

type colorRangeMask struct {
	imgRGB *image.RGBA
	coloring.ColorRange
	// cnt int32
}

const onColorRange = "on ColorRange()"

func ColorRange(colorRange coloring.ColorRange) Mask {
	return &colorRangeMask{ColorRange: colorRange}
}

func (mask *colorRangeMask) Side() int {
	return 1
}

const onColorRangePrepare = "on ColorRange.GetNext()"

func (mask *colorRangeMask) Prepare(onData interface{}) error {
	switch v := onData.(type) {
	case image.RGBA:
		mask.imgRGB = &v
	case *image.RGBA:
		mask.imgRGB = v
	case frame2.LayerRGBA:
		mask.imgRGB = &v.RGBA
	case *frame2.LayerRGBA:
		mask.imgRGB = &v.RGBA
	case frame2.Frame:
		mask.imgRGB = &v.RGBA
	case *frame2.Frame:
		mask.imgRGB = &v.RGBA
	}
	if mask.imgRGB == nil {
		if onData == nil {
			return fmt.Errorf("onData == nil (%#v) / "+onColorRangePrepare, onData)
		}
		return fmt.Errorf("wrong data (%T) / "+onColorRangePrepare, onData)
	}

	// mask.cnt = 0
	return nil
}

func (mask colorRangeMask) Info() common.Map {
	return common.Map{
		"name":     "cr_" + mask.ColorRange.String(),
		"colorMin": mask.ColorRange.ColorMin,
		"colorMax": mask.ColorRange.ColorMax,
	}
}

func (mask colorRangeMask) Classes() layers.Classes {
	return nil
}

func (mask *colorRangeMask) Calculate(x, y int) pix.Value {
	imgRGB := mask.imgRGB
	offset := (y-imgRGB.Rect.Min.Y)*imgRGB.Stride + (x-imgRGB.Rect.Min.X)*imagelib.NumColorsRGBA
	r, g, b := imgRGB.Pix[offset], imgRGB.Pix[offset+1], imgRGB.Pix[offset+2]

	if r >= mask.ColorRange.ColorMin.R && r <= mask.ColorRange.ColorMax.R &&
		g >= mask.ColorRange.ColorMin.G && g <= mask.ColorRange.ColorMax.G &&
		b >= mask.ColorRange.ColorMin.B && b <= mask.ColorRange.ColorMax.B {
		return pix.ValueMax
	}

	return 0
}
