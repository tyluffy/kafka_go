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

package codec

import (
	"errors"
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

type SaslHandshakeReq struct {
	BaseReq
	SaslMechanism string
}

// DecodeSaslHandshakeReq SaslHandshakeReq
func DecodeSaslHandshakeReq(bytes []byte, version int16) (saslHandshakeReq *SaslHandshakeReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Info("Recovered in f", r, string(debug.Stack()))
			saslHandshakeReq = nil
			err = errors.New("codec failed")
		}
	}()
	req := &SaslHandshakeReq{}
	idx := 0
	req.CorrelationId, idx = readCorrId(bytes, idx)
	req.ClientId, idx = readClientId(bytes, idx)
	req.SaslMechanism, idx = readSaslMechanism(bytes, idx)
	return req, nil
}
