package finalizer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/database"
)

type finalizerDB interface {
	FetchTransactionsByAddressAndSelector(
		address common.Address, selector []byte, from int64, to int64,
	) ([]database.Transaction, error)
	FetchTransactionsByAddressAndSelectorFromBlockNumber(
		address common.Address, selector []byte, from int64,
	) ([]database.Transaction, error)
	FetchLogsByAddressAndTopic0(address common.Address, topic0 common.Hash, from int64, to int64) ([]database.Log, error)
}

type finalizerDBImpl struct {
	client *gorm.DB
}

func (db finalizerDBImpl) FetchTransactionsByAddressAndSelector(
	address common.Address, selector []byte, from, to int64,
) ([]database.Transaction, error) {
	return database.FetchTransactionsByAddressAndSelectorTimestamp(context.TODO(), db.client, database.TxParams{
		ToAddress:   address,
		FunctionSel: [4]byte(selector),
		From:        from,
		To:          to,
	})
}

func (db finalizerDBImpl) FetchTransactionsByAddressAndSelectorFromBlockNumber(
	address common.Address, selector []byte, from int64,
) ([]database.Transaction, error) {
	return database.FetchTransactionsByAddressAndSelectorFromBlockNumber(context.TODO(), db.client, database.TxParams{
		ToAddress:   address,
		FunctionSel: [4]byte(selector),
		From:        from,
	})
}

func (db finalizerDBImpl) FetchLogsByAddressAndTopic0(
	address common.Address, topic0 common.Hash, from, to int64,
) ([]database.Log, error) {
	return database.FetchLogsByAddressAndTopic0Timestamp(context.TODO(), db.client, database.LogsParams{
		Address: address,
		Topic0:  topic0,
		From:    from,
		To:      to,
	})
}
