package convolution

import (
	"fmt"
	"math"
	"strconv"

	"github.com/pavlo67/common/common/imagelib/pix"

	"github.com/pavlo67/common/common"
)

var _ Mask = &averagingLeftTopMask{}

type averagingLeftTopMask struct {
	lyr  *methods.Layer
	side int
}

func AveragingLeftTop(side int) Mask {
	if side < 1 {
		side = 1
	}

	return &averagingLeftTopMask{
		side: side,
	}
}

const onAveragingPrepare = "on AveragingLeftTop.Prepare()"

func (mask *averagingLeftTopMask) Prepare(onData interface{}) error {
	switch v := onData.(type) {
	case methods.Layer:
		mask.lyr = &v
	case *methods.Layer:
		mask.lyr = v
	}
	if mask.lyr == nil {
		if onData == nil {
			return fmt.Errorf("onData == nil (%#v) / "+onAveragingPrepare, onData)
		}
		return fmt.Errorf("wrong data (%T) / "+onAveragingPrepare, onData)
	}

	return nil
}

func (mask averagingLeftTopMask) Stat() interface{} {
	return nil
}

func (mask averagingLeftTopMask) Info() common.Map {
	return common.Map{
		"name": "avg_" + strconv.Itoa(mask.side),
		"side": mask.side,
	}
}

func (mask averagingLeftTopMask) Calculate(x, y int) pix.Value {

	lyr := mask.lyr
	if lyr == nil || x < lyr.Rect.Min.X || x >= lyr.Rect.Max.X || y < lyr.Rect.Min.Y || y >= lyr.Rect.Max.Y {
		return pix.ValueMiddle
	}

	xMax, yMax := min(x+mask.side, lyr.Rect.Max.X), min(y+mask.side, lyr.Rect.Max.Y)
	offset, vSum := (y-lyr.Rect.Min.Y)*lyr.Stride+(-lyr.Rect.Min.X), pix.ValueSum(0)
	for yi := y; yi < yMax; yi++ {
		for xi := x; xi < xMax; xi++ {
			vSum += pix.ValueSum(lyr.Pix[offset+xi])
		}
		offset += lyr.Stride
	}

	return pix.Value(math.Round(float64(vSum) / float64((yMax-y)*(xMax-x))))
}
