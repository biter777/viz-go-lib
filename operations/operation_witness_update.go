package operations

import (
	"github.com/biter777/viz-go-lib/encoding/transaction"
)

// WitnessUpdateOperation represents witness_update operation data.
type WitnessUpdateOperation struct {
	Owner           string `json:"owner"`
	URL             string `json:"url"`
	BlockSigningKey string `json:"block_signing_key"`
}

// Type function that defines the type of operation WitnessUpdateOperation.
func (op *WitnessUpdateOperation) Type() OpType {
	return TypeWitnessUpdate
}

// Data returns the operation data WitnessUpdateOperation.
func (op *WitnessUpdateOperation) Data() interface{} {
	return op
}

// MarshalTransaction is a function of converting type WitnessUpdateOperation to bytes.
func (op *WitnessUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeWitnessUpdate.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.URL)
	enc.EncodePubKey(op.BlockSigningKey)
	return enc.Err()
}
