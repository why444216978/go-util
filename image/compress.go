package image

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"

	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
)

type Options struct {
	Quality         int
	Size            int
	Compressibility float64
	MaxDepth        int
	ResampleFilter  imaging.ResampleFilter
}

func Compress(r io.Reader, o *Options) (image.Image, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	if o == nil {
		return img, nil
	}

	if o.Compressibility == 0 {
		o.Compressibility = 0.75
	}

	if o.MaxDepth == 0 {
		o.MaxDepth = 3
	}

	return compress(img, o, 1)
}

func compress(img image.Image, o *Options, depth int) (image.Image, error) {
	size, err := getSize(img, o.Quality)
	if err != nil {
		return nil, err
	}

	if size <= o.Size {
		return img, nil
	}

	if depth > o.MaxDepth {
		return nil, errors.Wrap(err, "mat depth")
	}

	bounds := img.Bounds()
	width := int(float64(bounds.Dx()) * 0.75)
	height := int(float64(bounds.Dy()) * 0.75)

	img = imaging.Resize(img, width, height, o.ResampleFilter)
	return compress(img, o, depth+1)
}

func getSize(img image.Image, quality int) (int, error) {
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: quality}); err != nil {
		return -1, err
	}
	return buf.Len(), nil
}
