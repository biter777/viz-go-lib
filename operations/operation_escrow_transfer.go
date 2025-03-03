package operations

import (
	"github.com/biter777/viz-go-lib/encoding/transaction"
	"github.com/biter777/viz-go-lib/types"
)

// EscrowTransferOperation represents escrow_transfer operation data.
type EscrowTransferOperation struct {
	From                 string       `json:"from"`
	To                   string       `json:"to"`
	Agent                string       `json:"agent"`
	EscrowID             uint32       `json:"escrow_id"`
	TokenAmount          *types.Asset `json:"token_amount"`
	Fee                  *types.Asset `json:"fee"`
	RatificationDeadline *types.Time  `json:"ratification_deadline"`
	EscrowExpiration     *types.Time  `json:"escrow_expiration"`
	JSONMeta             string       `json:"json_meta"`
}

// Type function that defines the type of operation EscrowTransferOperation.
func (op *EscrowTransferOperation) Type() OpType {
	return TypeEscrowTransfer
}

// Data returns the operation data EscrowTransferOperation.
func (op *EscrowTransferOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type EscrowTransferOperation to bytes.
func (op *EscrowTransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeEscrowTransfer.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.TokenAmount)
	enc.Encode(op.EscrowID)
	enc.Encode(op.Agent)
	enc.Encode(op.Fee)
	enc.Encode(op.JSONMeta)
	enc.Encode(op.RatificationDeadline)
	enc.Encode(op.EscrowExpiration)
	return enc.Err()
}
