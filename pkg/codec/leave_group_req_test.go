package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeLeaveGroupReqV4(t *testing.T) {
	bytes := testHex2Bytes(t, "0000000c002f636f6e73756d65722d38646437623936622d366239342d346139622d623263632d3363623538393863396364662d31002538646437623936622d366239342d346139622d623263632d3363623538393863396364660255636f6e73756d65722d38646437623936622d366239342d346139622d623263632d3363623538393863396364662d312d34333361636236612d653665632d343561612d623738642d366132343963666630376663000000")
	leaveGroupReq, err := DecodeLeaveGroupReq(bytes, 4)
	assert.Nil(t, err)
	assert.Equal(t, 12, leaveGroupReq.CorrelationId)
	assert.Equal(t, "consumer-8dd7b96b-6b94-4a9b-b2cc-3cb5898c9cdf-1", leaveGroupReq.ClientId)
	assert.Equal(t, "8dd7b96b-6b94-4a9b-b2cc-3cb5898c9cdf", leaveGroupReq.GroupId)
	assert.Len(t, leaveGroupReq.Members, 1)
	leaveGroupMember := leaveGroupReq.Members[0]
	assert.Equal(t, "consumer-8dd7b96b-6b94-4a9b-b2cc-3cb5898c9cdf-1-433acb6a-e6ec-45aa-b78d-6a249cff07fc", leaveGroupMember.MemberId)
	assert.Nil(t, leaveGroupMember.GroupInstanceId)
}
