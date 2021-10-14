package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodeFindCoordinatorRespV3(t *testing.T) {
	protocolConfig := &KafkaProtocolConfig{}
	findCoordinatorResp := NewFindCoordinatorResp(0, protocolConfig)
	bytes := findCoordinatorResp.Bytes()
	expectBytes := testHex2Bytes(t, "0000001f000000000000000000000000000000000a6c6f63616c686f73740000238400")
	assert.Equal(t, expectBytes, bytes)
}
