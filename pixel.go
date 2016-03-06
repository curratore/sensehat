package sensehat

import (
	"bytes"
	"encoding/binary"
)

type Pixel struct {
	// Raw
	R int
	G int
	B int
}

func (p *Pixel) ToArray() []int {
	return []int{p.R, p.G, p.B}
}

func (p *Pixel) Pack() []byte {
	r := (p.R >> 3) & 0x1f
	g := (p.G >> 2) & 0x3f
	b := (p.B >> 3) & 0x1f

	b16 := (r << 11) + (g << 5) + b

	buf := new(bytes.Buffer)

	_ = binary.Write(buf, binary.LittleEndian, uint16(b16))

	return buf.Bytes()
}

func (p *Pixel) Unpack(packed []byte) *Pixel {
	var result uint16

	buf := bytes.NewReader(packed)
	_ = binary.Read(buf, binary.LittleEndian, &result)

	r := (result & 0xF800) >> 11
	g := (result & 0x7E0) >> 5
	b := (result & 0x1F)

	p.R = int(r << 3)
	p.G = int(g << 2)
	p.B = int(b << 3)

	return p
}

func (p *Pixel) Valid() bool {

	// raise ValueError('Pixel elements must be between 0 and 255')
	for _, item := range p.ToArray() {
		if item > 255 || item < 0 {
			return false
		}
	}

	return true
}
