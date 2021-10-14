package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeJoinGroupReqV6(t *testing.T) {
	bytes := testHex2Bytes(t, "00000008002f636f6e73756d65722d37336664633964612d306439322d346537622d613761372d6563323636663637633137312d31002537336664633964612d306439322d346537622d613761372d65633236366636376331373100002710000493e0010009636f6e73756d6572020672616e676535000100000001002437363465646565332d303037652d343865302d623966392d646637663731336666373037ffffffff000000000000")
	joinGroupReq, err := DecodeJoinGroupReq(bytes, 6)
	assert.Nil(t, err)
	assert.Equal(t, 8, joinGroupReq.CorrelationId)
	assert.Equal(t, "consumer-73fdc9da-0d92-4e7b-a7a7-ec266f67c171-1", joinGroupReq.ClientId)
	assert.Equal(t, "73fdc9da-0d92-4e7b-a7a7-ec266f67c171", joinGroupReq.GroupId)
	assert.Equal(t, 10_000, joinGroupReq.SessionTimeout)
	assert.Equal(t, 300_000, joinGroupReq.RebalanceTimeout)
	assert.Equal(t, "", joinGroupReq.MemberId)
	assert.Equal(t, "consumer", joinGroupReq.ProtocolType)
	assert.Len(t, joinGroupReq.GroupProtocols, 1)
	groupProtocol := joinGroupReq.GroupProtocols[0]
	assert.Equal(t, "range", groupProtocol.ProtocolName)
}
