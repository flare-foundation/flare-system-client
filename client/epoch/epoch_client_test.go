package epoch

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/flare-foundation/flare-system-client/client/shared"
	"github.com/flare-foundation/flare-system-client/utils"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/system"

	"github.com/flare-foundation/go-flare-common/pkg/database"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestEpochClientMainline(t *testing.T) {
	systemsManagerClient := newTestSystemsManagerClient()
	relayClient := newTestRelayClient()
	registryClient := newTestRegistryClient()

	c := &client{
		db:                   testDB{},
		systemsManagerClient: systemsManagerClient,
		relayClient:          relayClient,
		registryClient:       registryClient,
		identityAddress:      common.HexToAddress("0x123456"),
		registrationEnabled:  true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return c.Run(ctx)
	})

	rewardEpochID := big.NewInt(2)
	signingPolicyBytes := []byte{1, 2, 3}

	t.Log("sending test VPBS")
	systemsManagerClient.sendTestVPBS(&system.FlareSystemsManagerVotePowerBlockSelected{
		RewardEpochId: rewardEpochID,
	})

	t.Log("sending test policy")
	relayClient.sendTestPolicy(&relay.RelaySigningPolicyInitialized{
		RewardEpochId:      rewardEpochID,
		Seed:               big.NewInt(1),
		SigningPolicyBytes: signingPolicyBytes,
	})

	t.Log("stopping runner")
	cancel()
	err := eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %s", err.Error())

	t.Run("registered voters", func(t *testing.T) {
		t.Logf("registered voters: %v", registryClient.registeredVoters)
		require.True(t, registryClient.registeredVoters["2"][common.HexToAddress("0x123456")])
		cupaloy.SnapshotT(t, registryClient.registeredVoters)
	})

	t.Run("signed policies", func(t *testing.T) {
		t.Logf("signed policies: %v", systemsManagerClient.signedPolicies)
		require.Equal(t, signingPolicyBytes, systemsManagerClient.signedPolicies["2"])
		cupaloy.SnapshotT(t, systemsManagerClient.signedPolicies)
	})
}

func TestEpochClientInvalidEpoch(t *testing.T) {
	systemsManagerClient := newTestSystemsManagerClient()
	relayClient := newTestRelayClient()
	registryClient := newTestRegistryClient()

	c := &client{
		db:                   testDB{},
		systemsManagerClient: systemsManagerClient,
		relayClient:          relayClient,
		registryClient:       registryClient,
		identityAddress:      common.HexToAddress("0x123456"),
		registrationEnabled:  true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return c.Run(ctx)
	})

	rewardEpochID := big.NewInt(0)

	// Reward Epoch ID is 0, so this should be ignored.
	t.Log("sending test VPBS")
	systemsManagerClient.sendTestVPBS(&system.FlareSystemsManagerVotePowerBlockSelected{
		RewardEpochId: rewardEpochID,
	})

	t.Log("stopping runner")
	cancel()
	err := eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %s", err.Error())

	t.Logf("registered voters: %v", registryClient.registeredVoters)
	require.Empty(t, registryClient.registeredVoters)

	t.Logf("signed policies: %v", systemsManagerClient.signedPolicies)
	require.Empty(t, systemsManagerClient.signedPolicies)
}

func TestEpochClientSigningErr(t *testing.T) {
	systemsManagerClient := newTestSystemsManagerClient()
	systemsManagerClient.signingErr = errors.New("signing error")

	relayClient := newTestRelayClient()
	registryClient := newTestRegistryClient()

	c := &client{
		db:                   testDB{},
		systemsManagerClient: systemsManagerClient,
		relayClient:          relayClient,
		registryClient:       registryClient,
		identityAddress:      common.HexToAddress("0x123456"),
		registrationEnabled:  true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return c.Run(ctx)
	})

	rewardEpochID := big.NewInt(2)
	signingPolicyBytes := []byte{1, 2, 3}

	t.Log("sending test VPBS")
	systemsManagerClient.sendTestVPBS(&system.FlareSystemsManagerVotePowerBlockSelected{
		RewardEpochId: rewardEpochID,
	})

	t.Log("sending test policy")
	relayClient.sendTestPolicy(&relay.RelaySigningPolicyInitialized{
		RewardEpochId:      rewardEpochID,
		Seed:               big.NewInt(1),
		SigningPolicyBytes: signingPolicyBytes,
	})

	t.Log("stopping runner")
	cancel()
	err := eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %s", err.Error())

	t.Logf("registered voters: %v", registryClient.registeredVoters)
	require.True(t, registryClient.registeredVoters["2"][common.HexToAddress("0x123456")])
	cupaloy.SnapshotT(t, registryClient.registeredVoters)

	t.Logf("signed policies: %v", systemsManagerClient.signedPolicies)
	require.Empty(t, systemsManagerClient.signedPolicies)
}

func TestEpochClientRewardEpochErr(t *testing.T) {
	systemsManagerClient := newTestSystemsManagerClient()
	systemsManagerClient.rewardEpochErr = errors.New("reward epoch error")

	relayClient := newTestRelayClient()
	registryClient := newTestRegistryClient()

	c := &client{
		db:                   testDB{},
		systemsManagerClient: systemsManagerClient,
		relayClient:          relayClient,
		registryClient:       registryClient,
		identityAddress:      common.HexToAddress("0x123456"),
		registrationEnabled:  true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return c.Run(ctx)
	})

	rewardEpochID := big.NewInt(2)

	t.Log("sending test VPBS")
	systemsManagerClient.sendTestVPBS(&system.FlareSystemsManagerVotePowerBlockSelected{
		RewardEpochId: rewardEpochID,
	})

	t.Log("stopping runner")
	cancel()
	err := eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %s", err.Error())

	t.Logf("registered voters: %v", registryClient.registeredVoters)
	require.Empty(t, registryClient.registeredVoters)

	t.Logf("signed policies: %v", systemsManagerClient.signedPolicies)
	require.Empty(t, systemsManagerClient.signedPolicies)
}

