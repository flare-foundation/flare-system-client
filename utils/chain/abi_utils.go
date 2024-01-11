package chain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func EventIDFromMetadata(metaData *bind.MetaData, eventName string) (string, error) {
	abi, err := metaData.GetAbi()
	if err != nil {
		return "", err
	}
	id := abi.Events[eventName].ID
	return id.String(), nil
}
