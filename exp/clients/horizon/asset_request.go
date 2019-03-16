package horizonclient

import (
	"context"
	"encoding/json"

	"github.com/stellar/go/protocols/horizon/base"

	"github.com/stellar/go/support/errors"
)

// BuildUrl creates the endpoint to be queried based on the data in the AssetRequest struct.
// If no data is set, it defaults to the build the URL for all assets
func (ar AssetRequest) BuildUrl() (endpoint string, err error) {
	surl := &StreamURL{
		horizonURL: "",
		resource:   "assets",

		ForAssetCode:   ar.ForAssetCode,
		ForAssetIssuer: ar.ForAssetIssuer,
		Order:          ar.Order,
		Cursor:         ar.Cursor,
		Limit:          ar.Limit,
	}

	res, err := surl.BuildUrl()

	if err != nil {
		return endpoint, err
	}

	endpoint = res

	return endpoint, err
}

func (ar AssetRequest) Stream(
	ctx context.Context,
	horizonURL string,
	client HTTP,
	handler func(interface{}),
) (err error) {
	surl := &StreamURL{
		horizonURL: horizonURL,
		resource:   "assets",

		Order:          ar.Order,
		Cursor:         ar.Cursor,
		Limit:          ar.Limit,
		ForAssetCode:   ar.ForAssetCode,
		ForAssetIssuer: ar.ForAssetIssuer,
	}

	return surl.Stream(ctx, client, func(data []byte) error {
		var asset base.Asset

		err = json.Unmarshal(data, &asset)
		if err != nil {
			return errors.Wrap(err, "Error unmarshaling data")
		}
		handler(asset)
		return nil
	})
}
