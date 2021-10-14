package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApiVersion280(t *testing.T) {
	assert.Len(t, buildKfk280ApiRespVersions(), 30)
}
