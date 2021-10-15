package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeIllegalSaslHandshakeAuthReq(t *testing.T) {
	bytes := make([]byte, 0)
	_, err := DecodeSaslHandshakeAuthReq(bytes, 0)
	assert.NotNil(t, err)
}

func TestDecodeSaslHandshakeAuthReqV2(t *testing.T) {
	bytes := testHex2Bytes(t, "7ffffffa002f636f6e73756d65722d33616332336137662d346333362d343064392d393964342d6163646134376430613438642d31000d00616c69636500616c69636500")
	saslHandshakeAuthReq, err := DecodeSaslHandshakeAuthReq(bytes, 2)
	assert.Nil(t, err)
	assert.Equal(t, 2147483642, saslHandshakeAuthReq.CorrelationId)
	assert.Equal(t, "consumer-3ac23a7f-4c36-40d9-99d4-acda47d0a48d-1", saslHandshakeAuthReq.ClientId)
	assert.Equal(t, "alice", saslHandshakeAuthReq.Username)
	assert.Equal(t, "alice", saslHandshakeAuthReq.Password)
}
