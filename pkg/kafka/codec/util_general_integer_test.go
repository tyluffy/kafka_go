package codec

import (
	"testing"
)

func Test_readInt64(t *testing.T) {
	timestampSlice := []byte{0x00, 0x00, 0x01, 0x7a, 0x92, 0xe3, 0x83, 0xdd}
	res, _ := readInt64(timestampSlice, 0)
	if res != 1625962021853 {
		t.Errorf("read failed")
	}
}
