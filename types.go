package bmp

import (
	"encoding/binary"
	"io"
)

type IconDir struct {
	Reserved, Type, Count uint16
}

func NewIconDir(count uint16) *IconDir {
	return &IconDir{Reserved: 0, Type: 1, Count: count}
}
func (ih *IconDir) Save(wr io.Writer) error {
	return binary.Write(wr, binary.LittleEndian, ih)
}

type IconDirEntry struct {
	Width, Height uint8
	ColorCount    uint8
	Reserved      uint8
	Planes        uint16
	BitCount      uint16
	BytesInRes    uint32
	ImageOffset   uint32
}

func NewIconDirEntry(width, height int, colorCount uint8, planes uint16, bitCount uint16, bytesInRes uint32, imageOffset uint32) *IconDirEntry {
	var w, h uint8
	if width == 256 {
		w = 0
	} else {
		w = uint8(width)
	}
	if height == 256 {
		h = 0
	} else {
		h = uint8(height)
	}
	return &IconDirEntry{
		Width: w, Height: h, ColorCount: colorCount, Reserved: 0,
		Planes:      planes,
		BitCount:    bitCount,
		BytesInRes:  bytesInRes,
		ImageOffset: imageOffset,
	}
}
func (ie IconDirEntry) Save(wr io.Writer) error {
	return binary.Write(wr, binary.LittleEndian, ie)
}

func LoadIconDirEntry(r io.Reader) (*IconDirEntry, error) {
	var ie IconDirEntry
	if err := binary.Read(r, binary.LittleEndian, &ie); err != nil {
		return nil, err
	}
	return &ie, nil
}

type IconImage struct {
	ICHeader BitmapInfoHeader
	ICColors []RGBQuad
	ICXor    []byte
	ICAnd    []byte
}

func LoadBitmapInfoHeader(r io.Reader) (*BitmapInfoHeader, error) {
	var bmih BitmapInfoHeader
	if err := binary.Read(r, binary.LittleEndian, &bmih); err != nil {
		return nil, err
	}
	return &bmih, nil
}
