// package horizon is an experimental horizon client that provides access to the horizon server
package horizon

import (
	"errors"
	"net/http"
	"net/url"
)

type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

// Error struct contains the problem returned by Horizon
type Error struct {
	Response *http.Response
	Problem  Problem
}

var (
	// ErrResultCodesNotPopulated is the error returned from a call to
	// ResultCodes() against a `Problem` value that doesn't have the
	// "result_codes" extra field populated when it is expected to be.
	ErrResultCodesNotPopulated = errors.New("result_codes not populated")

	// ErrEnvelopeNotPopulated is the error returned from a call to
	// Envelope() against a `Problem` value that doesn't have the
	// "envelope_xdr" extra field populated when it is expected to be.
	ErrEnvelopeNotPopulated = errors.New("envelope_xdr not populated")

	// ErrResultNotPopulated is the error returned from a call to
	// Result() against a `Problem` value that doesn't have the
	// "result_xdr" extra field populated when it is expected to be.
	ErrResultNotPopulated = errors.New("result_xdr not populated")
)

// HTTP represents the HTTP client that a horizon client uses to communicate
type HTTP interface {
	Do(req *http.Request) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

// Client struct contains data for creating an horizon client that connects to the stellar network
type Client struct {
	HorizonURL string
	HTTP       HTTP
}

// ClientInterface contains methods implemented by the horizon client
type ClientInterface interface {
	AccountDetail(request AccountRequest) (Account, error)
	AccountData(request AccountRequest) (AccountData, error)
	AllEffects(request EffectRequest) (EffectsPage, error)
	LedgerEffects(request EffectRequest) (EffectsPage, error)
	OperationEffects(request EffectRequest) (EffectsPage, error)
	TransactionEffects(request EffectRequest) (EffectsPage, error)
}

// DefaultTestNetClient is a default client to connect to test network
var DefaultTestNetClient = &Client{
	HorizonURL: "https://horizon-testnet.stellar.org",
	HTTP:       http.DefaultClient,
}

// DefaultPublicNetClient is a default client to connect to public network
var DefaultPublicNetClient = &Client{
	HorizonURL: "https://horizon.stellar.org",
	HTTP:       http.DefaultClient,
}

type HorizonRequest interface {
	BuildUrl() (string, error)
}

// AccountRequest struct contains data for making requests to the accounts endpoint of an horizon server
type AccountRequest struct {
	AccountId string
	DataKey   string

	Order  Order
	Cursor string
	Limit  int
}

type EffectRequest struct {
	AccountId       string
	LedgerId        string
	OperationId     string
	TransactionHash string

	Order  Order
	Cursor string
	Limit  int
}
