package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeIllegalSaslHandshakeReq(t *testing.T) {
	bytes := make([]byte, 0)
	_, err := DecodeSaslHandshakeReq(bytes, 0)
	assert.NotNil(t, err)
}
