package frame

import (
	"image"
	"math"
	"testing"

	"github.com/pavlo67/common/common/mathlib"
	"github.com/pavlo67/common/common/mathlib/plane"
	"github.com/pavlo67/imagelib/imagelib"
)

func TestFrame_PolyChainToOuter(t *testing.T) {
	tests := []struct {
		name             string
		Position         plane.Position
		DPM              float64
		RectInner        image.Rectangle
		pChInner         plane.PolyChain
		pChOuterExpected plane.PolyChain
	}{
		{
			name:             "",
			Position:         plane.Position{Point2: plane.Point2{1, 1}, XToYAngle: math.Pi},
			DPM:              1,
			RectInner:        image.Rectangle{Max: image.Point{5, 5}},
			pChInner:         plane.PolyChain{{0, 0}},
			pChOuterExpected: plane.PolyChain{{3, -1}},
		},
		{
			name:             "",
			Position:         plane.Position{Point2: plane.Point2{1, 1}, XToYAngle: math.Pi},
			DPM:              2,
			RectInner:        image.Rectangle{Max: image.Point{5, 5}},
			pChInner:         plane.PolyChain{{0, 0}},
			pChOuterExpected: plane.PolyChain{{2, 0}},
		},
		{
			name:             "",
			Position:         plane.Position{Point2: plane.Point2{1, 1}, XToYAngle: math.Pi},
			DPM:              2,
			RectInner:        image.Rectangle{Max: image.Point{5, 5}},
			pChInner:         plane.PolyChain{{2, 2}},
			pChOuterExpected: plane.PolyChain{{1, 1}},
		},
		{
			name:             "",
			Position:         plane.Position{Point2: plane.Point2{1, 1}, XToYAngle: math.Pi},
			DPM:              2,
			RectInner:        image.Rectangle{Max: image.Point{5, 3}},
			pChInner:         plane.PolyChain{{0, 0}},
			pChOuterExpected: plane.PolyChain{{2, 0.5}},
		},
		{
			name:             "",
			Position:         plane.Position{Point2: plane.Point2{1, 1}, XToYAngle: math.Pi / 2},
			DPM:              2,
			RectInner:        image.Rectangle{Max: image.Point{5, 3}},
			pChInner:         plane.PolyChain{{0, 0}},
			pChOuterExpected: plane.PolyChain{{0.5, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fr := Frame{
				LayerRGBA: LayerRGBA{
					RGBA:     image.RGBA{Rect: tt.RectInner},
					Settings: imagelib.Settings{DPM: tt.DPM},
				},
				Position: tt.Position,
			}
			pChOuter := fr.PointsToOuter(tt.pChInner...)

			for i, pOuter := range pChOuter {
				distance := pOuter.DistanceTo(tt.pChOuterExpected[i])
				if math.IsNaN(distance) || distance > mathlib.Eps {
					t.Errorf("PointsToOuter() = %v, pChOuterExpected %v", pChOuter, tt.pChOuterExpected)
				}

				pInner := fr.PointToInner(pOuter)
				distanceFinal := pInner.DistanceTo(tt.pChInner[i])
				if distanceFinal > mathlib.Eps {
					t.Errorf("pInner = %v, pOuter = %v, pOuterInner = %v", tt.pChInner[i], pOuter, pInner)
				}

			}
		})
	}
}
