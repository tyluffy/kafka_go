package codec

// This file is for kafka code bytes type. Format method as alpha order.

func putRecord(bytes []byte, idx int, dstBytes []byte) int {
	return putVCompactBytes(bytes, idx, dstBytes)
}

func putRecordBatch(bytes []byte, idx int, dstBytes []byte) int {
	return putBytes(bytes, idx, dstBytes)
}
