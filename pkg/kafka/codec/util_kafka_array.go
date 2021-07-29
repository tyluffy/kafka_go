package codec

func putArrayLen(bytes []byte, idx int, length int) int {
	return putInt(bytes, idx, length)
}

func readArrayLen(bytes []byte, idx int) (int, int) {
	return readInt(bytes, idx)
}

func putCompactArrayLen(bytes []byte, idx int, length int) int {
	return putUVarint(bytes, idx, uint32(length+1))
}

func readCompactArrayLen(bytes []byte, idx int) (int, int) {
	uVarint, i := readUVarint(bytes, idx)
	return int(uVarint - 1), i
}
