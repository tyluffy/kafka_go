package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeIllegalOffsetCommitReq(t *testing.T) {
	bytes := make([]byte, 0)
	_, err := DecodeOffsetCommitReq(bytes, 0)
	assert.NotNil(t, err)
}

func TestDecodeOffsetCommitReqV8(t *testing.T) {
	bytes := testHex2Bytes(t, "0000000b002f636f6e73756d65722d38646437623936622d366239342d346139622d623263632d3363623538393863396364662d31002538646437623936622d366239342d346139622d623263632d3363623538393863396364660000000155636f6e73756d65722d38646437623936622d366239342d346139622d623263632d3363623538393863396364662d312d34333361636236612d653665632d343561612d623738642d366132343963666630376663000207746573742d35020000000000000000000000010000000001000000")
	offsetCommitReq, err := DecodeOffsetCommitReq(bytes, 8)
	assert.Nil(t, err)
	assert.Equal(t, 11, offsetCommitReq.CorrelationId)
	assert.Equal(t, 1, offsetCommitReq.GenerationId)
	assert.Equal(t, "consumer-8dd7b96b-6b94-4a9b-b2cc-3cb5898c9cdf-1", offsetCommitReq.ClientId)
	assert.Equal(t, "8dd7b96b-6b94-4a9b-b2cc-3cb5898c9cdf", offsetCommitReq.GroupId)
	assert.Equal(t, "consumer-8dd7b96b-6b94-4a9b-b2cc-3cb5898c9cdf-1-433acb6a-e6ec-45aa-b78d-6a249cff07fc", offsetCommitReq.MemberId)
	assert.Len(t, offsetCommitReq.OffsetCommitTopicReqList, 1)
	offsetTopic := offsetCommitReq.OffsetCommitTopicReqList[0]
	assert.Equal(t, "test-5", offsetTopic.Topic)
	offsetPartition := offsetTopic.OffsetPartitions[0]
	assert.Equal(t, 0, offsetPartition.PartitionId)
	var expectedOffset int64 = 1
	assert.Equal(t, expectedOffset, offsetPartition.Offset)
}
