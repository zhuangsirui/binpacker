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

func TestShiftInt16(t *testing.T) {
	buf := new(bytes.Buffer)
	p := NewPacker(buf)
	u := NewUnpacker(buf)
	p.PushInt16(-1)
	i, err := u.ShiftInt16()
	assert.Equal(t, err, nil, "Has error.")
	assert.Equal(t, i, int16(-1), "uint16 error.")
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

func TestShiftInt32(t *testing.T) {
	buf := new(bytes.Buffer)
	p := NewPacker(buf)
	u := NewUnpacker(buf)
	p.PushInt32(-1)
	i, err := u.ShiftInt32()
	assert.Equal(t, err, nil, "Has error.")
	assert.Equal(t, i, int32(-1), "int32 error.")
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

func TestShiftInt64(t *testing.T) {
	buf := new(bytes.Buffer)
	p := NewPacker(buf)
	u := NewUnpacker(buf)
	p.PushInt64(-1)
	i, err := u.ShiftInt64()
	assert.Equal(t, err, nil, "Has error.")
	assert.Equal(t, i, int64(-1), "int64 error.")
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
	p.PushInt16(-1)
	p.PushUint32(1)
	p.PushInt32(-1)
	p.PushUint64(1)
	p.PushInt64(-1)
	p.PushString("Hi")
	var b byte
	var bs []byte
	var ui16 uint16
	var i16 int16
	var ui32 uint32
	var i32 int32
	var ui64 uint64
	var i64 int64
	var s string
	u.FetchByte(&b).
		FetchBytes(2, &bs).
		FetchUint16(&ui16).
		FetchInt16(&i16).
		FetchUint32(&ui32).
		FetchInt32(&i32).
		FetchUint64(&ui64).
		FetchInt64(&i64).
		FetchString(2, &s)
	assert.Equal(t, u.err, nil, "Has Error.")
	assert.Equal(t, u.Error(), nil, "Has Error.")
	assert.Equal(t, b, byte(0x01), "byte error.")
	assert.Equal(t, bs, []byte("Hi"), "bytes error.")
	assert.Equal(t, ui16, uint16(1), "uint16 error.")
	assert.Equal(t, i16, int16(-1), "int16 error.")
	assert.Equal(t, ui32, uint32(1), "uint32 error.")
	assert.Equal(t, i32, int32(-1), "int32 error.")
	assert.Equal(t, ui64, uint64(1), "uint64 error.")
	assert.Equal(t, i64, int64(-1), "int64 error.")
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
