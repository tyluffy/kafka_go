package network

import (
	"github.com/paashzj/kafka_go/pkg/kafka/codec"
	"github.com/paashzj/kafka_go/pkg/kafka/log"
	"github.com/paashzj/kafka_go/pkg/kafka/network/context"
	"github.com/paashzj/kafka_go/pkg/kafka/service"
	"github.com/panjf2000/gnet"
	"k8s.io/klog/v2"
)

func (s *Server) SyncGroup(ctx *context.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	if version == 4 || version == 5 {
		return s.ReactSyncGroupVersion(ctx, frame, version)
	}
	klog.Error("unknown sync group version ", version)
	return nil, gnet.Close
}

func (s *Server) ReactSyncGroupVersion(ctx *context.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	req, err := codec.DecodeSyncGroupReq(frame, version)
	if err != nil {
		return nil, gnet.Close
	}
	log.Codec().Info("sync group req", req)
	lowReq := &service.SyncGroupReq{}
	lowReq.GroupId = req.GroupId
	lowReq.GenerationId = req.GenerationId
	lowReq.MemberId = req.MemberId
	lowReq.GroupInstanceId = req.GroupInstanceId
	lowReq.ProtocolType = req.ProtocolType
	lowReq.ProtocolName = req.ProtocolName
	lowReq.GroupAssignments = make([]*service.GroupAssignment, len(req.GroupAssignments))
	for i, groupAssignment := range req.GroupAssignments {
		g := &service.GroupAssignment{}
		g.MemberAssignment = groupAssignment.MemberAssignment
		g.MemberId = groupAssignment.MemberId
		lowReq.GroupAssignments[i] = g
	}
	resp := codec.NewSyncGroupResp(req.CorrelationId)
	lowResp, err := s.kafkaImpl.GroupSync(ctx.Addr, lowReq)
	if err != nil {
		return nil, gnet.Close
	}
	resp.ProtocolType = lowResp.ProtocolType
	resp.ProtocolName = lowResp.ProtocolName
	resp.MemberAssignment = lowResp.MemberAssignment
	return resp.Bytes(version), gnet.None
}
