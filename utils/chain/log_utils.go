package chain

import "github.com/ethereum/go-ethereum/common"

func ParseTopic(topic string) common.Hash {
	if topic == "NULL" {
		return common.Hash{}
	}
	return common.HexToHash(topic)
}
