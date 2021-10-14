package codec

type KafkaProtocolConfig struct {
	ClusterId     string
	AdvertiseHost string
	AdvertisePort int
	NeedSasl      bool
	MaxConn       int32
}
