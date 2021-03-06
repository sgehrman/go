package horizonclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOperationRequestBuildUrl(t *testing.T) {
	er := OperationRequest{}
	endpoint, err := er.BuildUrl()

	// It should return valid all operations endpoint and no errors
	require.NoError(t, err)
	assert.Equal(t, "operations", endpoint)

	er = OperationRequest{ForAccount: "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU"}
	endpoint, err = er.BuildUrl()

	// It should return valid account operations endpoint and no errors
	require.NoError(t, err)
	assert.Equal(t, "accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU/operations", endpoint)

	er = OperationRequest{ForLedger: "123"}
	endpoint, err = er.BuildUrl()

	// It should return valid ledger operations endpoint and no errors
	require.NoError(t, err)
	assert.Equal(t, "ledgers/123/operations", endpoint)

	er = OperationRequest{ForTransaction: "123"}
	endpoint, err = er.BuildUrl()

	// It should return valid transaction operations endpoint and no errors
	require.NoError(t, err)
	assert.Equal(t, "transactions/123/operations", endpoint)

	er = OperationRequest{Cursor: "123456", Limit: 30, Order: OrderAsc}
	endpoint, err = er.BuildUrl()
	// It should return valid all operations endpoint with query params and no errors
	require.NoError(t, err)
	assert.Equal(t, "operations?cursor=123456&limit=30&order=asc", endpoint)

}
