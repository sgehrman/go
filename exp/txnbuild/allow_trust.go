package txnbuild

import (
	"github.com/stellar/go/amount"
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/xdr"
)

// AllowTrust represents the Stellar allow trust operation. See
// https://www.stellar.org/developers/guides/concepts/list-of-operations.html
type AllowTrust struct {
	Trustor   string
	Type      *Asset
	Authorize bool
	xdrOp     xdr.AllowTrustOp
}

// BuildXDR for AllowTrust returns a fully configured XDR Operation.
func (at *AllowTrust) BuildXDR() (xdr.Operation, error) {
	var err error
	err = at.xdrOp.Trustor.SetAddress(at.Trustor)
	if err != nil {
		return xdr.Operation{}, errors.Wrap(err, "Failed to set trustor address")
	}

	if at.Type.IsNative() {
		return xdr.Operation{}, errors.New("Trustline doesn't exist for a native (XLM) asset")
	}

	// TODO: AllowTrust has a special asset op type. Need to map to it
	at.xdrOp.Asset, err = ct.Line.ToXDR()
	if err != nil {
		return xdr.Operation{}, errors.Wrap(err, "Can't convert asset for trustline to XDR")
	}

	ct.xdrOp.Limit, err = amount.Parse(ct.Limit)
	if err != nil {
		return xdr.Operation{}, errors.Wrap(err, "Failed to parse limit amount")
	}

	opType := xdr.OperationTypeChangeTrust
	body, err := xdr.NewOperationBody(opType, ct.xdrOp)
	if err != nil {
		return xdr.Operation{}, errors.Wrap(err, "Failed to build XDR OperationBody")
	}

	return xdr.Operation{Body: body}, nil
}
