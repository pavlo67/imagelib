package opencvlib

import (
	"fmt"
	"image"
	"math"

	"gocv.io/x/gocv"

	"github.com/pavlo67/common/common/errors"
)

const onPrepare = "on opencvlib.Prepare()"

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
