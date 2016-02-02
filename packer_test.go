package binpacker

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushByte(t *testing.T) {
	b := new(bytes.Buffer)
	p := NewPacker(b)
	p.PushByte(0x01)
	assert.Equal(t, p.Error(), nil, "Has error.")
	assert.Equal(t, b.Bytes(), []byte{1}, "byte error.")
}

func TestPushBytes(t *testing.T) {
	b := new(bytes.Buffer)
	p := NewPacker(b)
	p.PushBytes([]byte{0x01, 0x002})
	assert.Equal(t, p.Error(), nil, "Has error.")
	assert.Equal(t, b.Bytes(), []byte{0x01, 0x02}, "bytes error.")
}

func TestPushUint16(t *testing.T) {
	b := new(bytes.Buffer)
	p := NewPacker(b)
	p.PushUint16(1)
	assert.Equal(t, p.Error(), nil, "Has error.")
	assert.Equal(t, b.Bytes(), []byte{1, 0}, "uint16 error.")
}

func TestPushUint32(t *testing.T) {
	b := new(bytes.Buffer)
	p := NewPacker(b)
	p.PushUint32(1)
	assert.Equal(t, p.Error(), nil, "Has error.")
	assert.Equal(t, b.Bytes(), []byte{1, 0, 0, 0}, "uint32 error.")
}

func TestPushUint64(t *testing.T) {
	b := new(bytes.Buffer)
	p := NewPacker(b)
	p.PushUint64(1)
	assert.Equal(t, p.Error(), nil, "Has error.")
	assert.Equal(t, b.Bytes(), []byte{1, 0, 0, 0, 0, 0, 0, 0}, "uint64 error.")
}

func TestPushString(t *testing.T) {
	b := new(bytes.Buffer)
	p := NewPacker(b)
	p.PushString("Hi")
	assert.Equal(t, p.Error(), nil, "Has error.")
	assert.Equal(t, b.Bytes(), []byte{'H', 'i'}, "string error.")
}

func TestCombinedPush(t *testing.T) {
	b := new(bytes.Buffer)
	p := NewPacker(b)
	p.PushUint16(1).PushString("Hi")
	assert.Equal(t, p.Error(), nil, "Has error.")
	assert.Equal(t, b.Bytes(), []byte{1, 0, 'H', 'i'}, "combine push error.")
}
