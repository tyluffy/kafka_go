package kafka

import (
	"flag"
	"github.com/paashzj/kafka_go/pkg/codec"
	"github.com/paashzj/kafka_go/pkg/log"
	"github.com/paashzj/kafka_go/pkg/network"
	"github.com/paashzj/kafka_go/pkg/service"
	"k8s.io/klog/v2"
	"os"
)

type ServerConfig struct {
	LogLevel string

	// 网络配置
	ListenAddr string
	MultiCore  bool
	NeedSasl   bool
	MaxConn    int32

	// Kafka协议配置
	ClusterId     string
	AdvertiseHost string
	AdvertisePort int
}

func Run(config *ServerConfig, impl service.KfkServer) (*ServerControl, error) {
	flagSet := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(flagSet)
	_ = flagSet.Set("v", config.LogLevel)
	klog.SetOutput(os.Stdout)
	log.Codec().Info("This is codec message, you will see the message of codec")
	log.Network().Info("This is network message， you will see the message of network")
	networkConfig := &network.Config{}
	networkConfig.ListenAddr = config.ListenAddr
	networkConfig.MultiCore = config.MultiCore
	kfkProtocolConfig := &codec.KafkaProtocolConfig{}
	kfkProtocolConfig.ClusterId = config.ClusterId
	kfkProtocolConfig.AdvertiseHost = config.AdvertiseHost
	kfkProtocolConfig.AdvertisePort = config.AdvertisePort
	kfkProtocolConfig.NeedSasl = config.NeedSasl
	kfkProtocolConfig.MaxConn = config.MaxConn
	serverControl := &ServerControl{}
	var err error
	serverControl.networkServer, err = network.Run(networkConfig, kfkProtocolConfig, impl)
	return serverControl, err
}
