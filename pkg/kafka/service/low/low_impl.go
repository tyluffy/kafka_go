package low

import "net"

type KfkImpl interface {
	FetchPartition(addr *net.Addr, topic string, req *FetchPartitionReq) (*FetchPartitionResp, error)

	GroupJoin(addr *net.Addr, req *JoinGroupReq) (*JoinGroupResp, error)

	GroupLeave(addr *net.Addr, req *LeaveGroupReq) (*LeaveGroupResp, error)

	GroupSync(addr *net.Addr, req *SyncGroupReq) (*SyncGroupResp, error)

	OffsetListPartition(addr *net.Addr, topic string, req *ListOffsetsPartitionReq) (*ListOffsetsPartitionResp, error)

	OffsetCommitPartition(addr *net.Addr, topic string, req *OffsetCommitPartitionReq) (*OffsetCommitPartitionResp, error)

	OffsetFetch(addr *net.Addr, topic string, partition int) (*OffsetFetchPartitionResp, error)

	SaslAuth(username string, password string) (bool, ErrorCode)

	Disconnect(addr *net.Addr)
}
