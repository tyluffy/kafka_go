package network

import (
	"github.com/paashzj/kafka_go/pkg/kafka/network/context"
	"github.com/paashzj/kafka_go/pkg/kafka/service"
)

func (s *Server) checkSasl(ctx *context.NetworkContext) bool {
	if !s.kafkaProtocolConfig.NeedSasl {
		return true
	}
	_, ok := s.SaslMap.Load(ctx.Addr)
	return ok
}

func (s *Server) checkSaslGroup(ctx *context.NetworkContext, groupId string) bool {
	if !s.kafkaProtocolConfig.NeedSasl {
		return true
	}
	saslReq, ok := s.SaslMap.Load(ctx.Addr)
	if !ok {
		return false
	}
	res, code := s.kafkaImpl.SaslAuthConsumerGroup(saslReq.(service.SaslReq), groupId)
	if code != 0 || !res {
		return false
	}
	return true
}

func (s *Server) checkSaslTopic(ctx *context.NetworkContext, topic string) bool {
	if !s.kafkaProtocolConfig.NeedSasl {
		return true
	}
	saslReq, ok := s.SaslMap.Load(ctx.Addr)
	if !ok {
		return false
	}
	res, code := s.kafkaImpl.SaslAuthTopic(saslReq.(service.SaslReq), topic)
	if code != 0 || !res {
		return false
	}
	return true
}