func TestEpochClientRegisterErr(t *testing.T) {
	systemsManagerClient := newTestSystemsManagerClient()
	relayClient := newTestRelayClient()
	registryClient := newTestRegistryClient()
	registryClient.registerErr = errors.New("register error")

	c := &client{
		db:                   testDB{},
		systemsManagerClient: systemsManagerClient,
		relayClient:          relayClient,
		registryClient:       registryClient,
		identityAddress:      common.HexToAddress("0x123456"),
		registrationEnabled:  true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return c.Run(ctx)
	})

	rewardEpochID := big.NewInt(2)

	t.Log("sending test VPBS")
	systemsManagerClient.sendTestVPBS(&system.FlareSystemsManagerVotePowerBlockSelected{
		RewardEpochId: rewardEpochID,
	})

	t.Log("stopping runner")
	cancel()
	err := eg.Wait()
	require.True(t, errors.Is(err, context.Canceled), "unexpected error: %s", err.Error())

	t.Logf("registered voters: %v", registryClient.registeredVoters)
	require.Empty(t, registryClient.registeredVoters)

	t.Logf("signed policies: %v", systemsManagerClient.signedPolicies)
	require.Empty(t, systemsManagerClient.signedPolicies)
}

type testDB struct{}

func (db testDB) FetchLogsByAddressAndTopic0Timestamp(
	address common.Address, topic0 common.Hash, from, to int64,
) ([]database.Log, error) {
	return nil, errors.New("not implemented")
}

type testSystemsManagerClient struct {
	rewardEpoch    *utils.EpochConfig
	rewardEpochErr error
	vpbsChan       chan *system.FlareSystemsManagerVotePowerBlockSelected
	signedPolicies map[string][]byte
	signingErr     error
}

func newTestSystemsManagerClient() testSystemsManagerClient {
	return testSystemsManagerClient{
		rewardEpoch: &utils.EpochConfig{
			Start:  time.Time{},
			Period: time.Hour,
		},
		vpbsChan:       make(chan *system.FlareSystemsManagerVotePowerBlockSelected),
		signedPolicies: make(map[string][]byte),
	}
}

func (c testSystemsManagerClient) sendTestVPBS(vpbs *system.FlareSystemsManagerVotePowerBlockSelected) {
	c.vpbsChan <- vpbs
}

func (c testSystemsManagerClient) RewardEpochFromChain() (*utils.EpochConfig, error) {
	return c.rewardEpoch, nil
}

func (c testSystemsManagerClient) VotePowerBlockSelectedListener(
	db epochClientDB, epoch *utils.EpochConfig,
) <-chan *system.FlareSystemsManagerVotePowerBlockSelected {
	return c.vpbsChan
}

func (c testSystemsManagerClient) SignNewSigningPolicy(
	epochID *big.Int, policy []byte,
) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetry(func() (any, error) {
		if c.signingErr != nil {
			return nil, c.signingErr
		}

		key := epochID.String()

		if _, ok := c.signedPolicies[key]; ok {
			return nil, errors.New("already signed")
		}

		c.signedPolicies[key] = policy

		return nil, nil
	}, 1, 0)
}

func (c testSystemsManagerClient) GetCurrentRewardEpochID() <-chan shared.ExecuteStatus[*big.Int] {
	return shared.ExecuteWithRetry(func() (*big.Int, error) {
		if c.rewardEpochErr != nil {
			return nil, c.rewardEpochErr
		}

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
	db epochClientDB, epoch *utils.EpochConfig,
) <-chan *relay.RelaySigningPolicyInitialized {
	return c.policyChan
}

type testRegistryClient struct {
	registeredVoters map[string]map[common.Address]bool
	registerErr      error
}

func newTestRegistryClient() testRegistryClient {
	return testRegistryClient{
		registeredVoters: make(map[string]map[common.Address]bool),
	}
}

func (c testRegistryClient) RegisterVoter(
	epochID *big.Int, address common.Address,
) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetry(func() (any, error) {
		if c.registerErr != nil {
			return nil, c.registerErr
		}

		key := epochID.String()

		if _, ok := c.registeredVoters[key]; !ok {
			c.registeredVoters[key] = make(map[common.Address]bool)
		}

		if c.registeredVoters[key][address] {
			return nil, errors.New("already registered")
		}

		c.registeredVoters[key][address] = true

		return nil, nil
	}, 1, 0)
}

func (c testSystemsManagerClient) SignUptimeVoteEnabledListener(db epochClientDB, epoch *utils.EpochConfig, i int64) <-chan *system.FlareSystemsManagerSignUptimeVoteEnabled {
	return make(chan *system.FlareSystemsManagerSignUptimeVoteEnabled)
}

func (c testSystemsManagerClient) SignUptimeVote(b *big.Int) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetry(func() (any, error) {
		return nil, nil
	}, 1, 0)
}

func (c testSystemsManagerClient) UptimeVoteSignedListener(db epochClientDB, epoch *utils.EpochConfig, window int64) <-chan *system.FlareSystemsManagerUptimeVoteSigned {
	return make(chan *system.FlareSystemsManagerUptimeVoteSigned)
}

func (c testSystemsManagerClient) SignRewards(b *big.Int, hash *common.Hash, claims int) <-chan shared.ExecuteStatus[any] {
	return shared.ExecuteWithRetry(func() (any, error) {
		return nil, nil
	}, 1, 0)
}
