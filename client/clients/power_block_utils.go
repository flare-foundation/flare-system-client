package clients

import (
	"flare-tlc/client/shared"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/contracts/system"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type votePowerBlockSelectedListener struct {
	C <-chan *system.FlareSystemManagerVotePowerBlockSelected

	epoch   *utils.Epoch
	address string
	topic0  string

	systemManager *system.FlareSystemManager
	db            *gorm.DB
	mockableTime  utils.TimeProvider
}

func NewVotePowerBlockSelectedListener(
	db *gorm.DB,
	systemManager *system.FlareSystemManager,
	address string,
	topic0 string,
) (*votePowerBlockSelectedListener, error) {
	epoch, err := EpochFromChain(systemManager)
	if err != nil {
		return nil, errors.Wrap(err, "error getting epoch from chain")
	}

	listener := &votePowerBlockSelectedListener{
		epoch:         epoch,
		address:       address,
		topic0:        topic0,
		systemManager: systemManager,
		db:            db,
		mockableTime:  &utils.RealTimeProvider{},
	}
	listener.C = listener.votePowerBlockChannel()

	return listener, nil
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
	contractLog, err := shared.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, err
	}
	return c.systemManager.FlareSystemManagerFilterer.ParseVotePowerBlockSelected(*contractLog)
}

func EpochFromChain(systemManager *system.FlareSystemManager) (*utils.Epoch, error) {
	epochStart, err := systemManager.RewardEpochsStartTs(nil)
	if err != nil {
		return nil, err
	}
	epochPeriod, err := systemManager.RewardEpochDurationSeconds(nil)
	if err != nil {
		return nil, err
	}
	return &utils.Epoch{
		Start:  time.Unix(int64(epochStart), 0),
		Period: time.Duration(epochPeriod) * time.Second,
	}, nil
}
