package service

import (
	"net"
)

type FetchReq struct {
	MinBytes          int
	MaxBytes          int
	MaxWaitTime       int
	FetchTopicReqList []*FetchTopicReq
}

type FetchTopicReq struct {
	Topic                 string
	FetchPartitionReqList []*FetchPartitionReq
}

type FetchTopicResp struct {
	Topic                  string
	FetchPartitionRespList []*FetchPartitionResp
}

type FetchPartitionReq struct {
	PartitionId int
	FetchOffset int64
}

type FetchPartitionResp struct {
	PartitionId      int
	HighWatermark    int64
	LastStableOffset int64
	LogStartOffset   int64
	RecordBatch      *RecordBatch
}

type RecordBatch struct {
	Offset          int64
	MessageSize     int
	LastOffsetDelta int
	FirstTimestamp  int64
	LastTimestamp   int64
	BaseSequence    int
	Records         []*Record
}

type Record struct {
	RelativeTimestamp int64
	RelativeOffset    int
	Key               []byte
	Value             string
	Headers           []byte
}

func Fetch(addr *net.Addr, impl KfkServer, req *FetchReq) ([]*FetchTopicResp, error) {
	reqList := req.FetchTopicReqList
	result := make([]*FetchTopicResp, len(reqList))
	for i, req := range reqList {
		f := &FetchTopicResp{}
		f.Topic = req.Topic
		f.FetchPartitionRespList = make([]*FetchPartitionResp, len(req.FetchPartitionReqList))
		for j, partitionReq := range req.FetchPartitionReqList {
			// todo error 处理
			partition, _ := impl.FetchPartition(addr, req.Topic, partitionReq)
			f.FetchPartitionRespList[j] = partition
		}
		result[i] = f
	}
	return result, nil
}
