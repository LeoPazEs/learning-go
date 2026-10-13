package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Width  int
	Height int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	v := uint8(((x + y) / 2) - (x * y) + (x ^ y))
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{200, 100}

	pic.ShowImage(m)
}
