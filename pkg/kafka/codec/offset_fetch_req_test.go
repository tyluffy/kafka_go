package codec

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeOffsetFetchReq(t *testing.T) {
	bytes, err := hex.DecodeString("0000000b002f636f6e73756d65722d61303332616233632d303831382d343937352d626439332d3735613431323030656162342d31002561303332616233632d303831382d343937352d626439332d373561343132303065616234020a746573742d7361736c02000000000000")
	assert.Nil(t, err)
	fetchReq, err := DecodeOffsetFetchReq(bytes, 6)
	assert.Nil(t, err)
	assert.Equal(t, 11, fetchReq.CorrelationId)
	assert.Equal(t, "consumer-a032ab3c-0818-4975-bd93-75a41200eab4-1", fetchReq.ClientId)
	assert.Equal(t, "a032ab3c-0818-4975-bd93-75a41200eab4", fetchReq.GroupId)
	assert.Len(t, fetchReq.TopicReqList, 1)
	fetchTopicReq := fetchReq.TopicReqList[0]
	assert.Equal(t, "test-sasl", fetchTopicReq.Topic)
	assert.Len(t, fetchTopicReq.PartitionReqList, 1)
	fetchPartitionReq := fetchTopicReq.PartitionReqList[0]
	assert.Equal(t, 0, fetchPartitionReq.PartitionId)
	assert.False(t, fetchReq.RequireStableOffset)
}
