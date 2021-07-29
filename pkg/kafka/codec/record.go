package codec

type Record struct {
	RecordAttributes  byte
	RelativeTimestamp int64
	RelativeOffset    int
	Key               []byte
	Value             string
	Headers           []byte
}

func (r *Record) BytesLength() int {
	result := 0
	result += LenRecordAttributes
	result += varint64Size(r.RelativeTimestamp)
	result += varintSize(len(r.Key))
	result += len(r.Key)
	result += varintSize(r.RelativeOffset)
	result += varintSize(len(r.Value))
	result += len(r.Value)
	result += varintSize(len(r.Headers))
	result += len(r.Headers)
	return result
}

func (r *Record) Bytes() []byte {
	bytes := make([]byte, r.BytesLength())
	idx := 0
	bytes[idx] = 0
	idx++
	idx = putVarint64(bytes, idx, r.RelativeTimestamp)
	idx = putVarint(bytes, idx, r.RelativeOffset)
	if r.Key == nil {
		idx = putVarint(bytes, idx, -1)
	} else {
		idx = putVarint(bytes, idx, len(r.Key))
		copy(bytes[idx:], r.Key)
		idx += len(r.Key)
	}
	idx = putVarint(bytes, idx, len(r.Value))
	copy(bytes[idx:], r.Value)
	idx += len(r.Value)
	idx = putVarint(bytes, idx, len(r.Headers))
	return bytes
}
