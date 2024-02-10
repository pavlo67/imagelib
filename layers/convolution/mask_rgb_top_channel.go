package convolution

import (
	"fmt"
	frame2 "github.com/pavlo67/imagelib/frame"
	"image"
	"strconv"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/imagelib/imagelib"

	"github.com/pavlo67/imagelib/imagelib/pix"
)

var _ Mask = &topChannelMask{}

type topChannelMask struct {
	imgRGB   *image.RGBA
	ch       int
	ch1, ch2 int
	thr      pix.ValueDelta
}

const onTopChannel = "on convolution.TopChannel()"

func TopChannel(ch int, thr pix.ValueDelta) Mask {
	return &topChannelMask{
		ch:  ch,
		ch1: (ch + 1) % 3,
		ch2: (ch + 2) % 3,
		thr: thr,
	}
}

func (mask *topChannelMask) Side() int {
	return 1
}

const onTopChannelPrepare = "on topChannelMask.Prepare()"

func (mask *topChannelMask) Prepare(onData interface{}) error {
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
			return fmt.Errorf("onData == nil (%#v)", onData)
		}
		return fmt.Errorf("wrong data (%T)", onData)
	}

	return nil
}

func (mask topChannelMask) Info() common.Map {
	return common.Map{
		"name": "top_ch_" + strconv.Itoa(int(mask.thr)),
		"thr":  mask.thr,
	}
}

func (mask topChannelMask) Stat() interface{} {
	return nil
}

func (mask topChannelMask) Calculate(x, y int) pix.Value {
	imgRGB := mask.imgRGB

	offset := (y-imgRGB.Rect.Min.Y)*imgRGB.Stride + (x-imgRGB.Rect.Min.X)*imagelib.NumColorsRGBA
	if offset < 0 || offset >= len(imgRGB.Pix)-imagelib.NumColorsRGB {
		return 0
	}

	rgb := imgRGB.Pix[offset : offset+imagelib.NumColorsRGB]

	var clrOtherMax uint8
	if rgb[mask.ch1] > rgb[mask.ch2] {
		clrOtherMax = rgb[mask.ch1]
	} else {
		clrOtherMax = rgb[mask.ch2]
	}

	vDelta := pix.ValueDelta(rgb[mask.ch]) - pix.ValueDelta(clrOtherMax) - mask.thr
	if vDelta > 0 {
		return pix.Value(vDelta)
	}

	return 0
}
