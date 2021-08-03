package network

import (
	"github.com/paashzj/kafka_go/pkg/kafka/codec"
	"github.com/paashzj/kafka_go/pkg/kafka/log"
	"github.com/paashzj/kafka_go/pkg/kafka/network/context"
	"github.com/paashzj/kafka_go/pkg/kafka/service"
	"github.com/panjf2000/gnet"
	"k8s.io/klog/v2"
)

func (s *Server) ListOffsets(ctx *context.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	if version == 5 {
		return s.ListOffsetsVersion(ctx, frame, version)
	}
	klog.Error("unknown offset version ", version)
	return nil, gnet.Close
}

func (s *Server) ListOffsetsVersion(ctx *context.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	saslReq, ok := s.SaslMap.Load(ctx.Addr)
	if !ok {
		return nil, gnet.Close
	}
	req, err := codec.DecodeListOffsetReq(frame, version)
	if err != nil {
		return nil, gnet.Close
	}
	log.Codec().Info("list offset req ", req)
	lowOffsetReqList := make([]*service.ListOffsetsTopicReq, len(req.OffsetTopics))
	for i, topicReq := range req.OffsetTopics {
		res, code := s.kafkaImpl.SaslAuthTopic(saslReq.(service.SaslReq), topicReq.Topic)
		if code != 0 || !res {
			return nil, gnet.Close
		}
		lowTopicReq := &service.ListOffsetsTopicReq{}
		lowTopicReq.Topic = topicReq.Topic
		lowTopicReq.OffsetPartitionReqList = make([]*service.ListOffsetsPartitionReq, len(topicReq.OffsetTopicPartitions))
		for j, partitionReq := range topicReq.OffsetTopicPartitions {
			lowPartitionReq := &service.ListOffsetsPartitionReq{}
			lowPartitionReq.PartitionId = partitionReq.PartitionId
			lowTopicReq.OffsetPartitionReqList[j] = lowPartitionReq
		}
		lowOffsetReqList[i] = lowTopicReq
	}
	lowOffsetRespList, err := service.Offset(ctx.Addr, s.kafkaImpl, lowOffsetReqList)
	if err != nil {
		return nil, gnet.Close
	}
	resp := codec.NewOffsetResp(req.CorrelationId)
	resp.OffsetTopics = make([]*codec.ListOffsetTopicResp, len(lowOffsetRespList))
	for i, lowTopicResp := range lowOffsetRespList {
		f := &codec.ListOffsetTopicResp{}
		f.Topic = lowTopicResp.Topic
		f.OffsetTopicPartitions = make([]*codec.ListOffsetPartitionResp, len(lowTopicResp.OffsetPartitionRespList))
		for j, p := range lowTopicResp.OffsetPartitionRespList {
			partitionResp := &codec.ListOffsetPartitionResp{}
			partitionResp.PartitionId = p.PartitionId
			partitionResp.ErrorCode = 0
			partitionResp.Timestamp = p.Time
			partitionResp.Offset = p.Offset
			partitionResp.LeaderEpoch = 0
			f.OffsetTopicPartitions[j] = partitionResp
		}
		resp.OffsetTopics[i] = f
	}
	return resp.Bytes(version), gnet.None
}
