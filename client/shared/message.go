package shared

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
)

type Message []byte

func (msg Message) Hash() []byte {
	return accounts.TextHash(crypto.Keccak256(msg))
}

type ProtocolMessage struct {
	ProtocolID    uint8
	VotingRoundID uint32
	Message       Message
}
