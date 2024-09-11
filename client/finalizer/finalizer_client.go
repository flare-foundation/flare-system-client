package finalizer

import (
	"encoding/hex"
	"flare-fsc/database"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type finalizerDB interface {
	FetchTransactionsByAddressAndSelector(
		common.Address, []byte, int64, int64,
	) ([]database.Transaction, error)
	FetchLogsByAddressAndTopic0(common.Address, string, int64, int64) ([]database.Log, error)
}

type finalizerDBImpl struct {
	client *gorm.DB
}

func (db finalizerDBImpl) FetchTransactionsByAddressAndSelector(
	address common.Address, selector []byte, from, to int64,
) ([]database.Transaction, error) {
	hexSelector := hex.EncodeToString(selector)
	return database.FetchTransactionsByAddressAndSelector(db.client, address.Hex(), hexSelector, from, to)
}

func (db finalizerDBImpl) FetchLogsByAddressAndTopic0(
	address common.Address, topic0 string, from, to int64,
) ([]database.Log, error) {
	return database.FetchLogsByAddressAndTopic0(db.client, address.Hex(), topic0, from, to)
}
