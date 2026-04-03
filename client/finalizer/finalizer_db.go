package finalizer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/flare-foundation/go-flare-common/pkg/database"
)

type finalizerDB interface {
	FetchTransactionsByAddressAndSelector(
		ctx context.Context, address common.Address, selector []byte, from int64, to int64,
	) ([]database.Transaction, error)
	FetchTransactionsByAddressAndSelectorFromBlockNumber(
		ctx context.Context, address common.Address, selector []byte, from int64,
	) ([]database.Transaction, error)
	FetchLogsByAddressAndTopic0(ctx context.Context, address common.Address, topic0 common.Hash, from int64, to int64) ([]database.Log, error)
}

type finalizerDBImpl struct {
	client *gorm.DB
}

// FetchTransactionsByAddressAndSelector fetches all transactions with selector sent to address in time range (from to].
func (db finalizerDBImpl) FetchTransactionsByAddressAndSelector(
	ctx context.Context, address common.Address, selector []byte, from, to int64,
) ([]database.Transaction, error) {
	return database.FetchTransactionsByAddressAndSelectorTimestamp(ctx, db.client, database.TxParams{
		ToAddress:   address,
		FunctionSel: [4]byte(selector),
		From:        from,
		To:          to,
	})
}

// FetchTransactionsByAddressAndSelectorFromBlockNumber fetches all transactions with selector sent to address in block range higher than from.
func (db finalizerDBImpl) FetchTransactionsByAddressAndSelectorFromBlockNumber(
	ctx context.Context, address common.Address, selector []byte, from int64,
) ([]database.Transaction, error) {
	return database.FetchTransactionsByAddressAndSelectorFromBlockNumber(ctx, db.client, database.TxParams{
		ToAddress:   address,
		FunctionSel: [4]byte(selector),
		From:        from,
	})
}

// FetchLogsByAddressAndTopic0 fetches all logs with topic0 emitted by address in time range (from to].
func (db finalizerDBImpl) FetchLogsByAddressAndTopic0(
	ctx context.Context, address common.Address, topic0 common.Hash, from, to int64,
) ([]database.Log, error) {
	return database.FetchLogsByAddressAndTopic0Timestamp(ctx, db.client, database.LogsParams{
		Address: address,
		Topic0:  topic0,
		From:    from,
		To:      to,
	})
}
