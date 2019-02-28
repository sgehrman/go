package horizonclient

import (
	"github.com/stellar/go/support/errors"
)

func sendRequest(hr HorizonRequest, c Client, a interface{}) (err error) {
	endpoint, err := hr.BuildUrl()
	if err != nil {
		return
	}

	resp, err := c.HTTP.Get(c.HorizonURL + endpoint)
	if err != nil {

		return
	}

	err = decodeResponse(resp, &a)
	return
}

func (c *Client) AccountDetail(request AccountRequest) (account Account, err error) {
	if request.AccountId == "" {
		err = errors.New("No account ID provided")
	}

	if err != nil {
		return
	}

	err = sendRequest(request, *c, &account)
	return
}

func (c *Client) AccountData(request AccountRequest) (accountData AccountData, err error) {
	if request.AccountId == "" || request.DataKey == "" {
		err = errors.New("Too few parameters")
	}

	if err != nil {
		return
	}

	err = sendRequest(request, *c, &accountData)
	return
}

func (c *Client) Effects(request EffectRequest) (effects EffectsPage, err error) {
	err = sendRequest(request, *c, &effects)
	return
}
