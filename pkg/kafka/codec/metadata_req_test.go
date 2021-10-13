package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeMetadataV9(t *testing.T) {
	bytes := testHex2Bytes(t, "00000002002f636f6e73756d65722d37336664633964612d306439322d346537622d613761372d6563323636663637633137312d3100022537363465646565332d303037652d343865302d623966392d6466376637313366663730370001000000")
	metadataTopicReq, err := DecodeMetadataTopicReq(bytes, 9)
	assert.Nil(t, err)
	assert.Equal(t, 2, metadataTopicReq.CorrelationId)
	assert.Equal(t, "consumer-73fdc9da-0d92-4e7b-a7a7-ec266f67c171-1", metadataTopicReq.ClientId)
	assert.Len(t, metadataTopicReq.Topics, 1)
	topicReq := metadataTopicReq.Topics[0]
	assert.Equal(t, "764edee3-007e-48e0-b9f9-df7f713ff707", topicReq.Topic)
	assert.False(t, metadataTopicReq.AllowAutoTopicCreation)
	assert.False(t, metadataTopicReq.IncludeClusterAuthorizedOperations)
	assert.False(t, metadataTopicReq.IncludeTopicAuthorizedOperations)
}
