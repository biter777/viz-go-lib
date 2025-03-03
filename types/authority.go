package types

import (
	"github.com/biter777/viz-go-lib/encoding/transaction"
)

// Authority is an additional structure used by other structures.
type Authority struct {
	WeightThreshold uint32         `json:"weight_threshold"`
	AccountAuths    StringInt64Map `json:"account_auths"`
	KeyAuths        StringInt64Map `json:"key_auths"`
}

// MarshalTransaction is a function of converting type Authority to bytes.
func (auth *Authority) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeNumber(uint32(auth.WeightThreshold))
	// encode AccountAuths as map[string]uint16
	enc.EncodeUVarint(uint64(len(auth.AccountAuths)))
	for k, v := range auth.AccountAuths {
		enc.EncodeString(k)
		enc.EncodeNumber(uint16(v))
	}
	// encode KeyAuths as map[PubKey]uint16
	enc.EncodeUVarint(uint64(len(auth.KeyAuths)))
	for k, v := range auth.KeyAuths {
		enc.EncodePubKey(k)
		enc.EncodeNumber(uint16(v))
	}
	return enc.Err()
}
