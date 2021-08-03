package network

import (
	"github.com/paashzj/kafka_go/pkg/kafka/codec"
	"github.com/paashzj/kafka_go/pkg/kafka/log"
	"github.com/paashzj/kafka_go/pkg/kafka/network/context"
	"github.com/paashzj/kafka_go/pkg/kafka/service"
	"github.com/panjf2000/gnet"
	"k8s.io/klog/v2"
)

func (s *Server) JoinGroup(ctx *context.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	if version == 6 || version == 7 {
		return s.ReactJoinGroupVersion(ctx, frame, version)
	}
	klog.Error("unknown join group version ", version)
	return nil, gnet.Close
}

func (s *Server) ReactJoinGroupVersion(ctx *context.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	saslReq, ok := s.SaslMap.Load(ctx.Addr)
	if !ok {
		return nil, gnet.Close
	}
	req, err := codec.DecodeJoinGroupReq(frame, version)
	res, code := s.kafkaImpl.SaslAuthConsumerGroup(saslReq.(service.SaslReq), req.GroupId)
	if code != 0 || !res {
		return nil, gnet.Close
	}
	if err != nil {
		return nil, gnet.Close
	}
	log.Codec().Info("join group req", req)
	lowReq := &service.JoinGroupReq{}
	lowReq.GroupId = req.GroupId
	lowReq.SessionTimeout = req.SessionTimeout
	lowReq.MemberId = req.MemberId
	lowReq.GroupInstanceId = req.GroupInstanceId
	lowReq.ProtocolType = req.ProtocolType
	lowReq.GroupProtocols = make([]*service.GroupProtocol, len(req.GroupProtocols))
	for i, groupProtocol := range req.GroupProtocols {
		g := &service.GroupProtocol{}
		g.ProtocolName = groupProtocol.ProtocolName
		g.ProtocolMetadata = groupProtocol.ProtocolMetadata
		lowReq.GroupProtocols[i] = g
	}
	resp := codec.NewJoinGroupResp(req.CorrelationId)
	lowResp, err := s.kafkaImpl.GroupJoin(ctx.Addr, lowReq)
	if err != nil {
		return nil, gnet.Close
	}
	log.Codec().Info("resp ", resp)
	resp.ErrorCode = int16(lowResp.ErrorCode)
	resp.GenerationId = lowResp.GenerationId
	resp.ProtocolType = lowResp.ProtocolType
	resp.ProtocolName = lowResp.ProtocolName
	resp.LeaderId = lowResp.LeaderId
	resp.MemberId = lowResp.MemberId
	resp.Members = make([]*codec.Member, len(lowResp.Members))
	for i, lowMember := range lowResp.Members {
		m := &codec.Member{}
		m.MemberId = lowMember.MemberId
		m.GroupInstanceId = lowMember.GroupInstanceId
		m.Metadata = lowMember.Metadata
		resp.Members[i] = m
	}
	return resp.Bytes(version), gnet.None
}
