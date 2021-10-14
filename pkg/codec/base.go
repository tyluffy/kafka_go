package codec

type BaseReq struct {
	CorrelationId int
	ClientId      string
}

type BaseResp struct {
	CorrelationId int
}
