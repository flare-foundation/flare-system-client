package epoch

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/system"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/flare-foundation/go-flare-common/pkg/merkle"
)

const timeout = 5 * time.Second     // maximal duration for fetching of the reward data
const maxRespSize = 100 * (1 << 20) // 100 MB for maximal response size of the reward

var (
	uint8Type, _   = abi.NewType("uint8", "", nil)
	uint64Type, _  = abi.NewType("uint64", "", nil)
	bytes20Type, _ = abi.NewType("bytes20", "", nil)
	uint120Type, _ = abi.NewType("uint120", "", nil)

	signRewardsArgs = abi.Arguments{
		{
			Type: int64Ty,
		},
		{
			Type: bytes32Ty,
		},
		{
			Type: bytes32Ty,
		},
	}
	weightClaimsType, _ = abi.NewType("tuple[]", "", []abi.ArgumentMarshaling{
		{Name: "RewardManagerId", Type: "uint256"},
		{Name: "NoOfWeightBasedClaims", Type: "uint265"},
	})
	weightClaimsArgs = abi.Arguments{
		{Type: weightClaimsType},
	}
	rewardClaimArgs = abi.Arguments{
		{Type: uint64Type},
		{Type: bytes20Type},
		{Type: uint120Type},
		{Type: uint8Type},
	}
)

type ClaimType uint8

const (
	Direct ClaimType = iota
	Fee
	WNat
	Mirror
	CChain
)

type rewardDistributionData struct {
	RewardEpochId uint64        `json:"rewardEpochId"`
	Network       string        `json:"network"`
	RewardClaims  []rewardClaim `json:"rewardClaims"`
	MerkleRoot    string        `json:"merkleRoot"`
	WeightClaims  int           `json:"noOfWeightBasedClaims"`
}

type rewardClaimBody struct {
	Beneficiary common.Address `json:"beneficiary"`
	Amount      string         `json:"amount"`
	Type        ClaimType      `json:"claimType"`
}

type rewardClaim struct {
	MerkleProof []common.Hash   `json:"merkleProof"`
	Body        rewardClaimBody `json:"body"`
}

func rewardClaimHash(epoch uint64, claim rewardClaimBody) (common.Hash, error) {
	amount, _ := new(big.Int).SetString(claim.Amount, 10)
	encoded, err := rewardClaimArgs.Pack(
		epoch,
		claim.Beneficiary,
		amount,
		claim.Type,
	)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to pack reward claim arguments")
	}
	return crypto.Keccak256Hash(encoded), nil
}

func encodeRewardsData(epochId *big.Int, chainId int, rewardHash *common.Hash, weightClaims int) []byte {
	weightClaimsWithId, err := weightClaimsArgs.Pack(
		[]system.IFlareSystemsManagerNumberOfWeightBasedClaims{{
			RewardManagerId: big.NewInt(int64(chainId)), NoOfWeightBasedClaims: big.NewInt(int64(weightClaims))},
		},
	)
	if err != nil {
		log.Fatalf("Failed to pack weight based claims arguments: %v", err)
	}

	weightClaimsWithIdHash := crypto.Keccak256Hash(weightClaimsWithId)
	packed, err := signRewardsArgs.Pack(epochId.Int64(), [32]byte(weightClaimsWithIdHash), [32]byte(*rewardHash))
	if err != nil {
		log.Fatalf("Failed to packed reward hash arguments: %v", err)
	}
	return packed
}

