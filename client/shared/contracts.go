package shared

import (
	"github.com/flare-foundation/go-flare-common/pkg/database"
	"github.com/flare-foundation/go-flare-common/pkg/events"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
)

func ParseSigningPolicyInitializedEvent(relay *relay.Relay, dbLog database.Log) (*relay.RelaySigningPolicyInitialized, error) {
	contractLog, err := events.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, err
	}
	return relay.ParseSigningPolicyInitialized(*contractLog)
}

func ParseProtocolMessageRelayedEvent(relay *relay.Relay, dbLog database.Log) (*relay.RelayProtocolMessageRelayed, error) {
	contractLog, err := events.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, err
	}
	return relay.ParseProtocolMessageRelayed(*contractLog)
}
