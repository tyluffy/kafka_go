package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeIllegalHeartbeatReq(t *testing.T) {
	bytes := make([]byte, 0)
	_, err := DecodeHeartbeatReq(bytes, 0)
	assert.NotNil(t, err)
}
