package horizonclient

import (
	"github.com/stellar/go/support/errors"
)

func SendRequest(hr HorizonRequest, c Client, a interface{}) (err error) {

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

	err = SendRequest(request, *c, &account)
	return

}

func (c *Client) AccountData(request AccountRequest) (accountData AccountData, err error) {

	if request.AccountId == "" || request.DataKey == "" {
		err = errors.New("Too few parameters")
	}

	if err != nil {
		return
	}

	endpoint, err := request.BuildUrl()
	if err != nil {
		return
	}

	resp, err := c.HTTP.Get(c.HorizonURL + endpoint)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &accountData)
	return

}

func (c *Client) AllEffects(request EffectRequest) (effects EffectsPage, err error) {
	if request.AccountId != "" || request.LedgerId != "" || request.OperationId != "" || request.TransactionHash != "" {
		err = errors.New("Too many parameters")
	}

	if err != nil {
		return
	}
	endpoint, err := request.BuildUrl()
	if err != nil {
		return
	}

	resp, err := c.HTTP.Get(c.HorizonURL + endpoint)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &effects)
	return

}

func (c *Client) LedgerEffects(request EffectRequest) (effects EffectsPage, err error) {
	if request.LedgerId == "" {
		err = errors.New("Ledger Id is not provided.")
	}

	if err != nil {
		return
	}
	endpoint, err := request.BuildUrl()
	if err != nil {
		return
	}

	resp, err := c.HTTP.Get(c.HorizonURL + endpoint)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &effects)
	return

}

func (c *Client) OperationEffects(request EffectRequest) (effects EffectsPage, err error) {
	if request.OperationId == "" {
		err = errors.New("Operation Id is not provided.")
	}

	if err != nil {
		return
	}
	endpoint, err := request.BuildUrl()
	if err != nil {
		return
	}

	resp, err := c.HTTP.Get(c.HorizonURL + endpoint)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &effects)
	return

}

func (c *Client) TransactionEffects(request EffectRequest) (effects EffectsPage, err error) {
	if request.TransactionHash == "" {
		err = errors.New("Transaction Hash is not provided.")
	}

	if err != nil {
		return
	}
	endpoint, err := request.BuildUrl()
	if err != nil {
		return
	}

	resp, err := c.HTTP.Get(c.HorizonURL + endpoint)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &effects)
	return

}
