package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeApiVersionReqV3(t *testing.T) {
	bytes := testHex2Bytes(t, "00000001002f636f6e73756d65722d37336664633964612d306439322d346537622d613761372d6563323636663637633137312d3100126170616368652d6b61666b612d6a61766106322e342e3000")
	apiReq, err := DecodeApiReq(bytes, 3)
	assert.Nil(t, err)
	assert.Equal(t, 1, apiReq.CorrelationId)
	assert.Equal(t, "consumer-73fdc9da-0d92-4e7b-a7a7-ec266f67c171-1", apiReq.ClientId)
	assert.Equal(t, "apache-kafka-java", apiReq.ClientSoftwareName)
	assert.Equal(t, "2.4.0", apiReq.ClientSoftwareVersion)
}
