package layers

import (
	"github.com/pavlo67/imagelib/imagelib"
	"github.com/pavlo67/imagelib/imagelib/pix"
	"image"
)

var _ imagelib.Described = &Layer{}

type Layer struct {
	image.Gray
	imagelib.Settings
	Metrics
}

type Layers map[string]*Layer

func (lyr Layer) Description() imagelib.Settings {
	return lyr.Settings
}

func (lyr *Layer) Offset(x, y int) int {
	return (y-lyr.Rect.Min.Y)*lyr.Stride + (x - lyr.Rect.Min.X)
}

func (lyr Layer) Length() int64 {
	return int64(lyr.Rect.Max.Y-lyr.Rect.Min.Y) * int64(lyr.Rect.Max.X-lyr.Rect.Min.X)
}

type Metrics struct {
	Min, Max, Avg pix.Value
	BlRat, WhRat  float64
	Criterion     float64 // TODO!!! keep it for shifting
	Classes       Classes
}

type Classes []int32

func (classes Classes) Range() pix.Value {
	if len(classes) <= 0 {
		return 0

	}

	r := pix.Value(256 / len(classes))
	if pix.Value(256%len(classes)) >= r {
		return 0
	}

	return r
}
