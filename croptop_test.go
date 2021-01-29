package croptop_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/croptop"
)

func TestImage(t *testing.T) {
	_, err := croptop.Decode(new(bytes.Buffer))
	assert.Error(t, err)

	r, err := os.Open("./fixtures/alicia-vikander.jpg")
	require.NoError(t, err)

	img, err := croptop.Decode(r)
	require.NoError(t, err)

	height := 691.2
	width := 460.8
	offsetX := 434.7
	offsetY := 50.48

	buf := new(bytes.Buffer)
	err = img.Width(width).Height(height).
		OffsetX(offsetX).OffsetY(offsetY).
		Crop().Encode(buf)
	assert.NoError(t, err)

	// verify the result for image
	out, err := imaging.Decode(buf)
	require.NoError(t, err)

	rect := out.Bounds()
	assert.Equal(t, 461, rect.Max.X)
	assert.Equal(t, 691, rect.Max.Y)
}
