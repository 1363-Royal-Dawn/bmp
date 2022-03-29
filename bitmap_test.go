package bmp

import (
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestOneBpp(t *testing.T) {

	output := OutputPNG
	inputFile, err := os.Open("testdata/monochrome_1000x1000.bmp")
	if err != nil {
		t.Errorf("%v", err)
	}
	defer inputFile.Close()

	configuration, err := setupBitmap(inputFile)
	if err != nil {
		t.Errorf("%v", err)
	}

	fil, err := os.Create("output_1bpp_monochrome_1000x1000.png")
	if err != nil {
		t.Errorf("%v", err)
	}
	defer fil.Close()

	img := oneBpp(configuration.Width, configuration.Height, configuration.ColorTable, configuration.Pixels, configuration.Stride)

	switch output {
	case OutputJPEG:
		if err := jpeg.Encode(fil, img, &jpeg.Options{
			Quality: jpeg.DefaultQuality,
		}); err != nil {
			log.Fatal(err)
		}
	case OutputPNG:
		if err := png.Encode(fil, img); err != nil {
			log.Fatal(err)
		}
	}

}
