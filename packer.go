package binpacker

import (
	"encoding/binary"
	"io"
)

type Packer struct {
	writer io.Writer
	endian binary.ByteOrder
	err    error
}

func NewPacker(writer io.Writer) *Packer {
	return &Packer{
		writer: writer,
		endian: binary.LittleEndian,
	}
}

func (p *Packer) Error() error {
	return p.err
}

func (p *Packer) PushByte(b byte) *Packer {
	return p.errFilter(func() {
		_, p.err = p.writer.Write([]byte{b})
	})
}

func (p *Packer) PushBytes(bytes []byte) *Packer {
	return p.errFilter(func() {
		_, p.err = p.writer.Write(bytes)
	})
}

func (p *Packer) PushUint16(i uint16) *Packer {
	return p.errFilter(func() {
		buffer := make([]byte, 2)
		p.endian.PutUint16(buffer, i)
		_, p.err = p.writer.Write(buffer)
	})
}

func (p *Packer) PushUint32(i uint32) *Packer {
	return p.errFilter(func() {
		buffer := make([]byte, 4)
		p.endian.PutUint32(buffer, i)
		_, p.err = p.writer.Write(buffer)
	})
}

func (p *Packer) PushUint64(i uint64) *Packer {
	return p.errFilter(func() {
		buffer := make([]byte, 8)
		p.endian.PutUint64(buffer, i)
		_, p.err = p.writer.Write(buffer)
	})
}

func (p *Packer) PushString(s string) *Packer {
	return p.errFilter(func() {
		_, p.err = p.writer.Write([]byte(s))
	})
}

func (p *Packer) errFilter(f func()) *Packer {
	if p.err == nil {
		f()
	}
	return p
}
