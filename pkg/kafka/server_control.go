package kafka

import (
	"errors"
	"github.com/paashzj/kafka_go/pkg/network"
	"github.com/panjf2000/gnet"
	"net"
)

type ServerControl struct {
	networkServer *network.Server
}

func (s *ServerControl) DisConnect(addr net.Addr) error {
	load, ok := s.networkServer.ConnMap.Load(addr)
	if !ok {
		return errors.New("no such addr connection")
	}
	conn := load.(gnet.Conn)
	return conn.Close()
}
