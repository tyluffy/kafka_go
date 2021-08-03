package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/paashzj/kafka_go/pkg/kafka/service/low"
	"k8s.io/klog/v2"
	"net"
)

type ExampleKafkaImpl struct {
}

func (e ExampleKafkaImpl) FetchPartition(addr *net.Addr, topic string, req *low.FetchPartitionReq) (*low.FetchPartitionResp, error) {
	partitionResp := &low.FetchPartitionResp{}
	partitionResp.PartitionId = req.PartitionId
	partitionResp.HighWatermark = 1
	partitionResp.LastStableOffset = 1
	partitionResp.LogStartOffset = 0
	partitionResp.RecordBatch = exampleRecordBatch()
	return partitionResp, nil
}

func exampleRecordBatch() *low.RecordBatch {
	r := &low.RecordBatch{}
	r.Offset = 0
	r.MessageSize = 64
	r.LastOffsetDelta = 0
	r.FirstTimestamp = 1625962021853
	r.LastTimestamp = 1625962021853
	r.BaseSequence = -1
	r.Records = make([]*low.Record, 1)
	record := &low.Record{}
	record.RelativeTimestamp = 0
	record.RelativeOffset = 0
	record.Key = nil
	record.Value = "ShootHzj"
	r.Records[0] = record
	return r
}

func (e ExampleKafkaImpl) GroupJoin(addr *net.Addr, req *low.JoinGroupReq) (*low.JoinGroupResp, error) {
	resp := &low.JoinGroupResp{}
	if req.MemberId == "" {
		resp.ErrorCode = low.MEMBER_ID_REQUIRED
		resp.GenerationId = -1
		resp.ProtocolName = ""
		resp.MemberId = uuid.New().String()
		return resp, nil
	}
	members := make([]*low.Member, 1)
	groupProtocols := req.GroupProtocols
	members[0] = &low.Member{MemberId: req.MemberId,
		GroupInstanceId: nil, Metadata: groupProtocols[0].ProtocolMetadata}
	resp.GenerationId = 1
	resp.ProtocolName = "range"
	resp.LeaderId = req.MemberId
	resp.MemberId = req.MemberId
	resp.Members = members
	return resp, nil
}

func (e ExampleKafkaImpl) GroupLeave(addr *net.Addr, req *low.LeaveGroupReq) (*low.LeaveGroupResp, error) {
	l := &low.LeaveGroupResp{}
	return l, nil
}

func (e ExampleKafkaImpl) GroupSync(addr *net.Addr, req *low.SyncGroupReq) (*low.SyncGroupResp, error) {
	groupAssignments := req.GroupAssignments
	resp := &low.SyncGroupResp{}
	if req.ProtocolType == "" {
		resp.ProtocolType = ""
		resp.ProtocolName = ""
	} else {
		resp.ProtocolType = "consumer"
		resp.ProtocolName = "range"
	}
	resp.MemberAssignment = groupAssignments[0].MemberAssignment
	return resp, nil
}

func (e ExampleKafkaImpl) OffsetListPartition(addr *net.Addr, topic string, req *low.ListOffsetsPartitionReq) (*low.ListOffsetsPartitionResp, error) {
	partitionResp := &low.ListOffsetsPartitionResp{}
	partitionResp.PartitionId = req.PartitionId
	partitionResp.Time = -1
	partitionResp.Offset = 0
	return partitionResp, nil
}

func (e ExampleKafkaImpl) OffsetCommitPartition(addr *net.Addr, topic string, req *low.OffsetCommitPartitionReq) (*low.OffsetCommitPartitionResp, error) {
	partitionResp := &low.OffsetCommitPartitionResp{}
	partitionResp.PartitionId = req.PartitionId
	partitionResp.ErrorCode = 0
	return partitionResp, nil
}

func (e ExampleKafkaImpl) OffsetFetch(addr *net.Addr, topic string, partition int) (*low.OffsetFetchPartitionResp, error) {
	return nil, nil
}

func (e ExampleKafkaImpl) SaslAuth(username string, password string) (bool, low.ErrorCode) {
	klog.Info("username ", username, "password ", password)
	return true, low.NONE
}

func (e ExampleKafkaImpl) Disconnect(addr *net.Addr) {
	fmt.Println("do nothing now.")
}

func (e ExampleKafkaImpl) Available(addr *net.Addr) bool {
	return true
}
