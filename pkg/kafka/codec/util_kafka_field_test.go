package codec

import (
	"encoding/hex"
	"testing"
)

func Test_readSaslAuthBytes(t *testing.T) {
	bytes, _ := hex.DecodeString("00616c69636500707764")
	username, pwd := readSaslAuthBytes(bytes, 0)
	if username != "alice" {
		t.Errorf("user name is not alice")
	}
	if pwd != "pwd" {
		t.Errorf("password is not pwd")
	}
}
