package network

import (
	"github.com/paashzj/kafka_go/pkg/kafka/codec"
	"github.com/paashzj/kafka_go/pkg/kafka/log"
	"github.com/paashzj/kafka_go/pkg/kafka/network/context"
	"github.com/paashzj/kafka_go/pkg/kafka/service"
	"github.com/panjf2000/gnet"
	"k8s.io/klog/v2"
)

func (s *Server) OffsetFetch(ctx *context.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	if version == 6 || version == 7 {
		return s.OffsetFetchVersion(ctx, frame, version)
	}
	klog.Error("unknown offset fetch version ", version)
	return nil, gnet.Close
}

func (s *Server) OffsetFetchVersion(ctx *context.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	saslReq, ok := s.SaslMap.Load(ctx.Addr)
	if !ok {
		return nil, gnet.Close
	}
	req, err := codec.DecodeOffsetFetchReq(frame, version)
	if err != nil {
		return nil, gnet.Close
	}
	log.Codec().Info("offset fetch req ", req)
	lowReq := &service.OffsetFetchReq{}
	lowReq.TopicReqList = make([]*service.OffsetFetchTopicReq, len(req.TopicReqList))
	for i, topicReq := range req.TopicReqList {
		res, code := s.kafkaImpl.SaslAuthTopic(saslReq.(service.SaslReq), topicReq.Topic)
		if code != 0 || !res {
			return nil, gnet.Close
		}
		lowTopicReq := &service.OffsetFetchTopicReq{}
		lowTopicReq.Topic = topicReq.Topic
		lowTopicReq.PartitionReqList = make([]*service.OffsetFetchPartitionReq, len(topicReq.PartitionReqList))
		for j, partitionReq := range topicReq.PartitionReqList {
			lowPartitionReq := &service.OffsetFetchPartitionReq{}
			lowPartitionReq.PartitionId = partitionReq.PartitionId
			lowTopicReq.PartitionReqList[j] = lowPartitionReq
		}
		lowReq.TopicReqList[i] = lowTopicReq
	}
	lowResp, err := service.OffsetFetch(ctx.Addr, s.kafkaImpl, lowReq)
	if err != nil {
		return nil, gnet.Close
	}
	resp := codec.NewOffsetFetchResp(req.CorrelationId)
	resp.ErrorCode = lowResp.ErrorCode
	resp.TopicRespList = make([]*codec.OffsetFetchTopicResp, len(lowResp.TopicRespList))
	for i, lowTopicResp := range lowResp.TopicRespList {
		f := &codec.OffsetFetchTopicResp{}
		f.Topic = lowTopicResp.Topic
		f.PartitionRespList = make([]*codec.OffsetFetchPartitionResp, len(lowTopicResp.PartitionRespList))
		for j, p := range lowTopicResp.PartitionRespList {
			partitionResp := &codec.OffsetFetchPartitionResp{}
			partitionResp.PartitionId = p.PartitionId
			partitionResp.Offset = p.Offset
			partitionResp.LeaderEpoch = p.LeaderEpoch
			partitionResp.Metadata = p.Metadata
			f.PartitionRespList[j] = partitionResp
		}
		resp.TopicRespList[i] = f
	}
	return resp.Bytes(version), gnet.None
}
