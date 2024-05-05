package images_files

import (
	"fmt"
	"image"

	"github.com/pavlo67/common/common/db"
	"github.com/pavlo67/common/common/errors"
	"github.com/pavlo67/common/common/filelib"
	"github.com/pavlo67/common/common/pnglib"
	"github.com/pavlo67/common/common/serialization"

	"github.com/pavlo67/imagelib/imagelib"
	"github.com/pavlo67/imagelib/images"
	"github.com/pavlo67/imagelib/sources"
)

var _ images.Operator = &imagesFiles{}

type imagesFiles struct {
	basePath string
	colored  bool
}

const onNew = "on images_files.New()"

func New(basePath string, colored bool) (images.Operator, db.Cleaner, error) {

	var err error
	if basePath, err = filelib.Dir(basePath); err != nil {
		return nil, nil, errors.Wrap(err, onNew)
	}

	imagesOp := imagesFiles{
		basePath: basePath,
		colored:  colored,
	}

	return &imagesOp, &imagesOp, nil
}

const onSave = "on imagesFiles.Save()"

func (op imagesFiles) Save(img image.Image, descr sources.Description, relPath string) error {
	imgPath, descrPath, err := images.RelPath(op.basePath, relPath, op.colored)
	if err != nil {
		return errors.Wrap(err, onSave)
	}

	if op.colored {
		if err := pnglib.Save(img, imgPath); err != nil {
			return fmt.Errorf("%s: %s / "+onSave, imgPath, err)
		}
	} else {
		imgGray, err := imagelib.ImageToGray(img)
		if err != nil {
			return fmt.Errorf("%s: %s / "+onSave, imgPath, err)
		} else if imgGray == nil {
			return fmt.Errorf("imgGray == nil: %s / "+onSave, imgPath)
		}

		if err = imagelib.SavePGM(imgGray, imgPath); err != nil {
			return fmt.Errorf("%s: %s / "+onSave, imgPath, err)
		}
	}

	if err = serialization.Save(descr, serialization.MarshalerJSON, descrPath); err != nil {
		return errors.Wrap(err, onSave)
	}

	return nil
}

const onCheck = "on imagesFiles.Check()"

func (op imagesFiles) Check(relPath string) (bool, error) {
	imgPath, descrPath, err := images.RelPath(op.basePath, relPath, op.colored)
	if err != nil {
		return false, errors.Wrap(err, onCheck)
	}

	imgExists, err := filelib.FileExists(imgPath, false)
	if err != nil {
		return false, fmt.Errorf("%s: %s / "+onCheck, imgPath, err)
	}

	if !imgExists {
		return false, nil
	}

	descrExists, err := filelib.FileExists(descrPath, false)
	if err != nil {
		return false, fmt.Errorf("%s: %s / "+onCheck, descrPath, err)
	} else if !descrExists {
		return false, fmt.Errorf("%s: descr is absent / "+onCheck, imgPath)
	}

	return true, nil
}

const onGet = "on imagesFiles.Get()"

func (op imagesFiles) Get(relPath string) (image.Image, *sources.Description, error) {
	imgPath, descrPath, err := images.RelPath(op.basePath, relPath, op.colored)
	if err != nil {
		return nil, nil, errors.Wrap(err, onGet)
	}

	descrExists, err := filelib.FileExists(descrPath, false)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %s / "+onGet, descrPath, err)
	} else if !descrExists {
		return nil, nil, nil
	}

	var descr sources.Description
	if err = serialization.Read(descrPath, serialization.MarshalerJSON, &descr); err != nil {
		return nil, nil, errors.Wrap(err, onGet)
	}

	imgExists, err := filelib.FileExists(imgPath, false)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %s / "+onGet, imgPath, err)
	} else if !imgExists {
		return nil, nil, nil
	}

	if op.colored {
		img, err := imagelib.Read(imgPath)
		if err != nil {
			return nil, nil, fmt.Errorf("%s: %s / "+onGet, imgPath, err)
		}
		return img, &descr, nil
	}

	imgGray, err := imagelib.ReadPGMSpecial(imgPath)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %s / "+onGet, imgPath, err)
	}
	return imgGray, &descr, nil
}
