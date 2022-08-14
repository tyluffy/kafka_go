// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/paashzj/kafka_go/pkg/service"
	"github.com/sirupsen/logrus"
	"net"
)

type ExampleKafkaImpl struct {
}

func (e ExampleKafkaImpl) PartitionNum(addr net.Addr, topic string) (int, error) {
	return 1, nil
}

func (e ExampleKafkaImpl) Fetch(addr net.Addr, req *service.FetchReq) ([]*service.FetchTopicResp, error) {
	reqList := req.FetchTopicReqList
	result := make([]*service.FetchTopicResp, len(reqList))
	for i, req := range reqList {
		f := &service.FetchTopicResp{}
		f.Topic = req.Topic
		f.FetchPartitionRespList = make([]*service.FetchPartitionResp, len(req.FetchPartitionReqList))
		for j, partitionReq := range req.FetchPartitionReqList {
			f.FetchPartitionRespList[j] = e.fetchPartition(partitionReq)
		}
		result[i] = f
	}
	return result, nil
}

func (e ExampleKafkaImpl) fetchPartition(req *service.FetchPartitionReq) *service.FetchPartitionResp {
	partitionResp := &service.FetchPartitionResp{}
	partitionResp.PartitionId = req.PartitionId
	partitionResp.HighWatermark = 1
	partitionResp.LastStableOffset = 1
	partitionResp.LogStartOffset = 0
	partitionResp.RecordBatch = exampleRecordBatch()
	return partitionResp
}

func exampleRecordBatch() *service.RecordBatch {
	r := &service.RecordBatch{}
	r.Offset = 0
	r.LastOffsetDelta = 0
	r.FirstTimestamp = 1625962021853
	r.LastTimestamp = 1625962021853
	r.BaseSequence = -1
	r.Records = make([]*service.Record, 1)
	record := &service.Record{}
	record.RelativeTimestamp = 0
	record.RelativeOffset = 0
	record.Key = nil
	record.Value = []byte("ShootHzj")
	r.Records[0] = record
	return r
}

func (e ExampleKafkaImpl) GroupJoin(addr net.Addr, req *service.JoinGroupReq) (*service.JoinGroupResp, error) {
	resp := &service.JoinGroupResp{}
	if req.MemberId == "" {
		resp.ErrorCode = service.MEMBER_ID_REQUIRED
		resp.GenerationId = -1
		resp.ProtocolName = ""
		resp.MemberId = uuid.New().String()
		return resp, nil
	}
	members := make([]*service.Member, 1)
	groupProtocols := req.GroupProtocols
	members[0] = &service.Member{MemberId: req.MemberId,
		GroupInstanceId: nil, Metadata: groupProtocols[0].ProtocolMetadata}
	resp.GenerationId = 1
	resp.ProtocolName = "range"
	resp.LeaderId = req.MemberId
	resp.MemberId = req.MemberId
	resp.Members = members
	return resp, nil
}

func (e ExampleKafkaImpl) GroupLeave(addr net.Addr, req *service.LeaveGroupReq) (*service.LeaveGroupResp, error) {
	l := &service.LeaveGroupResp{}
	return l, nil
}

func (e ExampleKafkaImpl) GroupSync(addr net.Addr, req *service.SyncGroupReq) (*service.SyncGroupResp, error) {
	groupAssignments := req.GroupAssignments
	resp := &service.SyncGroupResp{}
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

func (e ExampleKafkaImpl) OffsetListPartition(addr net.Addr, topic string, req *service.ListOffsetsPartitionReq) (*service.ListOffsetsPartitionResp, error) {
	partitionResp := &service.ListOffsetsPartitionResp{}
	partitionResp.PartitionId = req.PartitionId
	partitionResp.Time = -1
	partitionResp.Offset = 0
	return partitionResp, nil
}

func (e ExampleKafkaImpl) OffsetCommitPartition(addr net.Addr, topic string, req *service.OffsetCommitPartitionReq) (*service.OffsetCommitPartitionResp, error) {
	partitionResp := &service.OffsetCommitPartitionResp{}
	partitionResp.PartitionId = req.PartitionId
	partitionResp.ErrorCode = 0
	return partitionResp, nil
}

func (e ExampleKafkaImpl) OffsetFetch(addr net.Addr, topic string, req *service.OffsetFetchPartitionReq) (*service.OffsetFetchPartitionResp, error) {
	return nil, nil
}

func (e ExampleKafkaImpl) OffsetLeaderEpoch(addr net.Addr, topic string, req *service.OffsetLeaderEpochPartitionReq) (*service.OffsetLeaderEpochPartitionResp, error) {
	partitionResp := &service.OffsetLeaderEpochPartitionResp{}
	partitionResp.ErrorCode = 0
	partitionResp.PartitionId = 0
	partitionResp.LeaderEpoch = 0
	partitionResp.Offset = 0
	return partitionResp, nil
}

func (e ExampleKafkaImpl) Produce(addr net.Addr, topic string, partition int, req *service.ProducePartitionReq) (*service.ProducePartitionResp, error) {
	partitionResp := &service.ProducePartitionResp{}
	partitionResp.Time = -1
	return partitionResp, nil
}

func (e ExampleKafkaImpl) SaslAuth(addr net.Addr, req service.SaslReq) (bool, service.ErrorCode) {
	logrus.Info("username ", req.Username, "password ", req.Password)
	return true, service.NONE
}

func (e ExampleKafkaImpl) SaslAuthTopic(addr net.Addr, req service.SaslReq, topic, permissionType string) (bool, service.ErrorCode) {
	logrus.Info("username ", req.Username, "password ", req.Password, "topic ", topic)
	return true, service.NONE
}

func (e ExampleKafkaImpl) SaslAuthConsumerGroup(addr net.Addr, req service.SaslReq, consumerGroup string) (bool, service.ErrorCode) {
	logrus.Info("username ", req.Username, "password ", req.Password, "group ", consumerGroup)
	return true, service.NONE
}

func (e ExampleKafkaImpl) HeartBeat(addr net.Addr, req service.HeartBeatReq) *service.HeartBeatResp {
	return &service.HeartBeatResp{}
}

func (e ExampleKafkaImpl) Disconnect(addr net.Addr) {
	fmt.Println("do nothing now.")
}
