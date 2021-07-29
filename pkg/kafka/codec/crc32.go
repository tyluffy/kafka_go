package codec

import (
	"encoding/binary"
	"hash/crc32"
)

var table = crc32.MakeTable(crc32.Castagnoli)

func putCrc32(bytes []byte, data []byte) {
	checksum := crc32.Checksum(data, table)
	binary.BigEndian.PutUint32(bytes, checksum)
}
