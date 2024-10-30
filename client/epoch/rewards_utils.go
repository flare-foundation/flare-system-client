package epoch

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/system"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

type rewardsHash struct {
	RewardEpochId         int    `json:"rewardEpochId"`
	NoOfWeightBasedClaims int    `json:"noOfWeightBasedClaims"`
	MerkleRoot            string `json:"merkleRoot"`
}

var (
	signRewardsArguments = abi.Arguments{
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
	weightClaimsArguments = abi.Arguments{
		{Type: weightClaimsType},
	}
)

func encodeRewardsData(epochId *big.Int, chainId int, rewardHash *common.Hash, weightClaims int) []byte {
	weightClaimsWithId, err := weightClaimsArguments.Pack(
		[]system.IFlareSystemsManagerNumberOfWeightBasedClaims{{
			RewardManagerId: big.NewInt(int64(chainId)), NoOfWeightBasedClaims: big.NewInt(int64(weightClaims))},
		},
	)
	if err != nil {
		log.Fatalf("Failed to pack weight based claims arguments: %v", err)
	}

	weightClaimsWithIdHash := crypto.Keccak256Hash(weightClaimsWithId)
	packed, err := signRewardsArguments.Pack(epochId.Int64(), [32]byte(weightClaimsWithIdHash), [32]byte(*rewardHash))
	if err != nil {
		log.Fatalf("Failed to packed reward hash arguments: %v", err)
	}
	return packed
}

func getRewardsHash(epochId *big.Int, rewardsConfig *config.Rewards) (*common.Hash, int, error) {
	prefix := rewardsConfig.PathPrefix
	if prefix == "" {
		return nil, 0, errors.New("rewards hash path prefix not set")
	}

	path := fmt.Sprintf("%s/%d/rewards-hash.json", prefix, epochId)
	bytes, err := fetchRewardsHashBytes(path)
	if err != nil {
		return nil, 0, err
	}

	var rewardHash rewardsHash
	err = json.Unmarshal(bytes, &rewardHash)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error decoding reward hash file")
	}

	if rewardHash.RewardEpochId != int(epochId.Int64()) {
		return nil, 0, errors.Errorf("invalid rewards hash epoch id: %d, expected: %d", rewardHash.RewardEpochId, epochId)
	}

	hashBytes, err := hexutil.Decode(rewardHash.MerkleRoot)
	if err != nil {
		return nil, 0, errors.Wrap(err, "invalid rewards merkle root")
	}
	if len(hashBytes) != common.HashLength {
		return nil, 0, errors.Errorf("invalid rewards merkle root length: %v", len(hashBytes))
	}

	hash := common.BytesToHash(hashBytes)
	return &hash, rewardHash.NoOfWeightBasedClaims, nil
}

func fetchRewardsHashBytes(path string) ([]byte, error) {
	var data []byte
	_, isUrl := parseUrl(path)
	if isUrl {
		logger.Infof("Fetching rewards hash from URL: %s", path)
		result := <-shared.ExecuteWithRetryChan(func() ([]byte, error) {
			resp, err := http.Get(path)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			bytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			return bytes, nil
		}, 3, 1*time.Second)

		if !result.Success {
			return nil, errors.Errorf("error fetching rewards hash: %s", result.Message)
		}
		data = result.Value
	} else {
		logger.Infof("Fetching rewards hash from disk: %s", path)
		file, err := os.Open(path)
		if err != nil {
			return nil, errors.Wrap(err, "error opening reward hash file")
		}
		defer file.Close()
		data, err = io.ReadAll(file)
		if err != nil {
			return nil, errors.Wrap(err, "error reading reward hash file")
		}
	}
	return data, nil
}

func parseUrl(s string) (*url.URL, bool) {
	url, err := url.ParseRequestURI(s)
	return url, err == nil && url.Scheme != "" && url.Host != ""
}
