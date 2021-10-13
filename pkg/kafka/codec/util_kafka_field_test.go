package codec

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_readSaslAuthBytes(t *testing.T) {
	bytes, err := hex.DecodeString("00616c69636500707764")
	assert.Nil(t, err)
	username, pwd := readSaslAuthBytes(bytes, 0)
	assert.Equal(t, "alice", username)
	assert.Equal(t, "pwd", pwd)
}
