package tiger

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

// uses unsafe for fast conversion from []byte to []uint64
// Overall benchmark time 104.62 MB/s
func (t *tiger) readFrom(buf []byte) []byte {
	for len(buf) >= BlockSize {
		// Pun the byte slice into an uint64 slice
		// We don't use its len directly as it'll be incorrect,
		// but we already know it has enough for the iteration.
		ptr := unsafe.Pointer(&buf)
		buf64 := []uint64(*(*[]uint64)(ptr))

		// Read the uint64 slice into an array
		x := [8]uint64{}
		for i := range x {
			x[i] = buf64[i]
		}
		buf = buf[BlockSize:]
		t.tigerBlock(x)
	}
	return buf
}

// encoding/binary version of the above
// Overall benchmark time 38.09 MB/s
func (t *tiger) readFrom_binary(b []byte) []byte {
	buf := bytes.NewBuffer(b)
	for buf.Len() >= BlockSize {
		x := [8]uint64{}

		// NOTE: Ideally, this would be bytes.NativeEndian, but we don't have that.
		// All supported platforms are LE (...and ARM) so this works out fine.
		binary.Read(buf, binary.LittleEndian, x[:])
		t.tigerBlock(x)
	}
	return buf.Bytes()
}
