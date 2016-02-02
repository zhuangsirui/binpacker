package binpacker

import (
	"encoding/binary"
	"io"
)

// Packer is a binary packer helps you pack data into an io.Writer.
type Packer struct {
	writer io.Writer
	endian binary.ByteOrder
	err    error
}

// NewPacker returns a *Packer hold an io.Writer.
func NewPacker(writer io.Writer) *Packer {
	return &Packer{
		writer: writer,
		endian: binary.LittleEndian,
	}
}

// Error returns an error if any errors exists
func (p *Packer) Error() error {
	return p.err
}

// PushByte write a single byte into writer.
func (p *Packer) PushByte(b byte) *Packer {
	return p.errFilter(func() {
		_, p.err = p.writer.Write([]byte{b})
	})
}

// PushBytes write a bytes array into writer.
func (p *Packer) PushBytes(bytes []byte) *Packer {
	return p.errFilter(func() {
		_, p.err = p.writer.Write(bytes)
	})
}

// PushUint16 write a uint16 into writer.
func (p *Packer) PushUint16(i uint16) *Packer {
	return p.errFilter(func() {
		buffer := make([]byte, 2)
		p.endian.PutUint16(buffer, i)
		_, p.err = p.writer.Write(buffer)
	})
}

// PushUint32 write a uint32 into writer.
func (p *Packer) PushUint32(i uint32) *Packer {
	return p.errFilter(func() {
		buffer := make([]byte, 4)
		p.endian.PutUint32(buffer, i)
		_, p.err = p.writer.Write(buffer)
	})
}

// PushUint64 write a uint64 into writer.
func (p *Packer) PushUint64(i uint64) *Packer {
	return p.errFilter(func() {
		buffer := make([]byte, 8)
		p.endian.PutUint64(buffer, i)
		_, p.err = p.writer.Write(buffer)
	})
}

// PushString write a string into writer.
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
