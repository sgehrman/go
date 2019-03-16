package horizonclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/stellar/go/support/errors"
)

// BuildUrl creates the endpoint to be queried based on the data in the AssetRequest struct.
// If no data is set, it defaults to the build the URL for all assets
func (ar AssetRequest) BuildUrl() (endpoint string, err error) {
	endpoint = "assets"
	queryParams := addQueryParams(ar.ForAssetCode, ar.ForAssetIssuer, ar.Cursor, ar.Limit, ar.Order)
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

func (er AssetRequest) Stream(
	ctx context.Context,
	horizonURL string,
	client HTTP,
	handler func(interface{}),
) (err error) {
	surl := &StreamURL{
		horizonURL: horizonURL,
		resource:   "assets",

		Order:          er.Order,
		Cursor:         er.Cursor,
		Limit:          er.Limit,
		ForAssetCode:   er.ForAssetCode,
		ForAssetIssuer: er.ForAssetIssuer,
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
