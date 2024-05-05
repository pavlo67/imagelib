package images_files

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/imagelib/images"
)

func TestImagesFiles(t *testing.T) {

	os.Setenv("ENV", "test")

	imagesOp, imagesCleanerOp, err := New("./test/", true)
	require.NotNil(t, imagesOp)
	require.NotNil(t, imagesCleanerOp)
	require.NoError(t, err)

	images.OperatorTestScenario(t, imagesOp, imagesCleanerOp)
}
