package codec

import (
	"encoding/hex"
	"testing"
)

func TestDecodeOffsetFetchReq(t *testing.T) {
	bytes, _ := hex.DecodeString("0000000b002f636f6e73756d65722d61303332616233632d303831382d343937352d626439332d3735613431323030656162342d31002561303332616233632d303831382d343937352d626439332d373561343132303065616234020a746573742d7361736c02000000000000")
	fetchReq, err := DecodeOffsetFetchReq(bytes, 6)
	if err != nil {
		panic(err)
	}
	if fetchReq.CorrelationId != 11 {
		codecError(t)
	}
	if fetchReq.ClientId != "consumer-a032ab3c-0818-4975-bd93-75a41200eab4-1" {
		codecError(t)
	}
	if fetchReq.GroupId != "a032ab3c-0818-4975-bd93-75a41200eab4" {
		codecError(t)
	}
	if len(fetchReq.TopicReqList) != 1 {
		codecError(t)
	}
	fetchTopicReq := fetchReq.TopicReqList[0]
	if fetchTopicReq.Topic != "test-sasl" {
		codecError(t)
	}
	if len(fetchTopicReq.PartitionReqList) != 1 {
		codecError(t)
	}
	fetchPartitionReq := fetchTopicReq.PartitionReqList[0]
	if fetchPartitionReq.PartitionId != 0 {
		codecError(t)
	}
	if fetchReq.RequireStableOffset {
		codecError(t)
	}
}
