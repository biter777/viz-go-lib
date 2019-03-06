package api

import (
	"encoding/json"

	"github.com/VIZ-Blockchain/viz-go-lib/types"
	"github.com/pkg/errors"
)

//account_history

//GetAccountHistory api request get_account_history
func (api *API) GetAccountHistory(account string, from uint64, limit uint32) ([]*types.OperationObject, error) {
	if limit > 10000 {
		return nil, errors.New("account_history: get_account_history -> limit must not exceed 10000")
	}
	if from == 0 {
		return nil, errors.New("account_history: get_account_history -> from can not have the value 0")
	}
	if from < uint64(limit) && !(from < uint64(0)) {
		return nil, errors.New("account_history: get_account_history -> from must be greater than or equal to the limit")
	}
	var raw json.RawMessage
	err := api.call("account_history", "get_account_history", []interface{}{account, from, limit}, &raw)
	if err != nil {
		return nil, err
	}
	var tmp1 [][]interface{}
	if err := json.Unmarshal([]byte(raw), &tmp1); err != nil {
		return nil, err
	}
	var resp []*types.OperationObject
	for _, v := range tmp1 {
		byteData, errm := json.Marshal(&v[1])
		if errm != nil {
			return nil, errm
		}
		var tmp *types.OperationObject
		if err := json.Unmarshal(byteData, &tmp); err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}
