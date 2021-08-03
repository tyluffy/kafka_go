package service

type SaslReq struct {
	Username string
	Password string
}

func SaslAuth(impl KfkServer, req SaslReq) (bool, ErrorCode) {
	return impl.SaslAuth(req)
}
