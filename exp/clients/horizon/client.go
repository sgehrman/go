package horizon

func (c *Client) AccountDetail(request AccountRequest) (account Account, err error) {

	endpoint, err := request.BuildUrl("account")
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

	endpoint, err := request.BuildUrl("data")
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
