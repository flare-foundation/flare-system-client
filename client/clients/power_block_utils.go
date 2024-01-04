package clients

import (
	"encoding/hex"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/chain"
	"flare-tlc/utils/contracts/system"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

type votePowerBlockSelectedListener struct {
	C <-chan *system.FlareSystemManagerVotePowerBlockSelected

	epoch   *utils.Epoch
	address string
	topic0  string

	systemManagerFilterer *system.FlareSystemManagerFilterer
	db                    *gorm.DB
	mockableTime          utils.TimeProvider
}

func NewVotePowerBlockSelectedListener(
	db *gorm.DB,
	systemManagerFilterer *system.FlareSystemManagerFilterer,
	epoch *utils.Epoch,
	address string,
	topic0 string,
) *votePowerBlockSelectedListener {
	listener := &votePowerBlockSelectedListener{
		epoch:                 epoch,
		address:               address,
		topic0:                topic0,
		systemManagerFilterer: systemManagerFilterer,
		db:                    db,
		mockableTime:          &utils.RealTimeProvider{},
	}
	listener.C = listener.votePowerBlockChannel()
	return listener
}

func (c *votePowerBlockSelectedListener) votePowerBlockChannel() <-chan *system.FlareSystemManagerVotePowerBlockSelected {
	out := make(chan *system.FlareSystemManagerVotePowerBlockSelected)
	go func() {
		ticker := time.NewTicker(10 * time.Second) // read from config
		eventRangeStart := c.epoch.StartTime(c.epoch.EpochIndex(c.mockableTime.Now()) - 1).Unix()
		for {
			<-ticker.C
			now := c.mockableTime.Now().Unix()
			logs, err := database.FetchLogsByAddressAndTopic0(c.db, c.address, c.topic0, eventRangeStart, now)
			if err != nil {
				logger.Error("Error fetching logs %w", err)
				continue
			}
			if len(logs) > 0 {
				// last log is the latest (sorted in FetchLogsByAddressAndTopic0)
				powerBlock, err := c.parseVotePowerBlockSelectedEvent(logs[len(logs)-1])
				if err != nil {
					logger.Error("Error parsing vote power block selected event %w", err)
					continue
				}
				out <- powerBlock
			}
			eventRangeStart = now
		}
	}()
	return out
}

func (c *votePowerBlockSelectedListener) parseVotePowerBlockSelectedEvent(dbLog database.Log) (*system.FlareSystemManagerVotePowerBlockSelected, error) {
	data, err := hex.DecodeString(dbLog.Data)
	if err != nil {
		return nil, err
	}
	contractLog := types.Log{
		Topics: []common.Hash{
			chain.ParseTopic(dbLog.Topic0),
			chain.ParseTopic(dbLog.Topic1),
			chain.ParseTopic(dbLog.Topic2),
			chain.ParseTopic(dbLog.Topic3),
		},
		Data: data,
		// Other fields are not used by log decoder
	}
	return c.systemManagerFilterer.ParseVotePowerBlockSelected(contractLog)
}
