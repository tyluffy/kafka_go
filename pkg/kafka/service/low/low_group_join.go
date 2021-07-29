package low

type JoinGroupReq struct {
	GroupId         string
	SessionTimeout  int
	MemberId        string
	GroupInstanceId *string
	ProtocolType    string
	GroupProtocols  []*GroupProtocol
}

type GroupProtocol struct {
	ProtocolName     string
	ProtocolMetadata string
}

type JoinGroupResp struct {
	ErrorCode    ErrorCode
	GenerationId int
	ProtocolType *string
	ProtocolName string
	LeaderId     string
	MemberId     string
	Members      []*Member
}

type Member struct {
	MemberId        string
	GroupInstanceId *string
	Metadata        string
}
