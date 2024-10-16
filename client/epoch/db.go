package epoch

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/flare-foundation/go-flare-common/pkg/database"
)

type epochClientDB interface {
	FetchLogsByAddressAndTopic0(common.Address, common.Hash, int64, int64) ([]database.Log, error)
}

type epochClientDBGorm struct {
	db *gorm.DB
}

func (g epochClientDBGorm) FetchLogsByAddressAndTopic0(
	address common.Address, topic0 common.Hash, fromBlock int64, toBlock int64,
) ([]database.Log, error) {
	return database.FetchLogsByAddressAndTopic0BlockNumber(context.TODO(), g.db, database.LogsParams{
		Address: address,
		Topic0:  [32]byte{},
		From:    fromBlock,
		To:      toBlock,
	})
}
