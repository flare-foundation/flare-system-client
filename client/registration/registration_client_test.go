package registration

import (
	"context"
	"flare-tlc/client/shared"
	"flare-tlc/database"
	"flare-tlc/utils"
	"flare-tlc/utils/contracts/relay"
	"flare-tlc/utils/contracts/system"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestRegistrationClient(t *testing.T) {
	c := newRegistrationClient()
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return c.Run(ctx)
	})

	cancel()

	err := eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %s", err.Error())
}

func newRegistrationClient() *registrationClient {
	return &registrationClient{
		db:                  testDB{},
		systemManagerClient: newTestSystemManagerClient(),
		relayClient:         &testRelayClient{},
		registryClient:      &testRegistryClient{},
		identityAddress:     "0x0",
	}

}

type testDB struct{}

func (db testDB) FetchLogsByAddressAndTopic0(
	address common.Address, topic0 string, from, to int64,
) ([]database.Log, error) {
	return nil, errors.New("not implemented")
}

type testSystemManagerClient struct {
	rewardEpoch *utils.Epoch
	vpbsChan    chan *system.FlareSystemManagerVotePowerBlockSelected
}

func newTestSystemManagerClient() testSystemManagerClient {
	return testSystemManagerClient{
		rewardEpoch: &utils.Epoch{
			Start:  time.Time{},
			Period: time.Hour,
		},
		vpbsChan: make(chan *system.FlareSystemManagerVotePowerBlockSelected),
	}
}

func (c testSystemManagerClient) RewardEpochFromChain() (*utils.Epoch, error) {
	return c.rewardEpoch, nil
}

func (c testSystemManagerClient) VotePowerBlockSelectedListener(
	db registrationClientDB, epoch *utils.Epoch,
) <-chan *system.FlareSystemManagerVotePowerBlockSelected {
	return c.vpbsChan
}

func (c testSystemManagerClient) SignNewSigningPolicy(
	epochID *big.Int, policy []byte,
) <-chan shared.ExecuteStatus[any] {
	return nil
}

func (c testSystemManagerClient) GetCurrentRewardEpochId() <-chan shared.ExecuteStatus[*big.Int] {
	return nil
}

type testRelayClient struct{}

func (c testRelayClient) SigningPolicyInitializedListener(
	db registrationClientDB, timestamp uint64,
) <-chan *relay.RelaySigningPolicyInitialized {
	return nil
}

type testRegistryClient struct{}

func (c testRegistryClient) RegisterVoter(
	epochID *big.Int, address string,
) <-chan shared.ExecuteStatus[any] {
	return nil
}
