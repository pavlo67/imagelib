package frame

import (
	"math"

	"github.com/pavlo67/common/common/mathlib"
	"github.com/pavlo67/common/common/mathlib/plane"
)

type Frame struct {
	LayerRGBA
	plane.Position
}

// TODO!!! be careful:
// the origin of image.Rectangle is in the left top corner, Ox to right, Oy to bottom
// the geometry origin is in the rectangle center (when frame.pos is zero), geometry Ox to right, geometry Oy to top

func (frame Frame) PlaneRectangle() plane.Rectangle {
	var halfSideX, halfSideY float64

	if frame.DPM > 0 && !math.IsInf(frame.DPM, 1) {
		rect := frame.RGBA.Rect
		halfSideX, halfSideY = max(0, float64(rect.Max.X-rect.Min.X)/(2*frame.DPM)), max(0, float64(rect.Max.Y-rect.Min.Y)/(2*frame.DPM))
	}

	return plane.Rectangle{Position: frame.Position, HalfSideX: halfSideX, HalfSideY: halfSideY}
}

func (frame Frame) PointsToOuter(pChInner ...plane.Point2) plane.PolyChain {
	rect := frame.RGBA.Rect
	center := plane.Point2{0.5 * float64(rect.Min.X+rect.Max.X-1), 0.5 * float64(rect.Min.Y+rect.Max.Y-1)}
	pChOuter := make(plane.PolyChain, len(pChInner))
	for i, p := range pChInner {
		radiusOuter := math.Sqrt((p.X-center.X)*(p.X-center.X)+(p.Y-center.Y)*(p.Y-center.Y)) / frame.DPM
		if radiusOuter <= mathlib.Eps {
			pChOuter[i] = frame.Point2

		} else {
			angleInner := plane.Point2{p.X - center.X, p.Y - center.Y}.XToYAngleFromOx()
			angleOuter := frame.XToYAngle - angleInner
			pChOuter[i] = plane.Point2{frame.Point2.X + radiusOuter*math.Cos(float64(angleOuter)), frame.Point2.Y + radiusOuter*math.Sin(float64(angleOuter))}
		}
	}

	return pChOuter
}

func (frame Frame) PointToInner(p2Outer plane.Point2) plane.Point2 {
	radiusInner := math.Sqrt((p2Outer.X-frame.X)*(p2Outer.X-frame.X)+(p2Outer.Y-frame.Y)*(p2Outer.Y-frame.Y)) * frame.DPM
	var angleOuter plane.XToYAngle
	if radiusInner > mathlib.Eps {
		angleOuter = plane.Point2{p2Outer.X - frame.X, p2Outer.Y - frame.Y}.XToYAngleFromOx()
	}
	angleInner := frame.XToYAngle - angleOuter
	rect := frame.RGBA.Rect
	center := plane.Point2{0.5 * float64(rect.Min.X+rect.Max.X-1), 0.5 * float64(rect.Min.Y+rect.Max.Y-1)}

	return plane.Point2{center.X + radiusInner*math.Cos(float64(angleInner)), center.Y + radiusInner*math.Sin(float64(angleInner))}
}

// MovingToInner calculates the frame moving over fixed image (inner --> outer)
func (frame Frame) MovingToInner(movingOuter plane.Point2) plane.Point2 {
	movingRadiusInner := math.Sqrt(movingOuter.X*movingOuter.X+movingOuter.Y*movingOuter.Y) * frame.DPM
	movingAngleInner := frame.XToYAngle - movingOuter.XToYAngleFromOx()

	return plane.Point2{movingRadiusInner * math.Cos(float64(movingAngleInner)), movingRadiusInner * math.Sin(float64(movingAngleInner))}
}

// MovingToOuter calculates the frame moving over fixed image (outer --> inner)
func (frame Frame) MovingToOuter(movingInner plane.Point2) plane.Point2 {
	if !(frame.DPM > 0) {
		return plane.Point2{math.NaN(), math.NaN()}
	}

	movingRadiusOuter := math.Sqrt(movingInner.X*movingInner.X+movingInner.Y*movingInner.Y) / frame.DPM
	movingAngleOuter := frame.XToYAngle - movingInner.XToYAngleFromOx()

	return plane.Point2{movingRadiusOuter * math.Cos(float64(movingAngleOuter)), movingRadiusOuter * math.Sin(float64(movingAngleOuter))}
}
