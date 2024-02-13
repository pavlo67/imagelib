package opencvlib

import (
	"image"

	"gocv.io/x/gocv"
	"golang.org/x/image/colornames"

	"github.com/pavlo67/common/common/errors"
)

func ConvexHull(points []image.Point) []image.Point {
	matHull := gocv.NewMat()
	defer matHull.Close()

	gocv.ConvexHull(gocv.NewPointVectorFromPoints(points), &matHull, false, false)

	// TODO: be careful!!!
	// if !returnPoints: matHull.T() == CV32S   (== []int == list of contour indices, required for gocv.ConvexityDefects())
	// if  returnPoints: matHull.T() == CV32SC2 (== []image.Point???)

	hullLength := matHull.Size()[0]
	convexHullIndices, convexHullPoints := make([]int, hullLength), make([]image.Point, hullLength)

	for i := 0; i < hullLength; i++ {
		convexHullIndices[i] = int(matHull.GetIntAt(i, 0))
		convexHullPoints[i] = points[convexHullIndices[i]]

		if i > 0 && convexHullIndices[i] < convexHullIndices[i-1] {
			// The convex hull indices are not monotonous, which can be in the case when the input contour contains self-intersections
		}
	}

	// pr.ConvexHullArea = gocv.ContourArea(gocv.NewPointVectorFromPoints(pr.ConvexHullPoints)) / (ls.DPM * ls.DPM)

	return convexHullPoints
}

const onFillOutsideContours = "on imagelib.GrayOutsideContours()"

func WhiteOutsideContours(imgGray image.Gray, psv gocv.PointsVector) (*image.Gray, error) {

	matImg, err := gocv.ImageGrayToMatGray(&imgGray)
	if err != nil {
		return nil, errors.Wrap(err, onFillOutsideContours)
	}

	matMaskCntrs := gocv.NewMatWithSize(imgGray.Rect.Dy(), imgGray.Rect.Dx(), gocv.MatTypeCV8U)
	gocv.FillPoly(&matMaskCntrs, psv, colornames.White)

	matMaskCntrsOutside := gocv.NewMat()
	gocv.BitwiseNot(matMaskCntrs, &matMaskCntrsOutside)

	matWhitedOutside := gocv.NewMat()
	gocv.BitwiseOr(matImg, matMaskCntrsOutside, &matWhitedOutside)

	imgWhitedOutside, err := matWhitedOutside.ToImage()
	if err != nil {
		return nil, errors.Wrap(err, onFillOutsideContours)
	}

	imgGrayWhitedOutside, _ := imgWhitedOutside.(*image.Gray)
	if imgGrayWhitedOutside == nil {
		return nil, errors.New("imgGrayWhitedOutside == nil / " + onFillOutsideContours)
	}

	return imgGrayWhitedOutside, nil
}

func BlackOutsideContours(imgGray image.Gray, psv gocv.PointsVector) (*image.Gray, error) {

	matImg, err := gocv.ImageGrayToMatGray(&imgGray)
	if err != nil {
		return nil, errors.Wrap(err, onFillOutsideContours)
	}

	matMaskCntrs := gocv.NewMatWithSize(imgGray.Rect.Dy(), imgGray.Rect.Dx(), gocv.MatTypeCV8U)
	gocv.FillPoly(&matMaskCntrs, psv, colornames.White)

	matImgMasked := gocv.NewMat()
	matImg.CopyToWithMask(&matImgMasked, matMaskCntrs)

	imgBlackedOutside, err := matImgMasked.ToImage()
	if err != nil {
		return nil, errors.Wrap(err, onFillOutsideContours)
	}

	imgGrayBlackedOutside, _ := imgBlackedOutside.(*image.Gray)
	if imgGrayBlackedOutside == nil {
		return nil, errors.New("imgGrayBlackedOutside == nil / " + onFillOutsideContours)
	}

	return imgGrayBlackedOutside, nil
}

