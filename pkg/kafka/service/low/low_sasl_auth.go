package low

func SaslAuth(impl KfkServer, username string, password string) (bool, ErrorCode) {
	return impl.SaslAuth(username, password)
}
