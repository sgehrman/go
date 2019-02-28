package horizonclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEffectRequestBuildUrl(t *testing.T) {

	er := EffectRequest{}
	endpoint, err := er.BuildUrl()

	// It should return valid all effects endpoint and no errors
	require.NoError(t, err)
	assert.Equal(t, "effects", endpoint)

	er = EffectRequest{LedgerId: "123"}
	endpoint, err = er.BuildUrl()

	// It should return valid ledger effects endpoint and no errors
	require.NoError(t, err)
	assert.Equal(t, "ledgers/123/effects", endpoint)

	er = EffectRequest{OperationId: "123"}
	endpoint, err = er.BuildUrl()

	// It should return valid operation effects endpoint and no errors
	require.NoError(t, err)
	assert.Equal(t, "operations/123/effects", endpoint)

	er = EffectRequest{TransactionHash: "123"}
	endpoint, err = er.BuildUrl()

	// It should return valid transaction effects endpoint and no errors
	require.NoError(t, err)
	assert.Equal(t, "transactions/123/effects", endpoint)

	er = EffectRequest{LedgerId: "123", OperationId: "789"}
	endpoint, err = er.BuildUrl()

	// error case: too many parameters for building any effect endpoint
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Invalid request. Too many parameters")
	}

}
