package horizon

import (
	"testing"

	"github.com/stellar/go/support/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestAccountDetail(t *testing.T) {
	hmock := httptest.NewClient()
	client := &Client{
		HorizonURL: "https://localhost/",
		HTTP:       hmock,
	}

	accountRequest := AccountRequest{AccountId: "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU"}

	// happy path
	hmock.On(
		"GET",
		"https://localhost/accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU",
	).ReturnString(200, accountResponse)

	account, err := client.AccountDetail(accountRequest)
	if assert.NoError(t, err) {
		assert.Equal(t, account.ID, "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU")
		assert.Equal(t, account.PT, "1")
		assert.Equal(t, account.Signers[0].Key, "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU")
		assert.Equal(t, account.Signers[0].Type, "ed25519_public_key")
		assert.Equal(t, account.Data["test"], "dGVzdA==")
		balance, err := account.GetNativeBalance()
		assert.Nil(t, err)
		assert.Equal(t, balance, "9999.9999900")
	}

	// failure response
	hmock.On(
		"GET",
		"https://localhost/accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU",
	).ReturnString(404, notFoundResponse)

	_, err = client.AccountDetail(accountRequest)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Horizon error")
		horizonError, ok := err.(*Error)
		assert.Equal(t, ok, true)
		assert.Equal(t, horizonError.Problem.Title, "Resource Missing")
	}

	// connection error
	hmock.On(
		"GET",
		"https://localhost/accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU",
	).ReturnError("http.Client error")

	_, err = client.AccountDetail(accountRequest)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "http.Client error")
		_, ok := err.(*Error)
		assert.Equal(t, ok, false)
	}
}

func TestAccountData(t *testing.T) {
	hmock := httptest.NewClient()
	client := &Client{
		HorizonURL: "https://localhost/",
		HTTP:       hmock,
	}

	accountRequest := AccountRequest{AccountId: "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU", DataKey: "test"}

	// happy path
	hmock.On(
		"GET",
		"https://localhost/accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU/data/test",
	).ReturnString(200, accountData)

	data, err := client.AccountData(accountRequest)
	if assert.NoError(t, err) {
		assert.Equal(t, data.Value, "dGVzdA==")
	}

}

var accountResponse = `{
  "_links": {
    "self": {
      "href": "https://horizon-testnet.stellar.org/accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU"
    },
    "transactions": {
      "href": "https://horizon-testnet.stellar.org/accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU/transactions{?cursor,limit,order}",
      "templated": true
    },
    "operations": {
      "href": "https://horizon-testnet.stellar.org/accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU/operations{?cursor,limit,order}",
      "templated": true
    },
    "payments": {
      "href": "https://horizon-testnet.stellar.org/accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU/payments{?cursor,limit,order}",
      "templated": true
    },
    "effects": {
      "href": "https://horizon-testnet.stellar.org/accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU/effects{?cursor,limit,order}",
      "templated": true
    },
    "offers": {
      "href": "https://horizon-testnet.stellar.org/accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU/offers{?cursor,limit,order}",
      "templated": true
    },
    "trades": {
      "href": "https://horizon-testnet.stellar.org/accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU/trades{?cursor,limit,order}",
      "templated": true
    },
    "data": {
      "href": "https://horizon-testnet.stellar.org/accounts/GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU/data/{key}",
      "templated": true
    }
  },
  "id": "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU",
  "paging_token": "1",
  "account_id": "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU",
  "sequence": "9865509814140929",
  "subentry_count": 1,
  "thresholds": {
    "low_threshold": 0,
    "med_threshold": 0,
    "high_threshold": 0
  },
  "flags": {
    "auth_required": false,
    "auth_revocable": false,
    "auth_immutable": false
  },
  "balances": [
    {
      "balance": "9999.9999900",
      "buying_liabilities": "0.0000000",
      "selling_liabilities": "0.0000000",
      "asset_type": "native"
    }
  ],
  "signers": [
    {
      "public_key": "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU",
      "weight": 1,
      "key": "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU",
      "type": "ed25519_public_key"
    }
  ],
  "data": {
    "test": "dGVzdA=="
  }
}`

var notFoundResponse = `{
  "type": "https://stellar.org/horizon-errors/not_found",
  "title": "Resource Missing",
  "status": 404,
  "detail": "The resource at the url requested was not found.  This is usually occurs for one of two reasons:  The url requested is not valid, or no data in our database could be found with the parameters provided.",
  "instance": "horizon-live-001/61KdRW8tKi-18408110"
}`

var accountData = `{
  "value": "dGVzdA=="
}`
