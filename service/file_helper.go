package service

import (
	"errors"
	"image"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"

	xdraw "golang.org/x/image/draw"
	_ "golang.org/x/image/webp"
)

func makeSquareAvatar(src io.Reader, dest io.Writer, resolution int) error {
	img, _, err := image.Decode(src)
	if err != nil {
		return err
	}
	var (
		bounds   = img.Bounds()
		width    = bounds.Dx()
		height   = bounds.Dy()
		cropArea image.Rectangle
	)
	if width > height {
		// Landscape image
		offset := (width - height) / 2
		cropArea = image.Rect(
			bounds.Min.X+offset,
			bounds.Min.Y,
			bounds.Min.X+offset+height,
			bounds.Max.Y,
		)
	} else {
		offset := (height - width) / 2
		cropArea = image.Rect(
			bounds.Min.X,
			bounds.Min.Y+offset,
			bounds.Max.X,
			bounds.Min.Y+offset+width,
		)
	}
	sub, ok := img.(SubImager)
	if !ok {
		return errors.New("this file does not implement SubImager interface.")
	}
	croppedImg := sub.SubImage(cropArea)
	canvas := image.NewRGBA(image.Rect(0, 0, resolution, resolution))

	xdraw.BiLinear.Scale(canvas, canvas.Bounds(), croppedImg, croppedImg.Bounds(), draw.Src, nil)

	return jpeg.Encode(dest, canvas, &jpeg.Options{Quality: jpeg.DefaultQuality})
}
