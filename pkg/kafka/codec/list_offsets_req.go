package codec

import (
	"errors"
	"k8s.io/klog/v2"
	"runtime/debug"
)

type ListOffsetReq struct {
	BaseReq
	ReplicaId      int
	IsolationLevel byte
	OffsetTopics   []*ListOffsetTopic
}

type ListOffsetTopic struct {
	Topic                 string
	OffsetTopicPartitions []*ListOffsetPartition
}

type ListOffsetPartition struct {
	PartitionId int
	LeaderEpoch int
	Time        int64
}

func DecodeListOffsetReq(bytes []byte, version int16) (offsetReq *ListOffsetReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			klog.Info("Recovered in f", r, string(debug.Stack()))
			offsetReq = nil
			err = errors.New("codec failed")
		}
	}()
	offsetReq = &ListOffsetReq{}
	idx := 0
	offsetReq.CorrelationId, idx = readCorrId(bytes, idx)
	offsetReq.ClientId, idx = readClientId(bytes, idx)
	offsetReq.ReplicaId, idx = readReplicaId(bytes, idx)
	offsetReq.IsolationLevel, idx = readIsolationLevel(bytes, idx)
	var length int
	length, idx = readInt(bytes, idx)
	offsetReq.OffsetTopics = make([]*ListOffsetTopic, length)
	for i := 0; i < length; i++ {
		topic := &ListOffsetTopic{}
		topic.Topic, idx = readTopicString(bytes, idx)
		var partitionLength int
		partitionLength, idx := readInt(bytes, idx)
		topic.OffsetTopicPartitions = make([]*ListOffsetPartition, partitionLength)
		for j := 0; j < partitionLength; j++ {
			partition := &ListOffsetPartition{}
			partition.PartitionId, idx = readInt(bytes, idx)
			partition.LeaderEpoch, idx = readLeaderEpoch(bytes, idx)
			partition.Time, idx = readTime(bytes, idx)
			topic.OffsetTopicPartitions[j] = partition
		}
		offsetReq.OffsetTopics[i] = topic
	}
	return offsetReq, nil
}
