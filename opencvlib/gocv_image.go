package opencvlib

import (
	"fmt"
	"github.com/pavlo67/common/common/logger"
	"github.com/pavlo67/common/common/mathlib/plane"
	"github.com/pavlo67/imagelib/imagelib"
	"image"
	"math"

	"gocv.io/x/gocv"

	"github.com/pavlo67/common/common/errors"
)

const onPrepare = "on Prepare()"

func Prepare(mat gocv.Mat, colorConversionCode gocv.ColorConversionCode, scale float64) (image.Image, error) {
	if !(scale > 0) || math.IsInf(scale, 1) {
		return nil, fmt.Errorf("wrong scale: %f / "+onPrepare, scale)
	}

	dims := mat.Size()
	if len(dims) != 2 {
		return nil, fmt.Errorf("wrong mat size: %+v / "+onPrepare, dims)
	}

	var matConverted *gocv.Mat

	if colorConversionCode < 0 {
		matConverted = &mat
	} else {
		matColorConverted := gocv.NewMat()
		defer matColorConverted.Close()

		gocv.CvtColor(mat, &matColorConverted, colorConversionCode)
		matConverted = &matColorConverted
	}

	if scale != 1 {
		matForResize := gocv.NewMat()
		defer matForResize.Close()

		gocv.Resize(*matConverted, &matForResize, image.Point{}, scale, scale, gocv.InterpolationDefault)
		matConverted = &matForResize
	}

	img, err := matConverted.ToImage()
	if err != nil {
		return nil, errors.Wrap(err, onPrepare)
	}

	return img, nil
}

const onPositionImage = "on PositionImage()"

func PositionImage(imgRGBA0 image.RGBA, scale0 float64, rotation plane.XToYAngle, imgSideX, imgSideY int, l logger.Operator) (*image.RGBA, error) {

	// TODO!!! don't convert gocv.Mat to imgRGBA

	imgRGBA, _, err := Resize(imgRGBA0, scale0)
	if err != nil {
		return nil, errors.Wrap(err, onPositionImage)
	} else if imgRGBA == nil {
		return nil, errors.New("imgRGBA == nil / " + onPositionImage)
	}

	var imgRGBRotated *image.RGBA

	if rotation == 0 {
		imgRGBRotated = imgRGBA

	} else {

		// TODO!!! why it doesn't work properly???
		// imgRGBRotated, err = Rotate(*imgRGBCentered, -rotation)

		imgRGBRotated, err = Rotate(*imagelib.ImageToRGBACopied(imgRGBA), float64(-rotation))
		if err != nil {
			return nil, errors.Wrap(err, onPositionImage)
		} else if imgRGBRotated == nil {
			return nil, errors.New("imgRGBRotated == nil / " + onPositionImage)
		}

	}

	if l != nil {
		l.Image("original.png", imagelib.GetImage1(&imgRGBA0))
		l.Image("resized.png", imagelib.GetImage1(imgRGBA))
		l.Image("rotated.png", imagelib.GetImage1(imgRGBRotated))
	}

	imgRect := image.Rectangle{Max: imgRGBRotated.Rect.Canon().Size()}
	imgRGBRotated.Rect = imgRect

	var imgRGBFinal *image.RGBA
	sub := false

	imgRectSub := imgRect

	if xDelta := imgRect.Max.X - imgSideX; xDelta > 0 {
		imgRectSub.Min.X += xDelta / 2
		imgRectSub.Max.X, sub = imgRectSub.Min.X+imgSideX, true
	}
	if yDelta := imgRect.Max.Y - imgSideY; yDelta > 0 {
		imgRectSub.Min.Y += yDelta / 2
		imgRectSub.Max.Y, sub = imgRectSub.Min.Y+imgSideY, true
	}

	if sub {
		imgRGBFinal, _ = imgRGBRotated.SubImage(imgRectSub).(*image.RGBA)

	} else {
		imgRGBFinal = imgRGBRotated

	}

	return imgRGBFinal, nil
}
