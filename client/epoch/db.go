package epoch

import (
	"flare-fsc/database"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type epochClientDB interface {
	FetchLogsByAddressAndTopic0(common.Address, string, int64, int64) ([]database.Log, error)
}

type epochClientDBGorm struct {
	db *gorm.DB
}

func (g epochClientDBGorm) FetchLogsByAddressAndTopic0(
	address common.Address, topic0 string, fromBlock int64, toBlock int64,
) ([]database.Log, error) {
	return database.FetchLogsByAddressAndTopic0(g.db, address.Hex(), topic0, fromBlock, toBlock)
}
