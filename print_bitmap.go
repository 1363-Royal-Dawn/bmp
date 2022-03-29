package bmp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type BitmapConfiguration struct {
	Width      int
	Height     int
	ColorTable RGBQuads
	Pixels     []byte
	Stride     int
}

func setupBitmap(r io.Reader) (*BitmapConfiguration, error) {
	var (
		bfh BitmapFileHeader
		bih BitmapInfoHeader
	)

	bfile, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	bitmapFileHeader := bfile[0:BitmapFileHeaderSize]
	if err := binary.Read(bytes.NewReader(bitmapFileHeader), binary.LittleEndian, &bfh); err != nil {
		return nil, err
	}

	bitmapInfoHeader := bfile[BitmapFileHeaderSize : BitmapFileHeaderSize+BitmapInfoHeaderSize]
	if err := binary.Read(bytes.NewReader(bitmapInfoHeader), binary.LittleEndian, &bih); err != nil {
		return nil, err
	}

	stride := int(bih.BitCount) * int(bih.Width)
	stride = ((stride + 31) / 32) * 4
	colorTableSize := 1 << bih.BitCount
	colorTable := make(RGBQuads, colorTableSize)
	colorTableOffset := BitmapFileHeaderSize + BitmapInfoHeaderSize
	for i := 0; i < colorTableSize; i++ {
		b := colorTableOffset + i
		colorTableOffset++
		g := colorTableOffset + i
		colorTableOffset++
		r := colorTableOffset + i
		colorTableOffset++
		a := colorTableOffset + i
		colorTable[i] = RGBQuad{bfile[b], bfile[g], bfile[r], bfile[a]}
	}
	fmt.Printf("Color(%03d): %v\n", len(colorTable), colorTable)
	pixelData := bfile[bfh.OffBits:]

	return &BitmapConfiguration{
		Width:      int(bih.Width),
		Height:     int(bih.Height),
		ColorTable: colorTable,
		Pixels:     pixelData,
		Stride:     stride,
	}, nil

}
func printBitmap(filename string, output OutputType) {
	bfile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("fisk", err)
	}
	fmt.Printf("Opening %s\n", filename)

	var (
		bfh BitmapFileHeader
		bih BitmapInfoHeader
	)
	bitmapFileHeader := bfile[0:BitmapFileHeaderSize]
	if err := binary.Read(bytes.NewReader(bitmapFileHeader), binary.LittleEndian, &bfh); err != nil {
		log.Fatal(err)
	}

	bitmapInfoHeader := bfile[BitmapFileHeaderSize : BitmapFileHeaderSize+BitmapInfoHeaderSize]
	if err := binary.Read(bytes.NewReader(bitmapInfoHeader), binary.LittleEndian, &bih); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("bfh %+v\n", bfh)
	fmt.Printf("bih %+v\n", bih)
	stride := int(bih.BitCount) * int(bih.Width)
	stride = ((stride + 31) / 32) * 4
	fmt.Printf("stride (bytes): %v\n", stride)
	colorTableSize := 1 << bih.BitCount
	colorTable := make(RGBQuads, colorTableSize)
	colorTableOffset := BitmapFileHeaderSize + BitmapInfoHeaderSize
	for i := 0; i < colorTableSize; i++ {
		b := colorTableOffset + i
		colorTableOffset++
		g := colorTableOffset + i
		colorTableOffset++
		r := colorTableOffset + i
		colorTableOffset++
		a := colorTableOffset + i
		colorTable[i] = RGBQuad{bfile[b], bfile[g], bfile[r], bfile[a]}
	}
	fmt.Printf("Color(%03d): %v\n", len(colorTable), colorTable)
	pixelData := bfile[bfh.OffBits:]

	switch bih.BitCount {
	case 1:
		oneBpp(int(bih.Width), int(bih.Height), colorTable, pixelData, stride)
	case 4:
		fourBpp(int(bih.Width), int(bih.Height), colorTable, pixelData, stride, filename, output)
	case 8:
		eightBpp(int(bih.Width), int(bih.Height), colorTable, pixelData, stride, filename, output)
	default:
		fmt.Printf("reading BMP file with %d bitcount\n", 1<<bih.BitCount)
	}

}
