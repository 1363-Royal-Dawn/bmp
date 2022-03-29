package bmp

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func high4(n int) int {
	return (n & 0xf0) >> 4
}

func low4(n int) int {
	return n & 0x0f
}

func fourBpp(width, height int, colorTable RGBQuads, pixelData []byte, stride int, filename string, output OutputType) {
	index := cursor(stride)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	pixels := make([]RGBQuads, 0)
	fmt.Printf("len %d\n", len(pixelData))
	for y := 0; y < (len(pixelData) / stride); y++ {
		row := make(RGBQuads, 0)
		for x := 0; x < stride; x++ {
			idx := index(x, y)
			element := pixelData[idx]

			i1 := high4(int(element))
			row = append(row, colorTable[i1])
			i2 := low4(int(element))
			row = append(row, colorTable[i2])

		}
		fmt.Printf("len %d\n", len(row))
		pixels = append(pixels, row)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixelColor := pixels[height-y-1][x]
			img.Set(x, y, color.RGBA{
				R: pixelColor.Red,
				G: pixelColor.Green,
				B: pixelColor.Blue,
				A: 255,
			})
		}
	}
	fil, err := os.Create("test_subject" + filepath.Base(filename) + ".png")
	if err != nil {
		log.Println(err)
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
