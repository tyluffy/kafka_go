package codec

func readBytes(bytes []byte, idx int) ([]byte, int) {
	length, idx := readInt(bytes, idx)
	return bytes[idx : idx+length], idx + length
}

func readCompactBytes(bytes []byte, idx int) ([]byte, int) {
	strLen, _ := readUVarint(bytes, idx)
	intLen := int(strLen)
	return bytes[idx+1 : idx+intLen], idx + intLen
}

func putBytes(bytes []byte, idx int, authBytes []byte) int {
	idx = putInt(bytes, idx, len(authBytes))
	copy(bytes[idx:], authBytes)
	return idx + len(authBytes)
}

func putCompactBytes(bytes []byte, idx int, compactBytes []byte) int {
	idx = putUVarint(bytes, idx, uint32(len(compactBytes)+1))
	copy(bytes[idx:], compactBytes)
	return idx + len(compactBytes)
}

func putVCompactBytes(bytes []byte, idx int, authBytes []byte) int {
	idx = putVarint(bytes, idx, len(authBytes))
	copy(bytes[idx:], authBytes)
	return idx + len(authBytes)
}

func BytesLen(bytes []byte) int {
	return 2 + len(bytes)
}

func CompactBytesLen(bytes []byte) int {
	return 1 + len(bytes)
}
