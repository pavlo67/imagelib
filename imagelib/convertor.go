package imagelib

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/pavlo67/common/common/errors"
)

const NumColorsRGB = 3
const NumColorsRGBA = 4

const ChRed = 0
const ChGreen = 1
const ChBlue = 2

const onImageToGray = "on ImageToGray()"

func ImageToGray(img image.Image) (*image.Gray, error) {
	if img == nil {
		return nil, fmt.Errorf(onImageToGray + ": nil img")
	}

	if grayPtr, ok := img.(*image.Gray); ok {
		if grayPtr != nil {
			return grayPtr, nil
		}
		return nil, fmt.Errorf(onImageToGray + ": nil img.(*Gray)")
	}

	rect := img.Bounds()
	gray := image.NewGray(rect)

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			gray.Set(x, y, img.At(x, y))
		}
	}

	return gray, nil
}

func ImageToGrayCopied(img image.Image) *image.Gray {
	if img == nil {
		return nil
	}

	rect := img.Bounds()
	gray := image.NewGray(rect)

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			gray.Set(x, y, img.At(x, y))
		}
	}

	return gray
}

const onImageToRGBA = "on ImageToRGBA()"

func ImageToRGBA(img image.Image) (*image.RGBA, error) {
	if img == nil {
		return nil, fmt.Errorf(onImageToRGBA + ": nil img")
	}

	if rgbaPtr, ok := img.(*image.RGBA); ok {
		if rgbaPtr != nil {
			return rgbaPtr, nil
		}
		return nil, fmt.Errorf(onImageToRGBA + ": nil img.(*RGBA)")
	}

	rect := img.Bounds()
	rgba := image.NewRGBA(rect)

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	return rgba, nil
}

func ImageToRGBACopied(img image.Image) *image.RGBA {
	if img == nil {
		return nil
	}

	rect := img.Bounds()
	rgba := image.NewRGBA(rect)

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	return rgba
}

const onReadImage = "on ReadImage()"

func ReadImage(srcFilename string) (image.Image, error) {
	srcFile, err := os.Open(srcFilename)
	if err != nil {
		return nil, errors.Wrap(err, onReadImage)
	}
	defer srcFile.Close()

	img, _, err := image.Decode(srcFile)
	if err != nil {
		return nil, errors.Wrapf(err, "on decoding %s / "+onReadImage, srcFilename)
	}

	return img, nil
}

const onRGBToGray = "on RGBToGray()"

func RGBToGray(rgba image.RGBA, colorNum int) (*image.Gray, error) {
	if colorNum >= NumColorsRGB {
		return nil, fmt.Errorf(onRGBToGray+": wrong color to get: %d", colorNum)
	}

	xWidth, yHeight := rgba.Rect.Max.X-rgba.Rect.Min.X, rgba.Rect.Max.Y-rgba.Rect.Min.Y
	if xWidth <= 0 || yHeight <= 0 {
		return nil, fmt.Errorf(onRGBToGray+": empty img.Rect (%#v)", rgba.Rect)
	}

	gray := image.Gray{
		Pix:    make([]uint8, xWidth*yHeight),
		Stride: xWidth,
		Rect:   image.Rectangle{Max: image.Point{X: xWidth, Y: yHeight}},
	}

	if colorNum >= 0 {
		for y := 0; y < yHeight; y++ {
			rgbaStride := (y-rgba.Rect.Min.Y)*rgba.Stride + colorNum
			grayStride := y * xWidth
			for x := 0; x < xWidth; x++ {
				gray.Pix[grayStride+x] = rgba.Pix[rgbaStride+(x-rgba.Rect.Min.X)*NumColorsRGBA]
			}
		}
	} else {
		for y := 0; y < yHeight; y++ {
			for x := 0; x < xWidth; x++ {
				gray.Set(x+rgba.Rect.Min.X, y+rgba.Rect.Min.Y, rgba.At(x, y))
			}
		}
	}

	return &gray, nil
}

func GrayValueAvg(img *image.RGBA, x, y float64, csSelected int) float64 {
	dx, dy := x-float64(int(x)), y-float64(int(y))
	if dx > 0 {
		if dy > 0 {
			return float64(img.Pix[int(y)*img.Stride+int(x)*NumColorsRGBA+csSelected])*(1-dx)*(1-dy) +
				float64(img.Pix[(int(y)+1)*img.Stride+int(x)*NumColorsRGBA+csSelected])*(1-dx)*dy +
				float64(img.Pix[(int(y)+1)*img.Stride+(int(x)+1)*NumColorsRGBA+csSelected])*dx*dy +
				float64(img.Pix[int(y)*img.Stride+(int(x)+1)*NumColorsRGBA+csSelected])*dx*(1-dy)
		}
		return float64(img.Pix[int(y)*img.Stride+int(x)*NumColorsRGBA+csSelected])*(1-dx) +
			float64(img.Pix[int(y)*img.Stride+(int(x)+1)*NumColorsRGBA+csSelected])*dx
	} else {
		if dy > 0 {
			return float64(img.Pix[int(y)*img.Stride+int(x)*NumColorsRGBA+csSelected])*(1-dy) +
				float64(img.Pix[(int(y)+1)*img.Stride+int(x)*NumColorsRGBA+csSelected])*dy
		}
		return float64(img.Pix[int(y)*img.Stride+int(x)*NumColorsRGBA+csSelected])
	}
}

func GrayValue(img *image.RGBA, x, y, csSelected int) float64 {
	return float64(img.Pix[y*img.Stride+x*NumColorsRGBA+csSelected])
}
