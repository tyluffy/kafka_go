package network

import (
	codec2 "github.com/paashzj/kafka_go/pkg/codec"
	"github.com/paashzj/kafka_go/pkg/log"
	"github.com/panjf2000/gnet"
	"k8s.io/klog/v2"
)

func (s *Server) FindCoordinator(frame []byte, version int16, config *codec2.KafkaProtocolConfig) ([]byte, gnet.Action) {
	if version == 3 {
		return s.FindCoordinatorVersion(frame, version, config)
	}
	klog.Error("unknown find coordinator version ", version)
	return nil, gnet.Close
}

func (s *Server) FindCoordinatorVersion(frame []byte, version int16, config *codec2.KafkaProtocolConfig) ([]byte, gnet.Action) {
	req, err := codec2.DecodeFindCoordinatorReq(frame, version)
	if err != nil {
		return nil, gnet.Close
	}
	log.Codec().Info("req ", req)
	resp := codec2.NewFindCoordinatorResp(req.CorrelationId, config)
	log.Codec().Info("resp ", resp)
	return resp.Bytes(), gnet.None
}
