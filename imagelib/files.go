package imagelib

import (
	"fmt"
	pnm "github.com/jbuchbinder/gopnm"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/filelib"
)

const onRead = "on Read()"

func Read(filename string) (image.Image, error) {

	srcFile, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, onRead)
	}
	defer srcFile.Close()

	img, _, err := image.Decode(srcFile)
	if err != nil {
		return nil, errors.Wrapf(err, "on decoding %s / "+onRead, filename)
	}

	return img, nil
}

const onReadGray = "on imagelib.ReadGray()"

func ReadGray(srcFilename string) (*image.Gray, error) {
	srcFile, err := os.Open(srcFilename)
	if err != nil {
		return nil, errors.Wrap(err, onReadGray)
	}
	defer srcFile.Close()

	src, _, err := image.Decode(srcFile)
	if err != nil {
		return nil, errors.Wrapf(err, "on decoding %s / "+onReadGray, srcFilename)
	}

	imgGray, err := ImageToGray(src)
	if err != nil {
		return nil, errors.Wrap(err, onReadGray)
	}

	return imgGray, nil
}

const onReadRGBA = "on imagelib.ReadRGBA()"

func ReadRGBA(filename string) (imgRGBA *image.RGBA, _ error) {
	srcFile, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, onReadRGBA)
	}
	defer srcFile.Close()

	src, _, err := image.Decode(srcFile)
	if err != nil {
		return nil, errors.Wrapf(err, "on decoding %s / "+onReadRGBA, filename)
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
		return nil, fmt.Errorf("wrong format (%T) of %s / "+onReadRGBA, src, filename)

	}

	return imgRGBA, nil
}

const onSavePGM = "on imagelib.SavePGM()"

func SavePGM(img image.Image, filename string) error {
	if img == nil {
		return errors.New("img == nil / " + onSavePGM)
	} else if path := filepath.Dir(filename); path != "" && path != "." && path != ".." {
		if _, err := filelib.Dir(path); err != nil {
			return errors.Wrapf(err, "can't create dir '%s' / "+onSavePGM, path)
		}
	}

	resFile, err := os.Create(filename)
	if err != nil {
		return errors.Wrap(err, onSavePGM)
	}
	defer resFile.Close()

	if err = pnm.Encode(resFile, img, pnm.PGM); err != nil {
		return errors.Wrap(err, onSavePGM)
	}
	return nil
}
