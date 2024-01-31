package registration

import (
	"context"
	"flare-tlc/client/shared"
	"flare-tlc/config"
	"flare-tlc/database"
	"flare-tlc/logger"
	"flare-tlc/utils"
	"flare-tlc/utils/contracts/relay"
	"flare-tlc/utils/contracts/system"
	"math/big"
	"testing"
	"time"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestRegistrationClient(t *testing.T) {
	logger.Configure(config.LoggerConfig{
		Level:   "DEBUG",
		Console: true,
	})

	systemManagerClient := newTestSystemManagerClient()
	relayClient := newTestRelayClient()
	registryClient := newTestRegistryClient()

	c := &registrationClient{
		db:                  testDB{},
		systemManagerClient: systemManagerClient,
		relayClient:         relayClient,
		registryClient:      registryClient,
		identityAddress:     "0x0",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return c.RunContext(ctx)
	})

	rewardEpochID := big.NewInt(2)

	t.Log("sending test VPBS")
	systemManagerClient.sendTestVPBS(&system.FlareSystemManagerVotePowerBlockSelected{
		RewardEpochId: rewardEpochID,
	})

	t.Log("sending test policy")
	relayClient.sendTestPolicy(&relay.RelaySigningPolicyInitialized{
		RewardEpochId: rewardEpochID,
		Seed:          big.NewInt(1),
	})

	t.Log("stopping runner")
	cancel()
	err := eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %s", err.Error())

	t.Run("registered voters", func(t *testing.T) {
		t.Logf("registered voters: %v", registryClient.registeredVoters)
		cupaloy.SnapshotT(t, registryClient.registeredVoters)
	})

	t.Run("signed policies", func(t *testing.T) {
		t.Logf("signed policies: %v", systemManagerClient.signedPolicies)
		cupaloy.SnapshotT(t, systemManagerClient.signedPolicies)
	})
}

type testDB struct{}

func (db testDB) FetchLogsByAddressAndTopic0(
	address common.Address, topic0 string, from, to int64,
) ([]database.Log, error) {
	return nil, errors.New("not implemented")
}

type testSystemManagerClient struct {
	rewardEpoch    *utils.Epoch
	vpbsChan       chan *system.FlareSystemManagerVotePowerBlockSelected
	signedPolicies map[string][]byte
}

func newTestSystemManagerClient() testSystemManagerClient {
	return testSystemManagerClient{
		rewardEpoch: &utils.Epoch{
			Start:  time.Time{},
			Period: time.Hour,
		},
		vpbsChan:       make(chan *system.FlareSystemManagerVotePowerBlockSelected),
		signedPolicies: make(map[string][]byte),
	}
}

func (c testSystemManagerClient) sendTestVPBS(vpbs *system.FlareSystemManagerVotePowerBlockSelected) {
	c.vpbsChan <- vpbs
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
	return shared.ExecuteWithRetry(func() (any, error) {
		key := epochID.String()

		if _, ok := c.signedPolicies[key]; ok {
			return nil, errors.New("already signed")
		}

		c.signedPolicies[key] = policy

		return nil, nil
	}, 1, 0)
}

func (c testSystemManagerClient) GetCurrentRewardEpochId() <-chan shared.ExecuteStatus[*big.Int] {
	return shared.ExecuteWithRetry(func() (*big.Int, error) {
		return big.NewInt(1), nil
	}, 1, 0)
}

type testRelayClient struct {
	policyChan chan *relay.RelaySigningPolicyInitialized
}

func newTestRelayClient() testRelayClient {
	return testRelayClient{
		policyChan: make(chan *relay.RelaySigningPolicyInitialized),
	}
}

func (c testRelayClient) sendTestPolicy(policy *relay.RelaySigningPolicyInitialized) {
	c.policyChan <- policy
}

func (c testRelayClient) SigningPolicyInitializedListener(
	db registrationClientDB, timestamp uint64,
) <-chan *relay.RelaySigningPolicyInitialized {
	return c.policyChan
}

type testRegistryClient struct {
	registeredVoters map[string]map[string]bool
}

func newTestRegistryClient() testRegistryClient {
	return testRegistryClient{
		registeredVoters: make(map[string]map[string]bool),
	}
}

func (c testRegistryClient) RegisterVoter(
	epochID *big.Int, address string,
) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetry(func() (any, error) {
		key := epochID.String()

		if _, ok := c.registeredVoters[key]; !ok {
			c.registeredVoters[key] = make(map[string]bool)
		}

		if c.registeredVoters[key][address] {
			return nil, errors.New("already registered")
		}

		c.registeredVoters[key][address] = true

		return nil, nil
	}, 1, 0)
}
