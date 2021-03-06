package horizonclient

import (
	"context"
	"encoding/json"

	"github.com/stellar/go/protocols/horizon/effects"
	"github.com/stellar/go/support/errors"
)

// EffectHandler is a function that is called when a new effect is received
type EffectHandler func(effects.Base)

// BuildUrl creates the endpoint to be queried based on the data in the EffectRequest struct.
// If no data is set, it defaults to the build the URL for all effects
func (er EffectRequest) BuildUrl() (endpoint string, err error) {

	surl := &StreamURL{
		horizonURL: "",
		resource:   "effects",

		ForAccount:     er.ForAccount,
		ForLedger:      er.ForLedger,
		ForOperation:   er.ForOperation,
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

func (er EffectRequest) Stream(
	ctx context.Context,
	horizonURL string,
	client HTTP,
	handler func(interface{}),
) (err error) {
	surl := &StreamURL{
		horizonURL: horizonURL,
		resource:   "effects",

		ForAccount:     er.ForAccount,
		ForLedger:      er.ForLedger,
		ForOperation:   er.ForOperation,
		ForTransaction: er.ForTransaction,
		Order:          er.Order,
		Cursor:         er.Cursor,
		Limit:          er.Limit,
	}

	return surl.Stream(ctx, client, func(data []byte) error {
		var effect effects.Base
		err = json.Unmarshal(data, &effect)
		if err != nil {
			return errors.Wrap(err, "Error unmarshaling data")
		}
		handler(effect)
		return nil
	})
}
