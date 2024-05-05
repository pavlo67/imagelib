package images

import (
	"fmt"
	"github.com/pavlo67/imagelib/sources"
	"image"
	"path/filepath"

	"github.com/pavlo67/common/common/joiner"
)

const InterfaceKey joiner.InterfaceKey = "images"
const InterfaceCleanerKey joiner.InterfaceKey = "images_cleaner"

type Operator interface {
	Get(path string) (image.Image, *sources.Description, error)
	Save(_ image.Image, _ sources.Description, path string) error
	Check(path string) (bool, error)
}

func RelPath(basePath, relPath string, colored bool) (string, string, error) {
	relImgPath := filepath.Clean(relPath)
	if len(relImgPath) < 1 || relImgPath[0] == '.' {
		return "", "", fmt.Errorf("wrong relative path: '%s'", relPath)
	}

	var ext string
	if colored {
		ext = ".png"
	} else {
		ext = ".pgm"
	}

	if base := filepath.Base(relImgPath); len(base) < 4 || base[len(base)-1:] != ext {
		relImgPath += ext
	}

	imgPath := filepath.Join(basePath, relImgPath)

	return imgPath, imgPath + ".json", nil
}
