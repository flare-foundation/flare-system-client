package epoch

import (
	"encoding/hex"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/flare-system-client/client/config"
	"github.com/stretchr/testify/require"
)

func Test_encodeRewardsData(t *testing.T) {
	t.Run("encodeRewardsData packs arguments correctly", func(t *testing.T) {
		rewardHash := common.HexToHash("0x9d91c7d9595969e7d21783b904f8707316d6267c656b60fad0e070e9c698a672")
		encoded := encodeRewardsData(
			big.NewInt(3),
			31337,
			&rewardHash,
			56,
		)
		encodedHashHex := hex.EncodeToString(crypto.Keccak256(encoded))
		require.Equal(t, "ac5a8c3adc6d9a499eb3bc1440a5e07b041d77b0a508bebe7429d189a41acc6a", encodedHashHex)
	})
}

func Test_verifyRewardData(t *testing.T) {
	type args struct {
		epochId       *big.Int
		identity      common.Address
		data          *rewardDistributionData
		rewardsConfig *config.RewardsConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *common.Hash
		want1   int
		wantErr bool
	}{
		{
			name: "valid reward data",
			args: args{
				epochId:  big.NewInt(1),
				identity: common.HexToAddress("0xA"),
				data: &rewardDistributionData{
					RewardEpochId: 1,
					Network:       "testnet",
					RewardClaims: []rewardClaim{
						{
							MerkleProof: []common.Hash{common.HexToHash("0x2f737c46d4149cda267ba09839ed3c92caa82663c4f39598959b8d82e4399338")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xA"),
								Amount:      "100",
								Type:        0,
							},
						},
						{
							MerkleProof: []common.Hash{common.HexToHash("0x7959f725a42594181d961f863bda384ea8a8b0142d02fb7f18de93ba3f941eb5")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xB"),
								Amount:      "1000",
								Type:        2,
							},
						},
					},
					MerkleRoot:   "0x4817c8012f9a80ba36ca047cb648bbfa01924b236994ab98b4ae91ee60fd8058",
					WeightClaims: 1,
				},
				rewardsConfig: &config.RewardsConfig{
					MinRewardWei: big.NewInt(100),
					MaxRewardWei: big.NewInt(10000),
				},
			},
			want:    hexToHashP("0x4817c8012f9a80ba36ca047cb648bbfa01924b236994ab98b4ae91ee60fd8058"),
			want1:   1,
			wantErr: false,
		},
		{
			name: "invalid merkle root",
			args: args{
				epochId:  big.NewInt(1),
				identity: common.HexToAddress("0xA"),
				data: &rewardDistributionData{
					RewardEpochId: 1,
					Network:       "testnet",
					RewardClaims: []rewardClaim{
						{
							MerkleProof: []common.Hash{common.HexToHash("0x2f737c46d4149cda267ba09839ed3c92caa82663c4f39598959b8d82e4399338")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xA"),
								Amount:      "100",
								Type:        0,
							},
						},
						{
							MerkleProof: []common.Hash{common.HexToHash("0x7959f725a42594181d961f863bda384ea8a8b0142d02fb7f18de93ba3f941eb5")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xB"),
								Amount:      "1000",
								Type:        2,
							},
						},
					},
					MerkleRoot:   "0x0",
					WeightClaims: 1,
				},
				rewardsConfig: &config.RewardsConfig{
					MinRewardWei: big.NewInt(100),
					MaxRewardWei: big.NewInt(10000),
				},
			},
			want:    nil,
			want1:   0,
			wantErr: true,
		},
		{
			name: "invalid epoch",
			args: args{
				epochId:  big.NewInt(1),
				identity: common.HexToAddress("0xA"),
				data: &rewardDistributionData{
					RewardEpochId: 2,
					Network:       "testnet",
					RewardClaims: []rewardClaim{
						{
							MerkleProof: []common.Hash{common.HexToHash("0x2f737c46d4149cda267ba09839ed3c92caa82663c4f39598959b8d82e4399338")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xA"),
								Amount:      "100",
								Type:        0,
							},
						},
						{
							MerkleProof: []common.Hash{common.HexToHash("0x7959f725a42594181d961f863bda384ea8a8b0142d02fb7f18de93ba3f941eb5")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xB"),
								Amount:      "1000",
								Type:        2,
							},
						},
					},
					MerkleRoot:   "0x4817c8012f9a80ba36ca047cb648bbfa01924b236994ab98b4ae91ee60fd8058",
					WeightClaims: 1,
				},
				rewardsConfig: &config.RewardsConfig{
					MinRewardWei: big.NewInt(100),
					MaxRewardWei: big.NewInt(10000),
				},
			},
			want:    nil,
			want1:   0,
			wantErr: true,
		},
		{
			name: "invalid weight claim count",
			args: args{
				epochId:  big.NewInt(1),
				identity: common.HexToAddress("0xA"),
				data: &rewardDistributionData{
					RewardEpochId: 1,
					Network:       "testnet",
					RewardClaims: []rewardClaim{
						{
							MerkleProof: []common.Hash{common.HexToHash("0x2f737c46d4149cda267ba09839ed3c92caa82663c4f39598959b8d82e4399338")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xA"),
								Amount:      "100",
								Type:        0,
							},
						},
						{
							MerkleProof: []common.Hash{common.HexToHash("0x7959f725a42594181d961f863bda384ea8a8b0142d02fb7f18de93ba3f941eb5")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xB"),
								Amount:      "1000",
								Type:        2,
							},
						},
					},
					MerkleRoot:   "0x4817c8012f9a80ba36ca047cb648bbfa01924b236994ab98b4ae91ee60fd8058",
					WeightClaims: 0,
				},
				rewardsConfig: &config.RewardsConfig{
					MinRewardWei: big.NewInt(100),
					MaxRewardWei: big.NewInt(10000),
				},
			},
			want:    nil,
			want1:   0,
			wantErr: true,
		},
		{
			name: "invalid claim merkle proof",
			args: args{
				epochId:  big.NewInt(1),
				identity: common.HexToAddress("0xA"),
				data: &rewardDistributionData{
					RewardEpochId: 1,
					Network:       "testnet",
					RewardClaims: []rewardClaim{
						{
							MerkleProof: []common.Hash{common.HexToHash("0x1")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xA"),
								Amount:      "100",
								Type:        0,
							},
						},
						{
							MerkleProof: []common.Hash{common.HexToHash("0x7959f725a42594181d961f863bda384ea8a8b0142d02fb7f18de93ba3f941eb5")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xB"),
								Amount:      "1000",
								Type:        2,
							},
						},
					},
					MerkleRoot:   "0x4817c8012f9a80ba36ca047cb648bbfa01924b236994ab98b4ae91ee60fd8058",
					WeightClaims: 1,
				},
				rewardsConfig: &config.RewardsConfig{
					MinRewardWei: big.NewInt(100),
					MaxRewardWei: big.NewInt(10000),
				},
			},
			want:    nil,
			want1:   0,
			wantErr: true,
		},
		{
			name: "reward amount too low",
			args: args{
				epochId:  big.NewInt(1),
				identity: common.HexToAddress("0xA"),
				data: &rewardDistributionData{
					RewardEpochId: 1,
					Network:       "testnet",
					RewardClaims: []rewardClaim{
						{
							MerkleProof: []common.Hash{common.HexToHash("0x2f737c46d4149cda267ba09839ed3c92caa82663c4f39598959b8d82e4399338")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xA"),
								Amount:      "100",
								Type:        0,
							},
						},
						{
							MerkleProof: []common.Hash{common.HexToHash("0x7959f725a42594181d961f863bda384ea8a8b0142d02fb7f18de93ba3f941eb5")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xB"),
								Amount:      "1000",
								Type:        2,
							},
						},
					},
					MerkleRoot:   "0x4817c8012f9a80ba36ca047cb648bbfa01924b236994ab98b4ae91ee60fd8058",
					WeightClaims: 1,
				},
				rewardsConfig: &config.RewardsConfig{
					MinRewardWei: big.NewInt(1000),
					MaxRewardWei: big.NewInt(10000),
				},
			},
			want:    nil,
			want1:   0,
			wantErr: true,
		},
		{
			name: "reward amount too high",
			args: args{
				epochId:  big.NewInt(1),
				identity: common.HexToAddress("0xA"),
				data: &rewardDistributionData{
					RewardEpochId: 1,
					Network:       "testnet",
					RewardClaims: []rewardClaim{
						{
							MerkleProof: []common.Hash{common.HexToHash("0x2f737c46d4149cda267ba09839ed3c92caa82663c4f39598959b8d82e4399338")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xA"),
								Amount:      "100",
								Type:        0,
							},
						},
						{
							MerkleProof: []common.Hash{common.HexToHash("0x7959f725a42594181d961f863bda384ea8a8b0142d02fb7f18de93ba3f941eb5")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xB"),
								Amount:      "1000",
								Type:        2,
							},
						},
					},
					MerkleRoot:   "0x4817c8012f9a80ba36ca047cb648bbfa01924b236994ab98b4ae91ee60fd8058",
					WeightClaims: 1,
				},
				rewardsConfig: &config.RewardsConfig{
					MinRewardWei: big.NewInt(0),
					MaxRewardWei: big.NewInt(99),
				},
			},
			want:    nil,
			want1:   0,
			wantErr: true,
		},
		{
			name: "reward missing",
			args: args{
				epochId:  big.NewInt(1),
				identity: common.HexToAddress("0xA"),
				data: &rewardDistributionData{
					RewardEpochId: 1,
					Network:       "testnet",
					RewardClaims: []rewardClaim{
						{
							MerkleProof: []common.Hash{common.HexToHash("0x2f737c46d4149cda267ba09839ed3c92caa82663c4f39598959b8d82e4399338")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xC"),
								Amount:      "100",
								Type:        0,
							},
						},
						{
							MerkleProof: []common.Hash{common.HexToHash("0xef06df783e02fd1147499bfba59770e709e07ec79f3524462217427db6d2cc83")},
							Body: rewardClaimBody{

								Beneficiary: common.HexToAddress("0xB"),
								Amount:      "1000",
								Type:        2,
							},
						},
					},
					MerkleRoot:   "0x1df994686aba50e4bed9312bd0f4e6c54aa81205c5b49a9e5ad0609555e681d9",
					WeightClaims: 1,
				},
				rewardsConfig: &config.RewardsConfig{
					MinRewardWei: big.NewInt(100),
					MaxRewardWei: big.NewInt(10000),
				},
			},
			want:    nil,
			want1:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := verifyRewardData(tt.args.epochId, tt.args.identity, tt.args.data, tt.args.rewardsConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyRewardData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verifyRewardData() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("verifyRewardData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func hexToHashP(hex string) *common.Hash {
	hash := common.HexToHash(hex)
	return &hash
}
