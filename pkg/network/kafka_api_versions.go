package network

import (
	codec2 "github.com/paashzj/kafka_go/pkg/codec"
	"github.com/paashzj/kafka_go/pkg/log"
	"github.com/panjf2000/gnet"
	"k8s.io/klog/v2"
)

func (s *Server) ApiVersions(frame []byte, version int16) ([]byte, gnet.Action) {
	if version == 0 || version == 3 {
		return s.ReactApiVersion(frame, version)
	}
	klog.Error("unknown api version ", version)
	return nil, gnet.Close
}

func (s *Server) ReactApiVersion(frame []byte, version int16) ([]byte, gnet.Action) {
	apiRequestV0, err := codec2.DecodeApiReq(frame, version)
	if err != nil {
		return nil, gnet.Close
	}
	log.Codec().Info("api request ", apiRequestV0)
	apiResponses := codec2.NewApiVersionResp(apiRequestV0.CorrelationId)
	return apiResponses.Bytes(version), gnet.None
}
