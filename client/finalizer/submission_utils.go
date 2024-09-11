package finalizer

import (
	"github.com/ethereum/go-ethereum/common"
)

type submissionContractClient struct {
	address common.Address
}

func NewSubmissionContractClient(address common.Address) *submissionContractClient {
	return &submissionContractClient{
		address: address,
	}
}
