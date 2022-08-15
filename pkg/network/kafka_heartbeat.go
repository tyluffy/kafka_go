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

func (s *Server) Heartbeat(frame []byte, version int16, context *ctx.NetworkContext) ([]byte, gnet.Action) {
	if version == 4 {
		return s.ReactHeartbeatVersion(frame, version, context)
	}
	logrus.Error("unknown heartbeat version ", version)
	return nil, gnet.Close
}

func (s *Server) ReactHeartbeatVersion(frame []byte, version int16, context *ctx.NetworkContext) ([]byte, gnet.Action) {
	heartbeatReqV4, r, stack := codec.DecodeHeartbeatReq(frame, version)
	if r != nil {
		logrus.Warn("decode sync group error", r, string(stack))
		return nil, gnet.Close
	}
	logrus.Debug("heart beat req ", heartbeatReqV4)
	heartBeatResp := codec.HeartbeatResp{
		BaseResp: codec.BaseResp{
			CorrelationId: heartbeatReqV4.CorrelationId,
		},
	}
	req := service.HeartBeatReq{}
	req.ClientId = heartbeatReqV4.ClientId
	req.GenerationId = heartbeatReqV4.GenerationId
	req.GroupInstanceId = heartbeatReqV4.GroupInstanceId
	req.MemberId = heartbeatReqV4.MemberId
	req.GroupId = heartbeatReqV4.GroupId
	beat := s.kafkaImpl.HeartBeat(context.Addr, req)
	heartBeatResp.ErrorCode = int16(beat.ErrorCode)
	return heartBeatResp.Bytes(version), gnet.None
}
