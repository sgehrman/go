package horizon

import "github.com/stellar/go/support/errors"

func (c *Client) AccountDetail(request AccountRequest) (account Account, err error) {

	if request.AccountId == "" {
		err = errors.New("No account ID provided")
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

	err = decodeResponse(resp, &account)
	return

}

func (c *Client) AccountData(request AccountRequest) (AccountData AccountData, err error) {

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

	err = decodeResponse(resp, &AccountData)
	return

}
