package low

func SaslAuth(impl KfkImpl, username string, password string) (bool, ErrorCode) {
	return impl.SaslAuth(username, password)
}