func fetchRewardData(epochId *big.Int, config *config.RewardsConfig) (*rewardDistributionData, error) {
	if config.UrlPrefix == "" {
		return nil, errors.New("reward data url prefix not set")
	}

	rewardsUrl, err := url.JoinPath(config.UrlPrefix, epochId.Text(10), "reward-distribution-data.json")
	if err != nil {
		return nil, errors.Errorf("cannot join url: %s", err)
	}

	logger.Infof("Fetching reward data at: %s", rewardsUrl)
	result := <-shared.ExecuteWithRetryChan(func() (*rewardDistributionData, error) {
		client := &http.Client{Timeout: timeout}

		resp, err := client.Get(rewardsUrl)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close() //nolint:errcheck

		if resp.StatusCode == http.StatusNotFound {
			return nil, nil // 404 is expected if data is not yet published, don't retry
		}
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("unexpected status code: %s", resp.Status)
		}

		respLimited := &io.LimitedReader{R: resp.Body, N: maxRespSize}

		decoder := json.NewDecoder(respLimited)

		var rewardData rewardDistributionData
		err = decoder.Decode(&rewardData)
		if err != nil {
			return nil, err
		}
		return &rewardData, nil
	}, 3, 1*time.Second)

	if !result.Success {
		return nil, errors.Errorf("unable to fetch reward data")
	}

	return result.Value, nil
}

func verifyRewardData(epochId *big.Int, identity common.Address, data *rewardDistributionData, rewardsConfig *config.RewardsConfig) (*common.Hash, int, error) {
	if data.RewardEpochId != epochId.Uint64() {
		return nil, 0, errors.Errorf("invalid rewards epoch id: %d, expected: %d", data.RewardEpochId, epochId)
	}

	root, err := verifyRoot(data)
	if err != nil {
		return nil, 0, errors.Wrap(err, "invalid rewards merkle root")
	}

	var myClaim *rewardClaim
	for i := range data.RewardClaims {
		if bytes.Equal(data.RewardClaims[i].Body.Beneficiary.Bytes(), identity.Bytes()) {
			myClaim = &data.RewardClaims[i]
		}
	}

	if myClaim == nil {
		return nil, 0, errors.Errorf("reward claim for our identity address %s not found in reward distribution data", identity.Hex())
	}

	claimHash, _ := rewardClaimHash(data.RewardEpochId, myClaim.Body)
	if !merkle.VerifyProof(claimHash, myClaim.MerkleProof, root) {
		return nil, 0, errors.Errorf("invalid merkle proof for our reward claim: %+v", myClaim.Body)
	}

	rewardAmount, ok := new(big.Int).SetString(myClaim.Body.Amount, 10)
	if !ok {
		return nil, 0, errors.Errorf("invalid reward amount: %s", myClaim.Body.Amount)
	}

	if rewardAmount.Cmp(rewardsConfig.MinRewardWei) < 0 {
		return nil, 0, errors.Errorf("reward amount %s is less than min reward %s, will not sign", rewardAmount, rewardsConfig.MinRewardWei)
	}
	if rewardsConfig.MaxRewardWei.Cmp(new(big.Int)) != 0 && rewardAmount.Cmp(rewardsConfig.MaxRewardWei) > 0 {
		return nil, 0, errors.Errorf("reward amount %s is greater than max reward %s, will not sign", rewardAmount, rewardsConfig.MaxRewardWei)
	}

	return &root, data.WeightClaims, nil
}

func verifyRoot(data *rewardDistributionData) (common.Hash, error) {
	var hashes []common.Hash
	var weightBasedClaims = 0
	for _, claim := range data.RewardClaims {
		body := claim.Body
		if body.Type == WNat || body.Type == Mirror || body.Type == CChain {
			weightBasedClaims++
		}
		hash, err := rewardClaimHash(data.RewardEpochId, body)
		if err != nil {
			return common.Hash{}, errors.Wrap(err, "unable to hash reward claim")
		}
		hashes = append(hashes, hash)
	}

	if weightBasedClaims != data.WeightClaims {
		return common.Hash{}, errors.Errorf("weight based claims count does not match: %d, expected: %d", weightBasedClaims, data.WeightClaims)
	}

	merkleTree := merkle.Build(hashes, false)
	root, err := merkleTree.Root()
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "unable to calculate merkle root")
	}

	if root.Hex() != data.MerkleRoot {
		return common.Hash{}, errors.Errorf("computed merkle root does not match: %s, expected: %s", root.Hex(), data.MerkleRoot)
	}
	return root, nil
}
