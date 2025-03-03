package operations

import (
	"github.com/biter777/viz-go-lib/types"
)

// FillVestingWithdrawOperation represents fill_vesting_withdraw operation data.
type FillVestingWithdrawOperation struct {
	FromAccount string      `json:"from_account"`
	ToAccount   string      `json:"to_account"`
	Withdrawn   types.Asset `json:"withdrawn"`
	Deposited   types.Asset `json:"deposited"`
}

// Type function that defines the type of operation FillVestingWithdrawOperation.
func (op *FillVestingWithdrawOperation) Type() OpType {
	return TypeFillVestingWithdraw
}

// Data returns the operation data FillVestingWithdrawOperation.
func (op *FillVestingWithdrawOperation) Data() interface{} {
	return op
}
