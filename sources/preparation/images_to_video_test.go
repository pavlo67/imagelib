package preparation

import (
	"image"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pavlo67/common/common/filelib"

	"github.com/pavlo67/imagelib/opencvlib"
	"github.com/pavlo67/imagelib/video"
)

func TestImagesToVideo(t *testing.T) {

	if os.Getenv("T") != "" {
		t.Skip()
	}

	// this logFile parameter determines a file for complete logger output duplication
	//envs, _ := config.PrepareTests(t, "../_env/", "")
	//require.NotNil(t, envs)

	// data ---------------------------------------------------------------------

	dataPath, err := filelib.GetDir("../../shifter/shifter_tests/_data/odometry/arcgis_0/")
	require.NoError(t, err)

	//jlistAllPath := filepath.Join(dataPath, shifter.FramesDescriptionFile)
	//var descrs []sources.Description
	//err = serialization.ReadAllPartsJSON(jlistAllPath, &descrs)
	//require.NoError(t, err)
	//require.Truef(t, len(descrs) > 0, "descrs = []")

	cfg := &video.Info{
		Rectangle: image.Rectangle{Max: image.Point{X: 1300, Y: 1040}},
		FPS:       3,
	}

	re, err := regexp.Compile(`^\d{4}\.png$`)
	require.NoError(t, err)
	require.NotNil(t, re)

	err = opencvlib.WriteMP4(filepath.Join(dataPath, filepath.Base(dataPath)+".mp4"), dataPath, *re, cfg.FPS, cfg.Rectangle.Max.X, cfg.Rectangle.Max.Y, true)
	require.NoError(t, err)

}
