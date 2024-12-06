package chain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func EventIDFromMetadata(metaData *bind.MetaData, eventName string) (common.Hash, error) {
	abi, err := metaData.GetAbi()
	if err != nil {
		return common.Hash{}, err
	}
	id := abi.Events[eventName].ID
	return id, nil
}

func FunctionSelector(signature string) (selector [4]byte) {
	copy(selector[:], crypto.Keccak256([]byte(signature)))
	return
}
