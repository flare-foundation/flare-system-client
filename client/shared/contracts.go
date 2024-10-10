package shared

import (
	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/database"
	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/events"

	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/contracts/relay"
)

func ParseSigningPolicyInitializedEvent(relay *relay.Relay, dbLog database.Log) (*relay.RelaySigningPolicyInitialized, error) {
	contractLog, err := events.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, err
	}
	return relay.RelayFilterer.ParseSigningPolicyInitialized(*contractLog)
}

func ParseProtocolMessageRelayedEvent(relay *relay.Relay, dbLog database.Log) (*relay.RelayProtocolMessageRelayed, error) {
	contractLog, err := events.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, err
	}
	return relay.RelayFilterer.ParseProtocolMessageRelayed(*contractLog)
}