//const onFill = "on areas.contoursPV.Fill()"
//
//func (cntrs contours.contoursPV) Fill(rect image.Rectangle) (*image.Gray, error) {
//	matFilled := gocv.NewMatWithSize(rect.Max.Y-rect.Min.Y, rect.Max.X-rect.Min.X, gocv.MatTypeCV8U)
//	defer matFilled.Close()
//
//	pointsVector := gocv.NewPointsVector()
//	for _, contour := range cntrs {
//		pointsVector.Append(gocv.NewPointVectorFromPoints(contour.PointsSub(rect.Min)))
//	}
//
//	gocv.FillPoly(&matFilled, pointsVector, colornames.White)
//
//	imgFilled, err := matFilled.ToImage()
//	if err != nil {
//		return nil, errors.Wrap(err, onFill)
//	} else if imgFilled == nil {
//		return nil, errors.New("imgFilled == nil / " + onFill)
//	}
//
//	imgFilledGray, _ := imgFilled.(*image.Gray)
//	if imgFilledGray == nil {
//		return nil, fmt.Errorf("imgFilledGray == nil (%T) / "+onFill, imgFilled)
//	}
//
//	imgFilledGray.Rect = rect
//	return imgFilledGray, nil
//}

//const onCoreCenter = "on Contour.CoreCenter()"
//
//func (cntr contours.Contour) CoreCenter(coreCheckSide int, returnCoreLayer bool) (*plane.Point2, *layers.Layer, error) {
//
//	if coreCheckSide <= 0 {
//		return nil, nil, fmt.Errorf("wrong coreCheckSide (%d) / "+onCoreCenter, coreCheckSide)
//	}
//
//	rect := imagelib.RectangleAround(coreCheckSide, cntr.Points...)
//
//	dx, dy := rect.Max.X-rect.Min.X, rect.Max.Y-rect.Min.Y
//	if dx <= 0 || dy <= 0 {
//		return nil, nil, fmt.Errorf("wrong rect (%v) / "+onCoreCenter, rect)
//	}
//
//	matMaskCntrs := gocv.NewMatWithSize(dy, dx, gocv.MatTypeCV8U)
//	cntrPoints := imagelib.ConvertImagePoints(cntr.Points, false, image.Point{-rect.Min.X, -rect.Min.Y}, 1)
//	gocv.FillPoly(&matMaskCntrs, gocv.NewPointsVectorFromPoints([][]image.Point{cntrPoints}), colornames.White)
//
//	img, err := matMaskCntrs.ToImage()
//	if err != nil {
//		return nil, nil, errors.Wrap(err, onCoreCenter)
//	}
//	imgGray, _ := img.(*image.Gray)
//	if imgGray == nil {
//		return nil, nil, fmt.Errorf("img (%T) isn't non-nil *image.Gray / "+onCoreCenter, img)
//	}
//
//	maskAveraging := convolution.AveragingLeftTop(coreCheckSide)
//	if maskAveraging == nil {
//		return nil, nil, fmt.Errorf("AveragingLeftTop(%d) == nil / "+onCoreCenter, coreCheckSide)
//	}
//
//	lyr0 := layers.Layer{
//		Gray: *imgGray,
//	}
//
//	lyr, err := convolution.Layer(&lyr0, maskAveraging, 1, false)
//	if err != nil {
//		return nil, nil, errors.Wrap(err, onCoreCenter)
//	} else if lyr == nil {
//		return nil, nil, fmt.Errorf("lyr == nil / " + onCoreCenter)
//	}
//
//	center, err := lyr.Center(lyr.Max)
//	if err != nil {
//		return nil, nil, errors.Wrap(err, onCoreCenter)
//	} else if center == nil {
//		return nil, nil, fmt.Errorf("center == nil / " + onCoreCenter)
//	}
//
//	var lyrCore *layers.Layer
//	if returnCoreLayer {
//		lyrCoreFull, _ := lyr.Thresholded(lyr.Max, false)
//		if lyrCoreFull != nil {
//			shift := image.Point{(coreCheckSide - 1) / 2, (coreCheckSide - 1) / 2}
//			lyrCoreFull.Rect = image.Rectangle{rect.Min.Add(shift), rect.Max.Add(shift)}
//			rectThr, _ := lyrCoreFull.RectThr(lyr.Max)
//			if rectThr != nil {
//				lyrCore = lyrCoreFull.SubLayer(*rectThr)
//			}
//		}
//	}
//
//	// log.Printf("cntr.Points: %v, rect: %v, center: %v, coreCheckSide: %d", cntr.Points, rect, center, coreCheckSide)
//
//	return &plane.Point2{
//		float64(rect.Min.X) + center.X + float64(coreCheckSide-1)*0.5,
//		float64(rect.Min.Y) + center.Y + float64(coreCheckSide-1)*0.5,
//	}, lyrCore, nil
//
//}
