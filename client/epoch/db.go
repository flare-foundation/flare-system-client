package epoch

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/flare-foundation/go-flare-common/pkg/database"
)

type epochClientDB interface {
	FetchLogsByAddressAndTopic0Timestamp(common.Address, common.Hash, int64, int64) ([]database.Log, error)
}

type epochClientDBGorm struct {
	db *gorm.DB
}

func (g epochClientDBGorm) FetchLogsByAddressAndTopic0Timestamp(
	address common.Address, topic0 common.Hash, fromTimestamp int64, toTimestamp int64,
) ([]database.Log, error) {
	return database.FetchLogsByAddressAndTopic0Timestamp(context.TODO(), g.db, database.LogsParams{
		Address: address,
		Topic0:  topic0,
		From:    fromTimestamp,
		To:      toTimestamp,
	})
}
