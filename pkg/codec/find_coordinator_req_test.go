package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeIllegalFindCoordinatorReq(t *testing.T) {
	bytes := make([]byte, 0)
	_, err := DecodeFindCoordinatorReq(bytes, 0)
	assert.NotNil(t, err)
}

func TestDecodeFindCoordinatorReqV3(t *testing.T) {
	bytes := testHex2Bytes(t, "00000000002f636f6e73756d65722d37336664633964612d306439322d346537622d613761372d6563323636663637633137312d31002537336664633964612d306439322d346537622d613761372d6563323636663637633137310000")
	findCoordinatorReq, err := DecodeFindCoordinatorReq(bytes, 3)
	assert.Nil(t, err)
	assert.Equal(t, 0, findCoordinatorReq.CorrelationId)
	assert.Equal(t, "consumer-73fdc9da-0d92-4e7b-a7a7-ec266f67c171-1", findCoordinatorReq.ClientId)
	assert.Equal(t, "73fdc9da-0d92-4e7b-a7a7-ec266f67c171", findCoordinatorReq.Key)
	var expectKeyType uint8 = 0
	assert.Equal(t, expectKeyType, findCoordinatorReq.KeyType)
}
