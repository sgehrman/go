package horizonclient

import (
	"context"
	"encoding/json"

	"github.com/stellar/go/protocols/horizon/operations"
	"github.com/stellar/go/support/errors"
)

// OperationHandler is a function that is called when a new effect is received
type OperationHandler func(operations.Base)

// BuildUrl creates the endpoint to be queried based on the data in the EffectRequest struct.
// If no data is set, it defaults to the build the URL for all effects
func (er OperationRequest) BuildUrl() (endpoint string, err error) {

	surl := &StreamURL{
		horizonURL: "",
		resource:   "operations",

		ForAccount:     er.ForAccount,
		ForLedger:      er.ForLedger,
		ForTransaction: er.ForTransaction,
		Order:          er.Order,
		Cursor:         er.Cursor,
		Limit:          er.Limit,
	}

	res, err := surl.BuildUrl()

	if err != nil {
		return endpoint, err
	}

	endpoint = res

	return endpoint, err
}

func (er OperationRequest) Stream(
	ctx context.Context,
	horizonURL string,
	client HTTP,
	handler func(interface{}),
) (err error) {
	surl := &StreamURL{
		horizonURL: horizonURL,
		resource:   "operations",

		ForAccount:     er.ForAccount,
		ForLedger:      er.ForLedger,
		ForTransaction: er.ForTransaction,
		Order:          er.Order,
		Cursor:         er.Cursor,
		Limit:          er.Limit,
	}

	return surl.Stream(ctx, client, func(data []byte) error {
		var operation operations.Base

		err = json.Unmarshal(data, &operation)
		if err != nil {
			return errors.Wrap(err, "Error unmarshaling data")
		}
		handler(operation)
		return nil
	})
}
