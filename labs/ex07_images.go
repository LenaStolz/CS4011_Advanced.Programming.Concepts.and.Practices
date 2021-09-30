package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
) 

type Image struct{
	width int
	height int
}

func (bild Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (bild Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, bild.width, bild.height)
}

func (bild Image) At(x, y int) color.Color {
	v := uint8(2 * x ^  - 2 * x - (y ^ x))
	return color.RGBA{ v, v, v ^ v ^ -v, 255}
}

func main() {
	m := Image{120, 240}
	pic.ShowImage(m)
}
