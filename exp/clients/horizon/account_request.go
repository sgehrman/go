package horizon

import (
	"fmt"
	"net/url"

	"github.com/stellar/go/support/errors"
)

// BuildUrl creates the endpoint to be queried based on the data in the AccountRequest struct.
// If only AccountId is present, then the endpoint for account details is returned.
// If both AccounId and DataKey are present, then the endpoint for getting account data is returned
func (ar AccountRequest) BuildUrl() (endpoint string, err error) {

	// to do:  check if ar.accountId  is a valid public key. I wonder if this is the right place for this.

	noOfParamsSet := checkParams(ar.DataKey, ar.AccountId)

	if noOfParamsSet >= 1 && ar.AccountId == "" {
		err = errors.New("Invalid request. Too few parameters")
	}

	if noOfParamsSet <= 0 {
		err = errors.New("Invalid request. No parameters")
	}

	if err != nil {
		return endpoint, err
	}

	if ar.DataKey != "" && ar.AccountId != "" {
		endpoint = fmt.Sprintf(
			"accounts/%s/data/%s",
			ar.AccountId,
			ar.DataKey,
		)
	} else if ar.AccountId != "" {
		endpoint = fmt.Sprintf(
			"accounts/%s",
			ar.AccountId,
		)
	}

	query := url.Values{}

	if ar.Cursor != "" {
		query.Add("cursor", ar.Cursor)
	}

	if ar.Limit > 0 {
		query.Add("limit", string(ar.Limit))
	}

	if ar.Order != "" {
		query.Add("order", string(ar.Order))
	}

	endpoint = fmt.Sprintf(
		"%s",
		endpoint,
	)

	queryParams := query.Encode()

	if queryParams != "" {
		endpoint = fmt.Sprintf(
			"%s?%s",
			endpoint,
			queryParams,
		)
	}

	_, err = url.Parse(endpoint)
	if err != nil {
		err = errors.Wrap(err, "failed to parse endpoint")
	}

	return endpoint, err

}
