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

func (s *Server) Fetch(ctx *ctx.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	if version == 10 || version == 11 {
		return s.ReactFetchVersion(ctx, frame, version)
	}
	logrus.Error("unknown fetch version ", version)
	return nil, gnet.Close
}

func (s *Server) ReactFetchVersion(ctx *ctx.NetworkContext, frame []byte, version int16) ([]byte, gnet.Action) {
	req, r, stack := codec.DecodeFetchReq(frame, version)
	if r != nil {
		logrus.Warn("decode sync group error", r, string(stack))
		return nil, gnet.Close
	}
	if !s.checkSasl(ctx) {
		return nil, gnet.Close
	}
	logrus.Debug("fetch req ", req)
	lowReq := &service.FetchReq{}
	lowReq.MaxBytes = req.MaxBytes
	lowReq.MinBytes = req.MinBytes
	lowReq.MaxWaitTime = req.MaxWaitTime
	lowReq.FetchTopicReqList = make([]*service.FetchTopicReq, len(req.TopicReqList))
	for i, topicReq := range req.TopicReqList {
		if !s.checkSaslTopic(ctx, topicReq.Topic, CONSUMER_PERMISSION_TYPE) {
			return nil, gnet.Close
		}
		lowTopicReq := &service.FetchTopicReq{}
		lowTopicReq.Topic = topicReq.Topic
		lowTopicReq.FetchPartitionReqList = make([]*service.FetchPartitionReq, len(topicReq.PartitionReqList))
		for j, partitionReq := range topicReq.PartitionReqList {
			lowPartitionReq := &service.FetchPartitionReq{}
			lowPartitionReq.PartitionId = partitionReq.PartitionId
			lowPartitionReq.FetchOffset = partitionReq.FetchOffset
			lowPartitionReq.ClientId = req.ClientId
			lowTopicReq.FetchPartitionReqList[j] = lowPartitionReq
		}
		lowReq.FetchTopicReqList[i] = lowTopicReq
	}
	lowTopicRespList, err := service.Fetch(ctx.Addr, s.kafkaImpl, lowReq)
	if err != nil {
		return nil, gnet.Close
	}
	resp := codec.NewFetchResp(req.CorrelationId)
	resp.TopicRespList = make([]*codec.FetchTopicResp, len(lowTopicRespList))
	for i, lowTopicResp := range lowTopicRespList {
		f := &codec.FetchTopicResp{}
		f.Topic = lowTopicResp.Topic
		f.PartitionRespList = make([]*codec.FetchPartitionResp, len(lowTopicResp.FetchPartitionRespList))
		for j, p := range lowTopicResp.FetchPartitionRespList {
			partitionResp := &codec.FetchPartitionResp{}
			partitionResp.PartitionIndex = p.PartitionId
			partitionResp.ErrorCode = int16(p.ErrorCode)
			partitionResp.HighWatermark = p.HighWatermark
			partitionResp.LastStableOffset = p.LastStableOffset
			partitionResp.LogStartOffset = p.LogStartOffset
			if p.RecordBatch != nil {
				partitionResp.RecordBatch = s.convertRecordBatchResp(p.RecordBatch)
			}
			partitionResp.AbortedTransactions = -1
			partitionResp.ReplicaId = -1
			f.PartitionRespList[j] = partitionResp
		}
		resp.TopicRespList[i] = f
	}
	return resp.Bytes(version), gnet.None
}

func (s *Server) convertRecordBatchResp(lowRecordBatch *service.RecordBatch) *codec.RecordBatch {
	recordBatch := &codec.RecordBatch{}
	recordBatch.Offset = lowRecordBatch.Offset
	recordBatch.LeaderEpoch = 0
	recordBatch.MagicByte = 2
	recordBatch.Flags = 0
	recordBatch.LastOffsetDelta = lowRecordBatch.LastOffsetDelta
	recordBatch.FirstTimestamp = lowRecordBatch.FirstTimestamp
	recordBatch.LastTimestamp = lowRecordBatch.LastTimestamp
	recordBatch.ProducerId = -1
	recordBatch.ProducerEpoch = -1
	recordBatch.BaseSequence = lowRecordBatch.BaseSequence
	recordBatch.Records = make([]*codec.Record, len(lowRecordBatch.Records))
	for i, r := range lowRecordBatch.Records {
		record := &codec.Record{}
		record.RecordAttributes = 0
		record.RelativeTimestamp = r.RelativeTimestamp
		record.RelativeOffset = r.RelativeOffset
		record.Key = r.Key
		record.Value = r.Value
		record.Headers = r.Headers
		recordBatch.Records[i] = record
	}
	return recordBatch
}
