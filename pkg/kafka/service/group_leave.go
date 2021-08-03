package service

type LeaveGroupReq struct {
	GroupId string
	Members []*LeaveGroupMember
}

type LeaveGroupMember struct {
	MemberId        string
	GroupInstanceId *string
}

type LeaveGroupResp struct {
	ErrorCode       ErrorCode
	Members         []*LeaveGroupMember
	MemberErrorCode ErrorCode
}
