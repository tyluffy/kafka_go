package network

import (
	"github.com/paashzj/kafka_go/pkg/network/context"
	"github.com/panjf2000/gnet"
	"k8s.io/klog/v2"
)

func (s *Server) Authed(context *context.NetworkContext) bool {
	if !s.kafkaProtocolConfig.NeedSasl {
		return true
	}
	return context.IsAuthed()
}

func (s *Server) AuthFailed() ([]byte, gnet.Action) {
	klog.Error("auth failed")
	return nil, gnet.Close
}
