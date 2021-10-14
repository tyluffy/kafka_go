package codec

import (
	"github.com/paashzj/kafka_go/pkg/codec/api"
)

// This file is for kafka code int16 type. Format method as alpha order.

func putApiKey(bytes []byte, idx int, x api.Code) int {
	return putInt16(bytes, idx, int16(x))
}

func putApiMaxVersion(bytes []byte, idx int, x int16) int {
	return putInt16(bytes, idx, x)
}

func putApiMinVersion(bytes []byte, idx int, x int16) int {
	return putInt16(bytes, idx, x)
}

func putProducerEpoch(bytes []byte, idx int, errorCode int16) int {
	return putInt16(bytes, idx, errorCode)
}

func putErrorCode(bytes []byte, idx int, errorCode int16) int {
	return putInt16(bytes, idx, errorCode)
}
