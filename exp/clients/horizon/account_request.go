package horizon

import (
	"fmt"
	"net/url"

	"github.com/stellar/go/support/errors"
)

func (ar AccountRequest) BuildUrl(rtype string) (endpoint string, err error) {

	switch rtype {
	case "account":
		endpoint = fmt.Sprintf(
			"accounts/%s",
			ar.AccountId,
		)
		break

	case "data":
		endpoint = fmt.Sprintf(
			"accounts/%s/data/%s",
			ar.AccountId,
			ar.DataKey,
		)
		break

	default:
		err = errors.Wrap(err, "Invalid request type, expected 'id' or 'data'")

	}

	if err != nil {
		return endpoint, err
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
		"%s?%s",
		endpoint,
		query.Encode(),
	)

	_, err = url.Parse(endpoint)
	if err != nil {
		err = errors.Wrap(err, "failed to parse endpoint")
	}

	return endpoint, err

}
