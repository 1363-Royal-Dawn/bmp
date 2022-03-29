package bmp

type Bitmap struct {
}

const BitmapFileHeaderSize = 14

type BitmapFileHeader struct {
	Type    uint16
	Size    uint32
	_       uint16
	_       uint16
	OffBits uint32
}

const BitmapInfoHeaderSize = 40

type BitmapInfoHeader struct {
	Size          uint32
	Width         int32
	Height        int32
	Planes        uint16
	BitCount      uint16 // 1 << BitCount (ColorTable if BitCount < 8)
	Compression   uint32
	SizeImage     uint32 // bytes to get
	XPelsPerMeter int32  // Picture Element
	YPelsPerMeter int32  // Picture Element
	ClrUsed       uint32
	ClrImportant  uint32
}

type RGBQuads []RGBQuad

type RGBQuad struct {
	Blue     byte
	Green    byte
	Red      byte
	Reserved byte
}
