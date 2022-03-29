package bmp

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func eightBpp(width int, height int, colorTable RGBQuads, pixelData []byte, stride int, filename string, output OutputType) {
	pixels := make(RGBQuads, 0)
	rows := len(pixelData) / stride
	for y := 0; y < rows; y++ {
		for x := 0; x < stride; x++ {
			idx := y*stride + x
			if idx%stride == width {
				break
			}
			pixel := pixelData[idx]
			pixels = append(pixels, colorTable[pixel])
		}
	}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := (height-1-y)*width + x
			element := pixels[idx]
			img.Set(x, y, color.RGBA{
				R: element.Red,
				G: element.Green,
				B: element.Blue,
				A: 255,
			})
		}
	}

	fil, err := os.Create("8bpp_test.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := fil.Close(); err != nil {
			log.Println(err)
		}
	}()

	if err := png.Encode(fil, img); err != nil {
		log.Fatal(err)
	}
}
