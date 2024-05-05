package images_files_jlist

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/imagelib/images"
)

func TestImagesFiles(t *testing.T) {

	os.Setenv("ENV", "test")

	imagesOp, imagesCleanerOp, err := New("./test/", true)
	require.NoError(t, err)
	require.NotNil(t, imagesOp)
	require.NotNil(t, imagesCleanerOp)

	images.OperatorTestScenario(t, imagesOp, imagesCleanerOp)
}
