package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_readSaslAuthBytes(t *testing.T) {
	bytes := testHex2Bytes(t, "00616c69636500707764")
	username, pwd := readSaslAuthBytes(bytes, 0)
	assert.Equal(t, "alice", username)
	assert.Equal(t, "pwd", pwd)
}
