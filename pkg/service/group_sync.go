package service

type SyncGroupReq struct {
	GroupId          string
	GenerationId     int
	MemberId         string
	GroupInstanceId  *string
	ProtocolType     string
	ProtocolName     string
	GroupAssignments []*GroupAssignment
}

type GroupAssignment struct {
	MemberId         string
	MemberAssignment string
}

type SyncGroupResp struct {
	ErrorCode        ErrorCode
	ProtocolType     string
	ProtocolName     string
	MemberAssignment string
}
