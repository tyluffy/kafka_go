package low

import (
	"net"
)

type ListOffsetsTopicReq struct {
	Topic                  string
	OffsetPartitionReqList []*ListOffsetsPartitionReq
}

type ListOffsetsTopicResp struct {
	Topic                   string
	OffsetPartitionRespList []*ListOffsetsPartitionResp
}

type ListOffsetsPartitionReq struct {
	PartitionId int
}

type ListOffsetsPartitionResp struct {
	PartitionId int
	Time        int64
	Offset      int64
}

func Offset(addr *net.Addr, impl KfkServer, reqList []*ListOffsetsTopicReq) ([]*ListOffsetsTopicResp, error) {
	result := make([]*ListOffsetsTopicResp, len(reqList))
	for i, req := range reqList {
		f := &ListOffsetsTopicResp{}
		f.Topic = req.Topic
		f.OffsetPartitionRespList = make([]*ListOffsetsPartitionResp, len(req.OffsetPartitionReqList))
		for j, partitionReq := range req.OffsetPartitionReqList {
			// todo ignore error
			f.OffsetPartitionRespList[j], _ = impl.OffsetListPartition(addr, f.Topic, partitionReq)
		}
		result[i] = f
	}
	return result, nil
}
