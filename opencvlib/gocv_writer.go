package opencvlib

import (
	"fmt"
	"regexp"

	"gocv.io/x/gocv"

	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/filelib"
)

const onWriteMP4 = "on videolib.WriteMP4()"

func WriteMP4(resultFilename, sourcePath string, sourceRegexp regexp.Regexp, fps float64, xWidth, yHeight int, isColor bool) error {
	fourcc, err := gocv.VideoWriterFile(resultFilename, "mp4v", fps, xWidth, yHeight, isColor)
	if err != nil {
		return errors.Wrap(err, onWriteMP4)
	} else if fourcc == nil {
		return errors.New("fourcc == nil / " + onWriteMP4)
	}

	filenames, err := filelib.List(sourcePath, &sourceRegexp, false, true)
	if err != nil {
		return errors.Wrap(err, onWriteMP4)
	}

	for i, filename := range filenames {
		img := gocv.IMRead(filename, gocv.IMReadAnyColor)
		fmt.Printf("#%d: %s --> %+v\n", i, filename, img.Size())

		if !fourcc.IsOpened() {
			return errors.New("fourcc is not open / " + onWriteMP4)
		} else if err = fourcc.Write(img); err != nil {
			return errors.Wrapf(err, "on writing image from %s / "+onWriteMP4, filename)
		}
	}

	if err = fourcc.Close(); err != nil {
		return errors.Wrap(err, onWriteMP4)
	}

	return nil
}
