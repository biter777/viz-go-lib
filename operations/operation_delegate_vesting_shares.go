package operations

import (
	"github.com/biter777/viz-go-lib/encoding/transaction"
	"github.com/biter777/viz-go-lib/types"
)

// DelegateVestingSharesOperation represents delegate_vesting_shares operation data.
type DelegateVestingSharesOperation struct {
	Delegator     string       `json:"delegator"`
	Delegatee     string       `json:"delegatee"`
	VestingShares *types.Asset `json:"vesting_shares"`
}

// Type function that defines the type of operation DelegateVestingSharesOperation.
func (op *DelegateVestingSharesOperation) Type() OpType {
	return TypeDelegateVestingShares
}

// Data returns the operation data DelegateVestingSharesOperation.
func (op *DelegateVestingSharesOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type DelegateVestingSharesOperation to bytes.
func (op *DelegateVestingSharesOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDelegateVestingShares.Code()))
	enc.Encode(op.Delegator)
	enc.Encode(op.Delegatee)
	enc.Encode(op.VestingShares)
	return enc.Err()
}
