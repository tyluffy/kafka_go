package main

import (
	"fmt"
	"github.com/google/uuid"
	service2 "github.com/paashzj/kafka_go/pkg/service"
	"k8s.io/klog/v2"
	"net"
)

type ExampleKafkaImpl struct {
}

func (e ExampleKafkaImpl) FetchPartition(addr *net.Addr, topic string, req *service2.FetchPartitionReq) (*service2.FetchPartitionResp, error) {
	partitionResp := &service2.FetchPartitionResp{}
	partitionResp.PartitionId = req.PartitionId
	partitionResp.HighWatermark = 1
	partitionResp.LastStableOffset = 1
	partitionResp.LogStartOffset = 0
	partitionResp.RecordBatch = exampleRecordBatch()
	return partitionResp, nil
}

func exampleRecordBatch() *service2.RecordBatch {
	r := &service2.RecordBatch{}
	r.Offset = 0
	r.MessageSize = 64
	r.LastOffsetDelta = 0
	r.FirstTimestamp = 1625962021853
	r.LastTimestamp = 1625962021853
	r.BaseSequence = -1
	r.Records = make([]*service2.Record, 1)
	record := &service2.Record{}
	record.RelativeTimestamp = 0
	record.RelativeOffset = 0
	record.Key = nil
	record.Value = "ShootHzj"
	r.Records[0] = record
	return r
}

func (e ExampleKafkaImpl) GroupJoin(addr *net.Addr, req *service2.JoinGroupReq) (*service2.JoinGroupResp, error) {
	resp := &service2.JoinGroupResp{}
	if req.MemberId == "" {
		resp.ErrorCode = service2.MEMBER_ID_REQUIRED
		resp.GenerationId = -1
		resp.ProtocolName = ""
		resp.MemberId = uuid.New().String()
		return resp, nil
	}
	members := make([]*service2.Member, 1)
	groupProtocols := req.GroupProtocols
	members[0] = &service2.Member{MemberId: req.MemberId,
		GroupInstanceId: nil, Metadata: groupProtocols[0].ProtocolMetadata}
	resp.GenerationId = 1
	resp.ProtocolName = "range"
	resp.LeaderId = req.MemberId
	resp.MemberId = req.MemberId
	resp.Members = members
	return resp, nil
}

func (e ExampleKafkaImpl) GroupLeave(addr *net.Addr, req *service2.LeaveGroupReq) (*service2.LeaveGroupResp, error) {
	l := &service2.LeaveGroupResp{}
	return l, nil
}

func (e ExampleKafkaImpl) GroupSync(addr *net.Addr, req *service2.SyncGroupReq) (*service2.SyncGroupResp, error) {
	groupAssignments := req.GroupAssignments
	resp := &service2.SyncGroupResp{}
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

func (e ExampleKafkaImpl) OffsetListPartition(addr *net.Addr, topic string, req *service2.ListOffsetsPartitionReq) (*service2.ListOffsetsPartitionResp, error) {
	partitionResp := &service2.ListOffsetsPartitionResp{}
	partitionResp.PartitionId = req.PartitionId
	partitionResp.Time = -1
	partitionResp.Offset = 0
	return partitionResp, nil
}

func (e ExampleKafkaImpl) OffsetCommitPartition(addr *net.Addr, topic string, req *service2.OffsetCommitPartitionReq) (*service2.OffsetCommitPartitionResp, error) {
	partitionResp := &service2.OffsetCommitPartitionResp{}
	partitionResp.PartitionId = req.PartitionId
	partitionResp.ErrorCode = 0
	return partitionResp, nil
}

func (e ExampleKafkaImpl) OffsetFetch(addr *net.Addr, topic string, partition int) (*service2.OffsetFetchPartitionResp, error) {
	return nil, nil
}

func (e ExampleKafkaImpl) SaslAuth(req service2.SaslReq) (bool, service2.ErrorCode) {
	klog.Info("username ", req.Username, "password ", req.Password)
	return true, service2.NONE
}

func (e ExampleKafkaImpl) SaslAuthTopic(req service2.SaslReq, topic string) (bool, service2.ErrorCode) {
	klog.Info("username ", req.Username, "password ", req.Password, "topic ", topic)
	return true, service2.NONE
}

func (e ExampleKafkaImpl) SaslAuthConsumerGroup(req service2.SaslReq, consumerGroup string) (bool, service2.ErrorCode) {
	klog.Info("username ", req.Username, "password ", req.Password, "group ", consumerGroup)
	return true, service2.NONE
}

func (e ExampleKafkaImpl) Disconnect(addr *net.Addr) {
	fmt.Println("do nothing now.")
}
