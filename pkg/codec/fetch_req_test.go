package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeIllegalFetchReq(t *testing.T) {
	bytes := make([]byte, 0)
	_, err := DecodeFetchReq(bytes, 0)
	assert.NotNil(t, err)
}

func TestDecodeFetchReqV11(t *testing.T) {
	bytes := testHex2Bytes(t, "0000000a002f636f6e73756d65722d38646437623936622d366239342d346139622d623263632d3363623538393863396364662d31ffffffff000001f40000000103200000000000000000000000000000010006746573742d350000000100000000000000000000000000000000ffffffffffffffff00100000000000000000")
	fetchReq, err := DecodeFetchReq(bytes, 11)
	assert.Nil(t, err)
	assert.Equal(t, 10, fetchReq.CorrelationId)
	assert.Equal(t, "consumer-8dd7b96b-6b94-4a9b-b2cc-3cb5898c9cdf-1", fetchReq.ClientId)
	assert.Equal(t, 500, fetchReq.MaxWaitTime)
	assert.Equal(t, 1, fetchReq.MinBytes)
	assert.Equal(t, 52428800, fetchReq.MaxBytes)
	var exepectedIsolationLevel uint8 = 0
	assert.Equal(t, exepectedIsolationLevel, fetchReq.IsolationLevel)
	assert.Equal(t, 0, fetchReq.FetchSessionId)
	assert.Equal(t, 0, fetchReq.FetchSessionEpoch)
	assert.Len(t, fetchReq.FetchTopics, 1)
	fetchTopicReq := fetchReq.FetchTopics[0]
	assert.Equal(t, "test-5", fetchTopicReq.Topic)
	assert.Len(t, fetchTopicReq.FetchPartitions, 1)
	fetchPartitionReq := fetchTopicReq.FetchPartitions[0]
	assert.Equal(t, 0, fetchPartitionReq.PartitionId)
	assert.Equal(t, 0, fetchPartitionReq.CurrentLeaderEpoch)
	var expectedFetchOffset int64 = 0
	assert.Equal(t, expectedFetchOffset, fetchPartitionReq.FetchOffset)
	assert.Equal(t, 0, fetchPartitionReq.LastFetchedEpoch)
	var expectedLogOffset int64 = -1
	assert.Equal(t, expectedLogOffset, fetchPartitionReq.LogStartOffset)
	assert.Equal(t, 1048576, fetchPartitionReq.PartitionMaxBytes)
}
