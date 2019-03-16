package horizonclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	hProtocol "github.com/stellar/go/protocols/horizon"

	"github.com/stellar/go/support/errors"
)

// BuildUrl creates the endpoint to be queried based on the data in the LedgerRequest struct.
// If no data is set, it defaults to the build the URL for all ledgers
func (lr LedgerRequest) BuildUrl() (endpoint string, err error) {
	endpoint = "ledgers"

	if lr.forSequence != 0 {
		endpoint = fmt.Sprintf(
			"%s/%d",
			endpoint,
			lr.forSequence,
		)
	} else {
		queryParams := addQueryParams(lr.Cursor, lr.Limit, lr.Order)
		if queryParams != "" {
			endpoint = fmt.Sprintf(
				"%s?%s",
				endpoint,
				queryParams,
			)
		}
	}

	_, err = url.Parse(endpoint)
	if err != nil {
		err = errors.Wrap(err, "failed to parse endpoint")
	}

	return endpoint, err
}

func (er LedgerRequest) Stream(
	ctx context.Context,
	horizonURL string,
	client HTTP,
	handler func(interface{}),
) (err error) {
	surl := &StreamURL{
		horizonURL: horizonURL,
		resource:   "ledgers",
		Order:      er.Order,
		Cursor:     er.Cursor,
		Limit:      er.Limit,
	}

	return surl.Stream(ctx, client, func(data []byte) error {
		var ledger hProtocol.Ledger
		err = json.Unmarshal(data, &ledger)

		if err != nil {
			return errors.Wrap(err, "Error unmarshaling data")
		}

		handler(ledger)
		return nil
	})
}
