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

package service

import (
	"net"
)

type KfkServer interface {
	PartitionNum(addr net.Addr, topic string) (int, error)

	// Fetch method called this already authed
	Fetch(addr net.Addr, req *FetchReq) ([]*FetchTopicResp, error)

	// GroupJoin method called this already authed
	GroupJoin(addr net.Addr, req *JoinGroupReq) (*JoinGroupResp, error)

	// GroupLeave method called this already authed
	GroupLeave(addr net.Addr, req *LeaveGroupReq) (*LeaveGroupResp, error)

	// GroupSync method called this already authed
	GroupSync(addr net.Addr, req *SyncGroupReq) (*SyncGroupResp, error)

	// OffsetListPartition method called this already authed
	OffsetListPartition(addr net.Addr, topic string, req *ListOffsetsPartitionReq) (*ListOffsetsPartitionResp, error)

	// OffsetCommitPartition method called this already authed
	OffsetCommitPartition(addr net.Addr, topic string, req *OffsetCommitPartitionReq) (*OffsetCommitPartitionResp, error)

	// OffsetFetch method called this already authed
	OffsetFetch(addr net.Addr, topic string, req *OffsetFetchPartitionReq) (*OffsetFetchPartitionResp, error)

	// OffsetLeaderEpoch method called this already authed
	OffsetLeaderEpoch(addr net.Addr, topic string, req *OffsetLeaderEpochPartitionReq) (*OffsetLeaderEpochPartitionResp, error)

	// Produce method called this already authed
	Produce(addr net.Addr, topic string, partition int, req *ProducePartitionReq) (*ProducePartitionResp, error)

	SaslAuth(addr net.Addr, req SaslReq) (bool, ErrorCode)

	SaslAuthTopic(addr net.Addr, req SaslReq, topic, permissionType string) (bool, ErrorCode)

	SaslAuthConsumerGroup(addr net.Addr, req SaslReq, consumerGroup string) (bool, ErrorCode)

	HeartBeat(addr net.Addr, req HeartBeatReq) *HeartBeatResp

	Disconnect(addr net.Addr)
}
