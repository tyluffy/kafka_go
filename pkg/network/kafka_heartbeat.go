package network

import (
	codec2 "github.com/paashzj/kafka_go/pkg/codec"
	"github.com/panjf2000/gnet"
	"k8s.io/klog/v2"
)

func (s *Server) Heartbeat(frame []byte, version int16) ([]byte, gnet.Action) {
	if version == 4 {
		return s.ReactHeartbeatVersion(frame, version)
	}
	klog.Error("unknown heartbeat version ", version)
	return nil, gnet.Close
}

func (s *Server) ReactHeartbeatVersion(frame []byte, version int16) ([]byte, gnet.Action) {
	heartbeatReqV4, err := codec2.DecodeHeartbeatReqV4(frame)
	if err != nil {
		return nil, gnet.Close
	}
	klog.Info("heart beat req ", heartbeatReqV4)
	heartBeatResp := codec2.NewHeartBeatResp(heartbeatReqV4.CorrelationId)
	return heartBeatResp.Bytes(), gnet.None
}
