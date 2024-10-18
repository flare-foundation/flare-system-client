package epoch

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/system"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

var (
	zeroBytes = [32]byte{}
	zeroHash  = crypto.Keccak256Hash(zeroBytes[:])

	int64Ty, _          = abi.NewType("int64", "int64", nil)
	bytes32Ty, _        = abi.NewType("bytes32", "bytes32", nil)
	uptimeVoteArguments = abi.Arguments{
		{ // reward epoch id
			Type: int64Ty,
		},
		{ // hash
			Type: bytes32Ty,
		},
	}
)

func getUptimeSignature(rewardEpochId *big.Int, privateKey *ecdsa.PrivateKey) (common.Hash, *system.IFlareSystemsManagerSignature, error) {
	logger.Infof("Signing uptime vote for epoch %v: zero hash %s", rewardEpochId, hex.EncodeToString(zeroHash[:]))

	toSign, err := uptimeVoteArguments.Pack(rewardEpochId.Int64(), zeroHash)
	if err != nil {
		return zeroHash, nil, errors.Wrapf(err, "error packing uptime vote arguments: %v, %v", rewardEpochId, zeroHash)
	}

	hashSignature, err := crypto.Sign(accounts.TextHash(crypto.Keccak256(toSign)), privateKey)
	if err != nil {
		return zeroHash, nil, err
	}

	signature := system.IFlareSystemsManagerSignature{
		R: [32]byte(hashSignature[0:32]),
		S: [32]byte(hashSignature[32:64]),
		V: hashSignature[64] + 27,
	}
	return zeroHash, &signature, nil
}
