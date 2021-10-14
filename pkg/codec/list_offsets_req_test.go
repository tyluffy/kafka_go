package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeListOffsetsReqV5(t *testing.T) {
	bytes := testHex2Bytes(t, "00000008002f636f6e73756d65722d38646437623936622d366239342d346139622d623263632d3363623538393863396364662d31ffffffff00000000010006746573742d35000000010000000000000000fffffffffffffffe")
	listOffsetReq, err := DecodeListOffsetReq(bytes, 4)
	assert.Nil(t, err)
	assert.Equal(t, 8, listOffsetReq.CorrelationId)
	assert.Equal(t, "consumer-8dd7b96b-6b94-4a9b-b2cc-3cb5898c9cdf-1", listOffsetReq.ClientId)
	var expectedReplicaId int32 = -1
	assert.Equal(t, expectedReplicaId, listOffsetReq.ReplicaId)
	var expectedIsolationLevel uint8 = 0
	assert.Equal(t, expectedIsolationLevel, listOffsetReq.IsolationLevel)
	assert.Len(t, listOffsetReq.OffsetTopics, 1)
	offsetTopic := listOffsetReq.OffsetTopics[0]
	assert.Equal(t, "test-5", offsetTopic.Topic)
	offsetPartition := offsetTopic.ListOffsetPartitions[0]
	assert.Equal(t, 0, offsetPartition.PartitionId)
	assert.Equal(t, 0, offsetPartition.LeaderEpoch)
	var expectedPartitionTime int64 = -2
	assert.Equal(t, expectedPartitionTime, offsetPartition.Time)
}
