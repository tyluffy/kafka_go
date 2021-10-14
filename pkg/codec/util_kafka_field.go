package codec

func readSaslAuthBytes(bytes []byte, idx int) (string, string) {
	usernameIdx := idx
	totalLength := idx + len(bytes)
	for i := idx + 1; i < totalLength; i++ {
		usernameIdx = i
		if bytes[i] == 0 {
			break
		}
	}
	return string(bytes[idx+1 : usernameIdx]), string(bytes[usernameIdx+1 : totalLength])
}

func putTaggedField(bytes []byte, idx int) int {
	bytes[idx] = 0
	return idx + 1
}

func readTaggedField(bytes []byte, idx int) int {
	if bytes[idx] != 0 {
		panic("not supported tagged field yet")
	}
	return idx + 1
}
