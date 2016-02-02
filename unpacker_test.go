package binpacker

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShiftByte(t *testing.T) {
	buf := new(bytes.Buffer)
	p := NewPacker(buf)
	u := NewUnpacker(buf)
	p.PushByte(0x01)
	b, err := u.ShiftByte()
	assert.Equal(t, err, nil, "Has error.")
	assert.Equal(t, b, byte(0x01), "byte error.")
}

func TestShiftBytes(t *testing.T) {
	buf := new(bytes.Buffer)
	p := NewPacker(buf)
	u := NewUnpacker(buf)
	p.PushBytes([]byte{0x01, 0x02})
	bs, err := u.ShiftBytes(2)
	assert.Equal(t, err, nil, "Has error.")
	assert.Equal(t, bs, []byte{0x01, 0x02}, "byte error.")
}

func TestShiftUint16(t *testing.T) {
	buf := new(bytes.Buffer)
	p := NewPacker(buf)
	u := NewUnpacker(buf)
	p.PushUint16(1)
	i, err := u.ShiftUint16()
	assert.Equal(t, err, nil, "Has error.")
	assert.Equal(t, i, uint16(1), "uint16 error.")
}

func TestShiftUint32(t *testing.T) {
	buf := new(bytes.Buffer)
	p := NewPacker(buf)
	u := NewUnpacker(buf)
	p.PushUint32(1)
	i, err := u.ShiftUint32()
	assert.Equal(t, err, nil, "Has error.")
	assert.Equal(t, i, uint32(1), "uint32 error.")
}

func TestShiftUint64(t *testing.T) {
	buf := new(bytes.Buffer)
	p := NewPacker(buf)
	u := NewUnpacker(buf)
	p.PushUint64(1)
	i, err := u.ShiftUint64()
	assert.Equal(t, err, nil, "Has error.")
	assert.Equal(t, i, uint64(1), "uint64 error.")
}

func TestShiftString(t *testing.T) {
	buf := new(bytes.Buffer)
	p := NewPacker(buf)
	u := NewUnpacker(buf)
	p.PushString("Hi")
	s, err := u.ShiftString(2)
	assert.Equal(t, err, nil, "Has error.")
	assert.Equal(t, s, "Hi", "string error.")
}

func TestRead(t *testing.T) {
	buf := new(bytes.Buffer)
	p := NewPacker(buf)
	u := NewUnpacker(buf)
	p.PushByte(0x01)
	p.PushBytes([]byte("Hi"))
	p.PushUint16(1)
	p.PushUint32(1)
	p.PushUint64(1)
	p.PushString("Hi")
	var b byte
	var bs []byte
	var i16 uint16
	var i32 uint32
	var i64 uint64
	var s string
	u.ReadByte(&b).
		ReadBytes(2, &bs).
		ReadUint16(&i16).
		ReadUint32(&i32).
		ReadUint64(&i64).
		ReadString(2, &s)
	assert.Equal(t, u.err, nil, "Has Error.")
	assert.Equal(t, u.Error(), nil, "Has Error.")
	assert.Equal(t, b, byte(0x01), "byte error.")
	assert.Equal(t, bs, []byte("Hi"), "bytes error.")
	assert.Equal(t, i16, uint16(1), "uint16 error.")
	assert.Equal(t, i32, uint32(1), "uint32 error.")
	assert.Equal(t, i64, uint64(1), "uint64 error.")
	assert.Equal(t, s, "Hi", "string error.")
}

func TestReadWithPerfix(t *testing.T) {
	buf := new(bytes.Buffer)
	p := NewPacker(buf)
	u := NewUnpacker(buf)
	var bs []byte
	var s string

	p.PushUint16(2)
	p.PushBytes([]byte("Hi"))
	u.BytesWithUint16Perfix(&bs)
	assert.Equal(t, bs, []byte("Hi"), "Bytes with perfix error.")
	p.PushUint16(2)
	p.PushString("Hi")
	u.StringWithUint16Perfix(&s)
	assert.Equal(t, s, "Hi", "String with perfix error.")

	p.PushUint32(2)
	p.PushBytes([]byte("Hi"))
	u.BytesWithUint32Perfix(&bs)
	assert.Equal(t, bs, []byte("Hi"), "Bytes with perfix error.")
	p.PushUint32(2)
	p.PushString("Hi")
	u.StringWithUint32Perfix(&s)
	assert.Equal(t, s, "Hi", "String with perfix error.")

	p.PushUint64(2)
	p.PushBytes([]byte("Hi"))
	u.BytesWithUint64Perfix(&bs)
	assert.Equal(t, bs, []byte("Hi"), "Bytes with perfix error.")
	p.PushUint64(2)
	p.PushString("Hi")
	u.StringWithUint64Perfix(&s)
	assert.Equal(t, s, "Hi", "String with perfix error.")
}
