package network

import (
	"github.com/paashzj/kafka_go/pkg/kafka/codec"
	"github.com/paashzj/kafka_go/pkg/kafka/log"
	"github.com/paashzj/kafka_go/pkg/kafka/network/context"
	"github.com/paashzj/kafka_go/pkg/kafka/service"
	"github.com/panjf2000/gnet"
	"k8s.io/klog/v2"
)

func (s *Server) OffsetCommit(ctx *context.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	if version == 8 {
		return s.OffsetCommitVersion(ctx, frame, version)
	}
	klog.Error("unknown fetch version ", version)
	return nil, gnet.Close
}

func (s *Server) OffsetCommitVersion(ctx *context.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	req, err := codec.DecodeOffsetCommitReq(frame, version)
	if err != nil {
		return nil, gnet.Close
	}
	log.Codec().Info("offset commit req ", req)
	lowReqList := make([]*service.OffsetCommitTopicReq, len(req.OffsetCommitTopicReqList))
	for i, topicReq := range req.OffsetCommitTopicReqList {
		lowTopicReq := &service.OffsetCommitTopicReq{}
		lowTopicReq.Topic = topicReq.Topic
		lowTopicReq.OffsetCommitPartitionReqList = make([]*service.OffsetCommitPartitionReq, len(topicReq.OffsetTopicPartitions))
		for j, partitionReq := range topicReq.OffsetTopicPartitions {
			lowPartitionReq := &service.OffsetCommitPartitionReq{}
			lowPartitionReq.PartitionId = partitionReq.PartitionId
			lowPartitionReq.OffsetCommitOffset = partitionReq.Offset
			lowTopicReq.OffsetCommitPartitionReqList[j] = lowPartitionReq
		}
		lowReqList[i] = lowTopicReq
	}
	lowTopicRespList, err := service.OffsetCommit(ctx.Addr, s.kafkaImpl, lowReqList)
	if err != nil {
		return nil, gnet.Close
	}
	resp := codec.NewOffsetCommitResp(req.CorrelationId)
	resp.Topics = make([]*codec.OffsetCommitTopicResp, len(lowTopicRespList))
	for i, lowTopicResp := range lowTopicRespList {
		f := &codec.OffsetCommitTopicResp{}
		f.Topic = lowTopicResp.Topic
		f.Partitions = make([]*codec.OffsetCommitPartitionResp, len(lowTopicResp.OffsetCommitPartitionRespList))
		for j, p := range lowTopicResp.OffsetCommitPartitionRespList {
			partitionResp := &codec.OffsetCommitPartitionResp{}
			partitionResp.PartitionId = p.PartitionId
			partitionResp.ErrorCode = int16(p.ErrorCode)
			f.Partitions[j] = partitionResp
		}
		resp.Topics[i] = f
	}
	return resp.Bytes(version), gnet.None
}
