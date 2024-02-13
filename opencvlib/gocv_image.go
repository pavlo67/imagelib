package opencvlib

import (
	"fmt"
	"image"
	"math"

	"gocv.io/x/gocv"

	"github.com/pavlo67/common/common/errors"
)

const onPrepare = "on video.Prepare()"

func Prepare(mat gocv.Mat, scale, rotation float64) (image.Image, error) {
	if !(scale > 0) || math.IsInf(scale, 1) {
		return nil, fmt.Errorf("wrong scale: %f / "+onPrepare, scale)
	} else if math.IsInf(rotation, 0) || math.IsNaN(rotation) {
		return nil, fmt.Errorf("wrong rotation: %f / "+onPrepare, rotation)
	}

	dims := mat.Size()
	if len(dims) != 2 {
		return nil, fmt.Errorf("wrong mat size: %+v / "+onPrepare, dims)
	}
	size := image.Point{dims[0], dims[1]}

	// TODO??? use native scaling & rotation

	var matFinal *gocv.Mat
	if rotation == 0 {
		if scale == 1 {
			matFinal = &mat

		} else {
			matForResize := gocv.NewMat()
			defer matForResize.Close()

			gocv.Resize(mat, &matForResize, image.Point{}, scale, scale, gocv.InterpolationDefault)
			matFinal = &matForResize
		}

	} else {
		matForRotation := gocv.NewMat()
		defer matForRotation.Close()

		rotationMatrix := gocv.GetRotationMatrix2D(image.Point{size.X / 2, size.Y / 2}, rotation, scale)
		gocv.WarpAffine(mat, &matForRotation, rotationMatrix, size)
		matFinal = &matForRotation
	}

	img, err := matFinal.ToImage()
	if err != nil {
		return nil, errors.Wrap(err, onPrepare)
	}

	return img, nil
}
