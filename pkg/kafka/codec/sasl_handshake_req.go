package codec

import (
	"errors"
	"k8s.io/klog/v2"
	"runtime/debug"
)

type SaslHandshakeReq struct {
	BaseReq
	ClientId      string
	SaslMechanism string
}

// DecodeSaslHandshakeReq SaslHandshakeReq
func DecodeSaslHandshakeReq(bytes []byte, version int16) (saslHandshakeReq *SaslHandshakeReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			klog.Info("Recovered in f", r, string(debug.Stack()))
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
