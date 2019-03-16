package horizonclient

import (
	"context"
	"encoding/json"

	"github.com/stellar/go/support/errors"
)

// BuildUrl creates the endpoint to be queried based on the data in the OfferRequest struct.
func (or OfferRequest) BuildUrl() (endpoint string, err error) {
	surl := &StreamURL{
		horizonURL: "",
		resource:   "offers",

		ForAccount: or.ForAccount,
		Order:      or.Order,
		Cursor:     or.Cursor,
		Limit:      or.Limit,
	}

	res, err := surl.BuildUrl()

	if err != nil {
		return endpoint, err
	}

	endpoint = res

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
