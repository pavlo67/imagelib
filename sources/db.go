package sources

import (
	"fmt"
	"image"

	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/geolib"
	"github.com/pavlo67/common/common/mathlib/plane"
	"github.com/pavlo67/common/common/serialization"

	"github.com/pavlo67/imagelib/imagelib"
)

const SideX = 700
const SideY = 700
const ZoomDefault = 18

func SteppedPoint(point geolib.Point, step image.Point) geolib.Point {
	return point.MovedAt(plane.Point2{float64(step.X * SideX / 2), float64(step.Y * SideY / 2)})
}

func ImageName(geoPoint geolib.Point, zoom int) string {
	if zoom < 0 || zoom > ZoomDefault {
		zoom = ZoomDefault
	}
	return fmt.Sprintf("%g-%g-%02d.png", geoPoint.Lat, geoPoint.Lon, zoom)
}

const onGetSaved = "on GetSaved()"

func GetSaved(path string) (imgRGB *image.RGBA, descr *Description, err error) {
	//if old {
	//	descr, err = GetDescriptionByKey(path, key)

	descr = &Description{}
	err = serialization.Read(path+".json", serialization.MarshalerJSON, descr)

	if err != nil {
		return nil, nil, errors.Wrap(err, onGetSaved)
	} else if descr == nil {
		return nil, nil, errors.New("descr == nil / " + onGetSaved)
	}

	img, err := imagelib.Read(path)

	if err != nil {
		return nil, nil, errors.Wrap(err, onGetSaved)
	} else if descr == nil {
		return nil, nil, errors.New("img == nil / " + onGetSaved)
	}

	imgRGB, err = imagelib.ImageToRGBA(img)
	if err != nil {
		return nil, nil, errors.Wrap(err, onGetSaved)
	} else if descr == nil {
		return nil, nil, errors.New("img == nil / " + onGetSaved)
	}

	return imgRGB, descr, err
}
