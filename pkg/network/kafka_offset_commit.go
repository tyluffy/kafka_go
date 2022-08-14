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

package network

import (
	"github.com/paashzj/kafka_go/pkg/network/ctx"
	"github.com/paashzj/kafka_go/pkg/service"
	"github.com/panjf2000/gnet"
	"github.com/protocol-laboratory/kafka-codec-go/codec"
	"github.com/sirupsen/logrus"
)

func (s *Server) OffsetCommit(ctx *ctx.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	if version == 2 || version == 8 {
		return s.OffsetCommitVersion(ctx, frame, version)
	}
	logrus.Error("unknown fetch version ", version)
	return nil, gnet.Close
}

func (s *Server) OffsetCommitVersion(ctx *ctx.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	req, r, stack := codec.DecodeOffsetCommitReq(frame, version)
	if r != nil {
		logrus.Warn("decode sync group error", r, string(stack))
		return nil, gnet.Close
	}
	if !s.checkSasl(ctx) {
		return nil, gnet.Close
	}
	logrus.Debug("offset commit req ", req)
	lowReqList := make([]*service.OffsetCommitTopicReq, len(req.TopicReqList))
	for i, topicReq := range req.TopicReqList {
		if !s.checkSaslTopic(ctx, topicReq.Topic, CONSUMER_PERMISSION_TYPE) {
			return nil, gnet.Close
		}
		lowTopicReq := &service.OffsetCommitTopicReq{}
		lowTopicReq.Topic = topicReq.Topic
		lowTopicReq.ReqList = make([]*service.OffsetCommitPartitionReq, len(topicReq.PartitionReqList))
		for j, partitionReq := range topicReq.PartitionReqList {
			lowPartitionReq := &service.OffsetCommitPartitionReq{}
			lowPartitionReq.PartitionId = partitionReq.PartitionId
			lowPartitionReq.OffsetCommitOffset = partitionReq.Offset
			lowPartitionReq.ClientId = req.ClientId
			lowTopicReq.ReqList[j] = lowPartitionReq
		}
		lowReqList[i] = lowTopicReq
	}
	lowTopicRespList, err := service.OffsetCommit(ctx.Addr, s.kafkaImpl, lowReqList)
	if err != nil {
		return nil, gnet.Close
	}
	resp := codec.OffsetCommitResp{
		BaseResp: codec.BaseResp{
			CorrelationId: req.CorrelationId,
		},
	}
	resp.TopicRespList = make([]*codec.OffsetCommitTopicResp, len(lowTopicRespList))
	for i, lowTopicResp := range lowTopicRespList {
		f := &codec.OffsetCommitTopicResp{}
		f.Topic = lowTopicResp.Topic
		f.PartitionRespList = make([]*codec.OffsetCommitPartitionResp, len(lowTopicResp.PartitionRespList))
		for j, p := range lowTopicResp.PartitionRespList {
			partitionResp := &codec.OffsetCommitPartitionResp{}
			partitionResp.PartitionId = p.PartitionId
			partitionResp.ErrorCode = int16(p.ErrorCode)
			f.PartitionRespList[j] = partitionResp
		}
		resp.TopicRespList[i] = f
	}
	return resp.Bytes(version), gnet.None
}
