package frame

import (
	"math"
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
			angleInner := plane.Point2{p.X - center.X, p.Y - center.Y}.Rotation()
			angleOuter := frame.Rotation - angleInner
			pChOuter[i] = plane.Point2{frame.Point2.X + radiusOuter*math.Cos(float64(angleOuter)), frame.Point2.Y + radiusOuter*math.Sin(float64(angleOuter))}
		}

		// log.Printf("center: %v", center)
		// log.Printf("radiusOuter: %f", radiusOuter)
		// log.Printf("angleInner: %f (%f degrees)", angleInner, angleInner*180/math.Pi)
		// log.Printf("angleOuter: %f (%f degrees)", angleOuter, angleOuter*180/math.Pi)
		// log.Printf("p (%v) --> pChOuter[i]: %v", p, pChOuter[i])
	}

	return pChOuter
}

func (frame Frame) PointToInner(p2 plane.Point2) plane.Point2 {
	radius := math.Sqrt((p2.X-frame.X)*(p2.X-frame.X)+(p2.Y-frame.Y)*(p2.Y-frame.Y)) * frame.DPM
	var angle plane.Rotation
	if radius > mathlib.Eps {
		angle = plane.Point2{p2.X - frame.X, p2.Y - frame.Y}.Rotation()
	}
	angleInternal := frame.Rotation - angle
	rect := frame.RGBA.Rect
	center := plane.Point2{0.5 * float64(rect.Min.X+rect.Max.X-1), 0.5 * float64(rect.Min.Y+rect.Max.Y-1)}

	return plane.Point2{center.X + radius*math.Cos(float64(angleInternal)), center.Y + radius*math.Sin(float64(angleInternal))}
}

func (frame Frame) MotionToInner(p2 plane.Point2) plane.Point2 {
	radius := math.Sqrt((p2.X-frame.X)*(p2.X-frame.X)+(p2.Y-frame.Y)*(p2.Y-frame.Y)) * frame.DPM
	var angle plane.Rotation
	if radius > mathlib.Eps {
		angle = plane.Point2{p2.X - frame.X, p2.Y - frame.Y}.Rotation()
	}
	angleInternal := frame.Rotation - angle

	return plane.Point2{radius * math.Cos(float64(angleInternal)), radius * math.Sin(float64(angleInternal))}
}
