package horizonclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/stellar/go/support/errors"
)

// BuildUrl creates the endpoint to be queried based on the data in the OfferRequest struct.
func (or OfferRequest) BuildUrl() (endpoint string, err error) {
	endpoint = fmt.Sprintf(
		"accounts/%s/offers",
		or.ForAccount,
	)

	queryParams := addQueryParams(or.Cursor, or.Limit, or.Order)
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

func (er OfferRequest) Stream(
	ctx context.Context,
	horizonURL string,
	client HTTP,
	handler func(interface{}),
) (err error) {
	surl := &StreamURL{
		horizonURL: horizonURL,
		resource:   "offers",

		ForAccount: er.ForAccount,
		Order:      er.Order,
		Cursor:     er.Cursor,
		Limit:      er.Limit,
	}

	return surl.Stream(ctx, client, func(data []byte) error {
		var objmap map[string]*json.RawMessage

		err = json.Unmarshal(data, &objmap)
		if err != nil {
			return errors.Wrap(err, "Error unmarshaling data")
		}
		handler(objmap)
		return nil
	})
}
