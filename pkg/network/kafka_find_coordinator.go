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
	"github.com/paashzj/kafka_go/pkg/codec"
	"github.com/panjf2000/gnet"
	"github.com/sirupsen/logrus"
)

func (s *Server) FindCoordinator(frame []byte, version int16, config *codec.KafkaProtocolConfig) ([]byte, gnet.Action) {
	if version == 0 || version == 3 {
		return s.FindCoordinatorVersion(frame, version, config)
	}
	logrus.Error("unknown find coordinator version ", version)
	return nil, gnet.Close
}

func (s *Server) FindCoordinatorVersion(frame []byte, version int16, config *codec.KafkaProtocolConfig) ([]byte, gnet.Action) {
	req, err := codec.DecodeFindCoordinatorReq(frame, version)
	if err != nil {
		return nil, gnet.Close
	}
	logrus.Info("req ", req)
	resp := codec.NewFindCoordinatorResp(req.CorrelationId, config)
	logrus.Info("resp ", resp)
	return resp.Bytes(version), gnet.None
}
