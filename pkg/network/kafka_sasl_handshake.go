package network

import (
	codec2 "github.com/paashzj/kafka_go/pkg/codec"
	"github.com/paashzj/kafka_go/pkg/log"
	"github.com/panjf2000/gnet"
	"k8s.io/klog/v2"
)

func (s *Server) SaslHandshake(frame []byte, version int16) ([]byte, gnet.Action) {
	if version == 1 || version == 2 {
		return s.ReactSaslVersion(frame, version)
	}
	klog.Error("unknown fetch version ", version)
	return nil, gnet.Close
}

func (s *Server) ReactSaslVersion(frame []byte, version int16) ([]byte, gnet.Action) {
	req, err := codec2.DecodeSaslHandshakeReq(frame, version)
	if err != nil {
		return nil, gnet.Close
	}
	log.Codec().Info("sasl handshake request ", req)
	saslHandshakeResp := codec2.NewSaslHandshakeResp(req.CorrelationId)
	saslHandshakeResp.EnableMechanisms = make([]*codec2.EnableMechanism, 1)
	saslHandshakeResp.EnableMechanisms[0] = &codec2.EnableMechanism{SaslMechanism: "PLAIN"}
	return saslHandshakeResp.Bytes(), gnet.None
}
