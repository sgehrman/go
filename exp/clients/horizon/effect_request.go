package horizon

import (
	"fmt"
	"net/url"

	"github.com/stellar/go/support/errors"
)

// BuildUrl creates the endpoint to be queried based on the data in the EffectRequest struct.
// If no data is set, it defaults to the build the URL for all effects
func (er EffectRequest) BuildUrl() (endpoint string, err error) {

	noOfParamsSet := checkParams(er.AccountId, er.LedgerId, er.OperationId, er.TransactionHash)

	if noOfParamsSet > 1 {
		err = errors.New("Invalid request. Too many parameters")
	}

	if err != nil {
		return endpoint, err
	}

	endpoint = "effects"

	if er.AccountId != "" {
		endpoint = fmt.Sprintf(
			"accounts/%s/effects",
			er.AccountId,
		)
	}

	if er.LedgerId != "" {
		endpoint = fmt.Sprintf(
			"ledgers/%s/effects",
			er.LedgerId,
		)
	}

	if er.OperationId != "" {
		endpoint = fmt.Sprintf(
			"operations/%s/effects",
			er.OperationId,
		)
	}

	if er.TransactionHash != "" {
		endpoint = fmt.Sprintf(
			"transactions/%s/effects",
			er.TransactionHash,
		)
	}

	query := url.Values{}

	if er.Cursor != "" {
		query.Add("cursor", er.Cursor)
	}

	if er.Limit > 0 {
		query.Add("limit", string(er.Limit))
	}

	if er.Order != "" {
		query.Add("order", string(er.Order))
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
