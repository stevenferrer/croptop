package croptop

import (
	"image"
	"io"
	"math"

	"github.com/disintegration/imaging"
)

// Image is an image
type Image struct {
	img     image.Image
	width   int
	height  int
	offsetX int
	offsetY int
}

// Decode reads the image from r
func Decode(r io.Reader) (*Image, error) {
	img, err := imaging.Decode(r)
	if err != nil {
		return nil, err
	}

	return &Image{img: img}, nil
}

// Width sets the bounding box width
func (i *Image) Width(width float64) *Image {
	i.width = int(math.Round(width))
	return i
}

// Height sets the bounding box height
func (i *Image) Height(height float64) *Image {
	i.height = int(math.Round(height))
	return i
}

// OffsetX sets the x-coordinate offset
func (i *Image) OffsetX(x float64) *Image {
	i.offsetX = int(math.Round(x))
	return i
}

// OffsetY sets the y-coordinate offset
func (i *Image) OffsetY(y float64) *Image {
	i.offsetY = int(math.Round(y))
	return i
}

// Crop crops and returns a new image
func (i *Image) Crop() *Image {
	rect := image.Rect(0, 0, i.width, i.height).
		Add(image.Pt(i.offsetX, i.offsetY))
	return &Image{img: imaging.Crop(i.img, rect)}
}

// Encode writes the image to w
func (i *Image) Encode(w io.Writer) error {
	return imaging.Encode(w, i.img, imaging.JPEG)
}
