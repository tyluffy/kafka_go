package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodecLogger(t *testing.T) {
	assert.NotNil(t, Codec())
}

func TestNetworkLogger(t *testing.T) {
	assert.NotNil(t, Network())
}
