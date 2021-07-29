package low

import "net"

type OffsetCommitTopicReq struct {
	Topic                        string
	OffsetCommitPartitionReqList []*OffsetCommitPartitionReq
}

type OffsetCommitTopicResp struct {
	Topic                         string
	OffsetCommitPartitionRespList []*OffsetCommitPartitionResp
}

type OffsetCommitPartitionReq struct {
	PartitionId        int
	OffsetCommitOffset int64
}

type OffsetCommitPartitionResp struct {
	PartitionId int
	ErrorCode   ErrorCode
}

func OffsetCommit(addr *net.Addr, impl KfkImpl, reqList []*OffsetCommitTopicReq) ([]*OffsetCommitTopicResp, error) {
	result := make([]*OffsetCommitTopicResp, len(reqList))
	for i, req := range reqList {
		f := &OffsetCommitTopicResp{}
		f.Topic = req.Topic
		f.OffsetCommitPartitionRespList = make([]*OffsetCommitPartitionResp, len(req.OffsetCommitPartitionReqList))
		for j, partitionReq := range req.OffsetCommitPartitionReqList {
			// todo ignore error
			f.OffsetCommitPartitionRespList[j], _ = impl.OffsetCommitPartition(addr, req.Topic, partitionReq)
		}
		result[i] = f
	}
	return result, nil
}
