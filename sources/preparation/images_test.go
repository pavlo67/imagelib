package preparation

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/common/common/filelib"

	"github.com/pavlo67/imagelib/imagelib"
)

func TestImages(t *testing.T) {

	if os.Getenv("T") != "" {
		t.Skip()
	}

	imagesPath := "/home/pavlo/0/partner/navigation_tests/_data/at_10_15/"

	numberedFiles, err := filelib.NumberedFilesSequence(imagesPath, RePGNStr, false)
	require.NoError(t, err)

	for i, nf := range numberedFiles {

		if i%10 == 0 {
			fmt.Println(i, " / ", len(numberedFiles))
		}

		//imgGray, err := imagelib.ReadGray(filepath.Join(imagesPath, fmt.Sprintf("%04d.pgm", descr.N)))

		filename := filepath.Join(imagesPath, fmt.Sprintf("%04d.pgm", nf.I))

		img, err := imagelib.ReadPGMSpecial(filename)
		require.NoError(t, err)
		require.NotNil(t, img)

	}

}
