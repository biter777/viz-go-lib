package operations

import (
	"github.com/biter777/viz-go-lib/encoding/transaction"
	"github.com/biter777/viz-go-lib/types"
)

// EscrowReleaseOperation represents escrow_release operation data.
type EscrowReleaseOperation struct {
	From        string       `json:"from"`
	To          string       `json:"to"`
	Agent       string       `json:"agent"`
	Who         string       `json:"who"`
	Receiver    string       `json:"receiver"`
	EscrowID    uint32       `json:"escrow_id"`
	TokenAmount *types.Asset `json:"token_amount"`
}

// Type function that defines the type of operation EscrowReleaseOperation.
func (op *EscrowReleaseOperation) Type() OpType {
	return TypeEscrowRelease
}

// Data returns the operation data EscrowReleaseOperation.
func (op *EscrowReleaseOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type EscrowReleaseOperation to bytes.
func (op *EscrowReleaseOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeEscrowRelease.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Agent)
	enc.Encode(op.Who)
	enc.Encode(op.Receiver)
	enc.Encode(op.EscrowID)
	enc.Encode(op.TokenAmount)
	return enc.Err()
}
