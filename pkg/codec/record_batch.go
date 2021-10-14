package codec

type RecordBatch struct {
	Offset      int64
	MessageSize int
	LeaderEpoch int
	MagicByte   byte
	// 8位flag字节
	Flags uint16

	LastOffsetDelta int
	FirstTimestamp  int64
	LastTimestamp   int64
	ProducerId      int64
	ProducerEpoch   int16
	BaseSequence    int
	Records         []*Record
}

func (r *RecordBatch) BytesLength() int {
	result := LenOffset + LenMessageSize + LenLeaderEpoch + LenMagicByte + LenCrc32
	// record batch flags
	result += 2
	result += LenOffsetDelta
	result += LenTime
	result += LenTime
	result += LenProducerId
	result += LenProducerEpoch
	result += LenBaseSequence
	result += LenArray
	for _, record := range r.Records {
		// record size
		bytesLength := record.BytesLength()
		result += varintSize(bytesLength)
		result += bytesLength
	}
	return result
}

func (r *RecordBatch) Bytes() []byte {
	bytes := make([]byte, r.BytesLength())
	idx := 0
	idx = putOffset(bytes, idx, r.Offset)
	idx = putMessageSize(bytes, idx, r.MessageSize)
	idx = putLeaderEpoch(bytes, idx, r.LeaderEpoch)
	bytes[idx] = r.MagicByte
	idx++
	crc32Start := idx
	// crc32后面计算
	idx += 4
	idx = putUInt16(bytes, idx, r.Flags)
	idx = putLastOffsetDelta(bytes, idx, r.LastOffsetDelta)
	idx = putTime(bytes, idx, r.FirstTimestamp)
	idx = putTime(bytes, idx, r.LastTimestamp)
	idx = putProducerId(bytes, idx, r.ProducerId)
	idx = putProducerEpoch(bytes, idx, r.ProducerEpoch)
	idx = putBaseSequence(bytes, idx, r.BaseSequence)
	idx = putArrayLen(bytes, idx, len(r.Records))
	for _, record := range r.Records {
		idx = putRecord(bytes, idx, record.Bytes())
	}
	putCrc32(bytes[crc32Start:], bytes[crc32Start+4:])
	return bytes
}
