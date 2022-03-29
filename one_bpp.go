package bmp

import (
	"image"
	"image/color"
	"log"
)

func cursor(size int) func(x, y int) int {
	return func(x, y int) int {
		return y*size + x
	}
}

func calculateOneBppColor(element byte, row RGBQuads, colorTable RGBQuads) RGBQuads {
	for bit := 7; bit >= 0; bit-- {
		bitValue := element & (1 << bit)
		if bitValue > 0 {
			row = append(row, colorTable[1])
		} else {
			row = append(row, colorTable[0])
		}
	}
	return row
}
func fatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type OutputType int8

const (
	OutputJPEG OutputType = iota
	OutputPNG
)

func oneBpp(width int, height int, colorTable RGBQuads, pixelData []byte, stride int) image.Image {
	index := cursor(stride)
	pixels := make([]RGBQuads, 0)
	rows := len(pixelData) / stride
	for y := 0; y < rows; y++ {
		row := make(RGBQuads, 0)
		for x := 0; x < stride; x++ {
			idx := index(x, y)
			element := pixelData[idx]
			row = calculateOneBppColor(element, row, colorTable)
		}
		pixels = append(pixels, row)
	}
	img := image.NewGray(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			element := pixels[height-y-1][x]
			img.Set(x, y, color.Gray{Y: element.Green})
		}
	}
	return img
}
