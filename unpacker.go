package binpacker

import (
	"encoding/binary"
	"io"
)

// Unpacker helps you unpack binary data from an io.Reader.
type Unpacker struct {
	reader io.Reader
	endian binary.ByteOrder
	err    error
}

// NewUnpacker returns a *Unpacker which hold an io.Reader.
func NewUnpacker(reader io.Reader) *Unpacker {
	return &Unpacker{
		reader: reader,
		endian: binary.LittleEndian,
	}
}

// Error returns an error if any errors exists
func (u *Unpacker) Error() error {
	return u.err
}

// ShiftByte fetch the first byte in io.Reader. Returns a byte and an error if
// exists.
func (u *Unpacker) ShiftByte() (byte, error) {
	buffer := make([]byte, 1)
	_, err := u.reader.Read(buffer)
	return buffer[0], err
}

// ReadByte fetch the first byte in io.Reader and set to b.
func (u *Unpacker) ReadByte(b *byte) *Unpacker {
	return u.errFilter(func() {
		*b, u.err = u.ShiftByte()
	})
}

// ShiftBytes fetch n bytes in io.Reader. Returns a byte array and an error if
// exists.
func (u *Unpacker) ShiftBytes(n uint64) ([]byte, error) {
	buffer := make([]byte, n)
	_, err := u.reader.Read(buffer)
	return buffer, err
}

// ReadBytes read n bytes and set to bytes.
func (u *Unpacker) ReadBytes(n uint64, bytes *[]byte) *Unpacker {
	return u.errFilter(func() {
		*bytes, u.err = u.ShiftBytes(n)
	})
}

// ShiftUint16 fetch 2 bytes in io.Reader and convert it to uint16.
func (u *Unpacker) ShiftUint16() (uint16, error) {
	buffer := make([]byte, 2)
	if _, err := u.reader.Read(buffer); err != nil {
		return 0, err
	}
	return u.endian.Uint16(buffer), nil
}

// ReadUint16 read 2 bytes, convert it to uint16 and set it to i.
func (u *Unpacker) ReadUint16(i *uint16) *Unpacker {
	return u.errFilter(func() {
		*i, u.err = u.ShiftUint16()
	})
}

// ShiftUint32 fetch 4 bytes in io.Reader and convert it to uint32.
func (u *Unpacker) ShiftUint32() (uint32, error) {
	buffer := make([]byte, 4)
	if _, err := u.reader.Read(buffer); err != nil {
		return 0, err
	}
	return u.endian.Uint32(buffer), nil
}

// ReadUint32 read 4 bytes, convert it to uint32 and set it to i.
func (u *Unpacker) ReadUint32(i *uint32) *Unpacker {
	return u.errFilter(func() {
		*i, u.err = u.ShiftUint32()
	})
}

// ShiftUint64 fetch 8 bytes in io.Reader and convert it to uint64.
func (u *Unpacker) ShiftUint64() (uint64, error) {
	buffer := make([]byte, 8)
	if _, err := u.reader.Read(buffer); err != nil {
		return 0, err
	}
	return u.endian.Uint64(buffer), nil
}

// ReadUint64 read 8 bytes, convert it to uint64 and set it to i.
func (u *Unpacker) ReadUint64(i *uint64) *Unpacker {
	return u.errFilter(func() {
		*i, u.err = u.ShiftUint64()
	})
}

// ShiftString fetch n bytes, convert it to string. Returns string and an error.
func (u *Unpacker) ShiftString(n uint64) (string, error) {
	buffer := make([]byte, n)
	if _, err := u.reader.Read(buffer); err != nil {
		return "", err
	}
	return string(buffer), nil
}

// ReadString read n bytes, convert it to string and set t to s.
func (u *Unpacker) ReadString(n uint64, s *string) *Unpacker {
	return u.errFilter(func() {
		*s, u.err = u.ShiftString(n)
	})
}

// StringWithUint16Perfix read 2 bytes as string length, then read N bytes,
// convert it to string and set it to s.
func (u *Unpacker) StringWithUint16Perfix(s *string) *Unpacker {
	return u.errFilter(func() {
		var n uint16
		n, u.err = u.ShiftUint16()
		u.ReadString(uint64(n), s)
	})
}

// StringWithUint32Perfix read 4 bytes as string length, then read N bytes,
// convert it to string and set it to s.
func (u *Unpacker) StringWithUint32Perfix(s *string) *Unpacker {
	return u.errFilter(func() {
		var n uint32
		n, u.err = u.ShiftUint32()
		u.ReadString(uint64(n), s)
	})
}

// StringWithUint64Perfix read 8 bytes as string length, then read N bytes,
// convert it to string and set it to s.
func (u *Unpacker) StringWithUint64Perfix(s *string) *Unpacker {
	return u.errFilter(func() {
		var n uint64
		n, u.err = u.ShiftUint64()
		u.ReadString(n, s)
	})
}

// BytesWithUint16Perfix read 2 bytes as bytes length, then read N bytes and set
// it to bytes.
func (u *Unpacker) BytesWithUint16Perfix(bytes *[]byte) *Unpacker {
	return u.errFilter(func() {
		var n uint16
		n, u.err = u.ShiftUint16()
		u.ReadBytes(uint64(n), bytes)
	})
}

// BytesWithUint32Perfix read 4 bytes as bytes length, then read N bytes and set
// it to bytes.
func (u *Unpacker) BytesWithUint32Perfix(bytes *[]byte) *Unpacker {
	return u.errFilter(func() {
		var n uint32
		n, u.err = u.ShiftUint32()
		u.ReadBytes(uint64(n), bytes)
	})
}

// BytesWithUint64Perfix read 8 bytes as bytes length, then read N bytes and set
// it to bytes.
func (u *Unpacker) BytesWithUint64Perfix(bytes *[]byte) *Unpacker {
	return u.errFilter(func() {
		var n uint64
		n, u.err = u.ShiftUint64()
		u.ReadBytes(n, bytes)
	})
}

func (u *Unpacker) errFilter(f func()) *Unpacker {
	if u.err == nil {
		f()
	}
	return u
}
