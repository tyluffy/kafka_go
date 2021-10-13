package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Marshal(t *testing.T) {
	r := &RecordBatch{}
	r.Offset = 0
	r.MessageSize = 64
	r.LeaderEpoch = 0
	r.MagicByte = 2
	r.Flags = 0
	r.LastOffsetDelta = 0
	r.FirstTimestamp = 1625962021853
	r.LastTimestamp = 1625962021853
	r.ProducerId = -1
	r.ProducerEpoch = -1
	r.BaseSequence = -1
	r.Records = make([]*Record, 1)
	record := &Record{}
	record.RecordAttributes = 0
	record.RelativeTimestamp = 0
	record.RelativeOffset = 0
	record.Key = nil
	record.Value = "ShootHzj"
	r.Records[0] = record
	assert.Equal(t, 76, r.BytesLength())
}
