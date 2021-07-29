package codec

func putSaslAuthBytes(bytes []byte, idx int, authBytes []byte) int {
	return putBytes(bytes, idx, authBytes)
}

func putSaslAuthBytesCompact(bytes []byte, idx int, authBytes []byte) int {
	return putCompactBytes(bytes, idx, authBytes)
}