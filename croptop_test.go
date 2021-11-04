package croptop_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/stevenferrer/croptop"
)

func TestCrop(t *testing.T) {
	_, err := croptop.Decode(new(bytes.Buffer))
	assert.Error(t, err)

	r, err := os.Open("./fixtures/alicia-vikander.jpg")
	require.NoError(t, err)

	img, err := croptop.Decode(r)
	require.NoError(t, err)

	height := 296.0
	width := 197.0
	offsetX := 666.0
	offsetY := 66.0

	buf := &bytes.Buffer{}
	err = img.Height(height).Width(width).
		OffsetX(offsetX).OffsetY(offsetY).
		Crop().Encode(buf)
	require.NoError(t, err)

	// verify the dimensions of the resulting image
	out, err := imaging.Decode(buf)
	require.NoError(t, err)

	rect := out.Bounds()
	// height
	assert.Equal(t, 296, rect.Max.Y)
	// width
	assert.Equal(t, 197, rect.Max.X)
}
