package codec

func putBool(bytes []byte, idx int, x bool) int {
	if x {
		bytes[idx] = 1
	}
	return idx + 1
}
