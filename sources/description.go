package sources

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/geolib"

	"github.com/pavlo67/imagelib/frame"
	"github.com/pavlo67/imagelib/layers/convolution"
)

type Key string

type ImageRef struct {
	ImagePath string
	SourceKey Key
}

func WrappedGeoPoint(geoPoint geolib.Point, sourcesPath string, sourceKey Key) ImageRef {
	return ImageRef{
		ImagePath: filepath.Join(sourcesPath, fmt.Sprintf("%g-%g-18.png", geoPoint.Lat, geoPoint.Lon)),
		SourceKey: sourceKey,
	}

}

const onParse = "on ImageRef.Parse()"

func (imageRef ImageRef) Parse() (*geolib.Point, int, error) {
	parts := strings.Split(filepath.Base(imageRef.ImagePath), "-")

	latFloat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "parsing latFloat (0) in %+v / "+onParse, imageRef)
	}

	lonFloat, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "parsing lonFloat (1) in %+v / "+onParse, imageRef)
	}

	part2Splitted := strings.Split(parts[2], ".")
	zoom, err := strconv.Atoi(part2Splitted[0])
	if err != nil {
		return nil, 0, errors.Wrapf(err, "parsing zoom (5) in %+v / "+onParse, imageRef)
	}

	return &geolib.Point{geolib.Degrees(latFloat), geolib.Degrees(lonFloat)}, zoom, nil

}

type Description struct {
	ImageRef  `                json:",inline"`
	GeoPoint  *geolib.Point  `json:",omitempty"`
	Bearing   geolib.Bearing `json:",omitempty"`
	DPM       float64
	PointsRaw []frame.PointRawGeo
	convolution.ClassesMetrics
}
