package convolution

import (
	"fmt"
	"strconv"

	"github.com/pavlo67/common/common/imagelib/pix"

	"github.com/pavlo67/common/common"
)

var _ Mask = &dilationMask{}

type dilationMask struct {
	lyr    *methods.Layer
	radius int
}

func Dilation(radius int) Mask {
	if radius < 1 {
		radius = 1
	}

	return &dilationMask{
		radius: radius,
	}
}

const onDilationPrepare = "on Dilation.Prepare()"

func (mask *dilationMask) Prepare(onData interface{}) error {
	switch v := onData.(type) {
	case methods.Layer:
		mask.lyr = &v
	case *methods.Layer:
		mask.lyr = v
	}
	if mask.lyr == nil {
		if onData == nil {
			return fmt.Errorf("onData == nil (%#v) / "+onDilationPrepare, onData)
		}
		return fmt.Errorf("wrong data (%T) / "+onDilationPrepare, onData)
	}
	if len(mask.lyr.Pix) < (mask.lyr.Rect.Max.Y-mask.lyr.Rect.Min.Y)*mask.lyr.Stride {
		return fmt.Errorf("wrong lyr.Pix (%d) for lyr.Rect (%v / %d) / "+onDilationPrepare, len(mask.lyr.Pix), mask.lyr.Rect, mask.lyr.Stride)
	}

	return nil
}

func (mask dilationMask) Stat() interface{} {
	return nil
}

func (mask dilationMask) Info() common.Map {
	return common.Map{
		"name":   "dil_" + strconv.Itoa(mask.radius),
		"radius": mask.radius,
	}
}

func (mask dilationMask) Calculate(x, y int) pix.Value {

	lyr := mask.lyr
	if lyr == nil {
		return 0
	}

	xMin, yMin := max(x-mask.radius, lyr.Rect.Min.X), max(y-mask.radius, lyr.Rect.Min.Y)
	xMax, yMax := min(x+mask.radius+1, lyr.Rect.Max.X), min(y+mask.radius+1, lyr.Rect.Max.Y)
	offset := (yMin-lyr.Rect.Min.Y)*lyr.Stride + (-lyr.Rect.Min.X)
	for yi := yMin; yi < yMax; yi++ {
		for xi := xMin; xi < xMax; xi++ {
			if lyr.Pix[offset+xi] > 0 {
				return pix.ValueMax
			}
		}
		offset += lyr.Stride
	}

	return 0
}
