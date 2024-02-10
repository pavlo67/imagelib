package convolution

import (
	"fmt"
	"image"
	"strconv"

	"github.com/pavlo67/common/common"
	"github.com/pavlo67/imagelib/imagelib"
	"github.com/pavlo67/imagelib/imagelib/pix"

	"github.com/pavlo67/imagelib/frame"
)

var _ Mask = &RGBBrightnessClassesMask{}

type RGBBrightnessClassesMask struct {
	imgRGB *image.RGBA
	stat   ClassesMetrics
}

type ClassesMetrics struct {
	ClassesNum int
	ClassRange pix.Value
	Cnt        []int32
}

const onRGBBrightnessClasses = "on RGBBrightnessClasses()"

func RGBBrightnessClasses(classRange pix.Value) (Mask, error) {
	if classRange == 0 {
		return nil, fmt.Errorf("classRange == 0 / " + onRGBBrightnessClasses)
	}
	classesNum := (int(pix.ValueMax) + 1) / int(classRange)
	if pix.ValueMax%classRange > 0 {
		classesNum++
	}

	return &RGBBrightnessClassesMask{
		stat: ClassesMetrics{
			ClassRange: classRange,
			ClassesNum: classesNum,
			Cnt:        make([]int32, classesNum),
		},
	}, nil
}

func (mask *RGBBrightnessClassesMask) Side() int {
	return 1
}

const onBrightnessPrepare = "on RGBBrightnessClasses.GetNext()"

func (mask *RGBBrightnessClassesMask) Prepare(onData interface{}) error {
	switch v := onData.(type) {
	case image.RGBA:
		mask.imgRGB = &v
	case *image.RGBA:
		mask.imgRGB = v
	case frame.LayerRGBA:
		mask.imgRGB = &v.RGBA
	case *frame.LayerRGBA:
		mask.imgRGB = &v.RGBA
	case frame.Frame:
		mask.imgRGB = &v.RGBA
	case *frame.Frame:
		mask.imgRGB = &v.RGBA
	}
	if mask.imgRGB == nil {
		if onData == nil {
			return fmt.Errorf("onData == nil (%#v) / "+onBrightnessPrepare, onData)
		}
		return fmt.Errorf("wrong data (%T) / "+onBrightnessPrepare, onData)
	}

	mask.stat.Cnt = make([]int32, mask.stat.ClassesNum)

	return nil
}

func (mask RGBBrightnessClassesMask) Info() common.Map {
	return common.Map{
		"name":       "br_classes_" + strconv.Itoa(int(mask.stat.ClassRange)),
		"classRange": mask.stat.ClassRange,
	}
}

func (mask RGBBrightnessClassesMask) Stat() interface{} {
	return mask.stat
}

func (mask *RGBBrightnessClassesMask) Calculate(x, y int) pix.Value {
	offset := (y-mask.imgRGB.Rect.Min.Y)*mask.imgRGB.Stride + (x-mask.imgRGB.Rect.Min.X)*imagelib.NumColorsRGBA

	brightness := pix.ValueSum(mask.imgRGB.Pix[offset]) + pix.ValueSum(mask.imgRGB.Pix[offset+1]) + pix.ValueSum(mask.imgRGB.Pix[offset+2])

	classNum := pix.Value(brightness / (3 * pix.ValueSum(mask.stat.ClassRange)))
	mask.stat.Cnt[classNum]++

	return mask.stat.ClassRange * classNum
}
