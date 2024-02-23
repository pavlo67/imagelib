package imagelib

import (
	"fmt"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/pnglib"
	"image"
	_ "image/jpeg"
	"os"
)

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

const onImageGray = "on ImageGray()"

func ImageGray(srcFilename, resFilename string, rect *image.Rectangle) (imgArea *image.Gray, original *image.Rectangle, err error) {
	srcFile, err := os.Open(srcFilename)
	if err != nil {
		// return nil, nil, errors.Wrapf(err, "on opening %s / "+onImageGray, srcFilename)
		return nil, nil, errors.Wrap(err, onImageGray)
	}
	defer srcFile.Close()

	src, _, err := image.Decode(srcFile)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "on decoding %s / "+onImageGray, srcFilename)
	}

	switch img := src.(type) {
	//case *img.Gray16:
	//	t.Logf("%T", img)
	case *image.Gray:
		if rect == nil {
			return img, &img.Rect, nil
		}

		imgArea := img.SubImage(*rect).(*image.Gray)

		if resFilename != "" {
			if err = pnglib.Save(imgArea, resFilename); err != nil {
				fmt.Fprint(os.Stderr, err)
			}
		}

		return imgArea, &img.Rect, nil
	}

	return nil, nil, fmt.Errorf("wrong format (%T) of %s / "+onImageGray, src, srcFilename)
}

const onImageRGBA = "on ImageRGBA()"

func ImageRGBA(filename string, rect *image.Rectangle) (imgRGBA *image.RGBA, original *image.Rectangle, err error) {
	srcFile, err := os.Open(filename)
	if err != nil {
		// return nil, nil, errors.Wrapf(err, "on opening %s / "+onImageRGBA, filename)
		return nil, nil, errors.Wrap(err, onImageRGBA)
	}
	defer srcFile.Close()

	src, _, err := image.Decode(srcFile)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "on decoding %s / "+onImageRGBA, filename)
	}

	switch img := src.(type) {
	case *image.RGBA:
		imgRGBA = img

	case *image.NRGBA:
		imgRGBA = &image.RGBA{
			Pix:    img.Pix,
			Stride: img.Stride,
			Rect:   img.Rect,
		}

	case *image.Gray:
		imgRGBA = image.NewRGBA(img.Rect)
		for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
			for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
				imgRGBA.Set(x, y, img.At(x, y))
			}
		}

	case *image.YCbCr:
		imgRGBA = image.NewRGBA(img.Rect)
		for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
			for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
				imgRGBA.Set(x, y, img.At(x, y))
			}
		}

	case *image.CMYK:
		imgRGBA = image.NewRGBA(img.Rect)
		for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
			for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
				imgRGBA.Set(x, y, img.At(x, y))
			}
		}

	default:
		return nil, nil, fmt.Errorf("wrong format (%T) of %s / "+onImageRGBA, src, filename)

	}

	if rect != nil {
		imgRGBA = imgRGBA.SubImage(*rect).(*image.RGBA)
	}

	//if resFilename != "" {
	//	if err = SavePNG(imgRGBA, resFilename); err != nil {
	//		fmt.Fprint(os.Stderr, err)
	//	}
	//}

	return imgRGBA, &imgRGBA.Rect, nil
}
