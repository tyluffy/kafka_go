package codec

func readIsolationLevel(bytes []byte, idx int) (byte, int) {
	return bytes[idx], idx + 1
}

func readCoordinatorType(bytes []byte, idx int) (byte, int) {
	return bytes[idx], idx + 1
}
