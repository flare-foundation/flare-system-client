package voters

import (
	"bytes"
	"encoding/binary"
	"errors"
	"flare-tlc/client/shared"
	"math/big"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type VoterData struct {
	index  int
	weight uint16
}

type VoterSet struct {
	voters      []common.Address
	weights     []uint16
	totalWeight uint16
	thresholds  []uint16

	voterDataMap map[common.Address]VoterData
}

func NewVoterSet(voters []common.Address, weights []uint16) *VoterSet {
	vs := VoterSet{
		voters:     voters,
		weights:    weights,
		thresholds: make([]uint16, len(weights)),
	}
	// sum does not exceed uint16, guaranteed by the smart contract
	for i, w := range weights {
		vs.thresholds[i] = vs.totalWeight
		vs.totalWeight += w
	}

	vMap := make(map[common.Address]VoterData)
	for i, voter := range vs.voters {
		if _, ok := vMap[voter]; !ok {
			vMap[voter] = VoterData{
				index:  i,
				weight: vs.weights[i],
			}
		}
	}
	vs.voterDataMap = vMap
	return &vs
}

// Initial seed for random voter selection for finalization reward calculation.
// Initial seed is calculated as a hash of protocol ID and voting round ID.
// The seed is used for the first random.
func InitialHashSeed(rewardEpochSeed *big.Int, protocolId byte, votingRoundId uint32) common.Hash {
	seed := make([]byte, 96)
	// 0-31 bytes are filled with the reward epoch seed
	if rewardEpochSeed != nil {
		rewardEpochSeed.FillBytes(seed[0:32])
	}
	// 32-63 bytes are filled with the protocol ID
	seed[63] = protocolId
	// 64-95 bytes are filled with the voting round ID
	binary.BigEndian.PutUint32(seed[92:96], votingRoundId)
	return common.BytesToHash(crypto.Keccak256(seed))
}

func RandomNumberSequence(initialSeed common.Hash, length int) []common.Hash {
	sequence := make([]common.Hash, length)
	currentSeed := initialSeed
	for i := 0; i < length; i++ {
		sequence[i] = currentSeed
		currentSeed = crypto.Keccak256Hash(currentSeed.Bytes())
	}
	return sequence
}

func (vs *VoterSet) SelectVoters(rewardEpochSeed *big.Int, protocolId byte, votingRoundId uint32, thresholdBIPS uint16) (mapset.Set[common.Address], error) {
	seed := InitialHashSeed(rewardEpochSeed, protocolId, votingRoundId)
	return vs.RandomSelectThresholdWeightVoters(seed, thresholdBIPS)
}

func (vs *VoterSet) RandomSelectThresholdWeightVoters(randomSeed common.Hash, thresholdBIPS uint16) (mapset.Set[common.Address], error) {
	// We limit the threshold to 5000 BIPS to avoid long running loops
	// In practice it will be used with around 1000 BIPS or lower.
	if thresholdBIPS > 5000 {
		return nil, errors.New("Threshold must be between 0 and 5000 BIPS")
	}

	selectedWeight := uint16(0)
	thresholdWeight := uint16(uint64(vs.totalWeight) * uint64(thresholdBIPS) / 10000)
	currentSeed := randomSeed
	selectedVoters := mapset.NewSet[common.Address]()

	// If threshold weight is not too big, the loop should end quickly
	for selectedWeight < thresholdWeight {
		index := vs.selectVoterIndex(currentSeed)
		selectedAddress := vs.voters[index]
		if !selectedVoters.Contains(selectedAddress) {
			selectedVoters.Add(selectedAddress)
			selectedWeight += vs.weights[index]
		}
		currentSeed = crypto.Keccak256Hash(currentSeed.Bytes())
	}
	return selectedVoters, nil
}

// Selects a random voter based provided random number.
func (vs *VoterSet) selectVoterIndex(randomNumber common.Hash) int {
	randomWeight := big.NewInt(0).SetBytes(randomNumber.Bytes())
	randomWeight = randomWeight.Mod(randomWeight, big.NewInt(int64(vs.totalWeight)))
	return vs.BinarySearch(uint16(randomWeight.Uint64()))
}

// Searches for the highest index of the threshold that is less than or equal to the value.
// Binary search is used.
func (vs *VoterSet) BinarySearch(value uint16) int {
	if value > vs.totalWeight {
		panic("Value must be between 0 and total weight")
	}
	left := 0
	right := len(vs.thresholds) - 1
	mid := 0
	if vs.thresholds[right] <= value {
		return right
	}
	for left < right {
		mid = (left + right) / 2
		if vs.thresholds[mid] < value {
			left = mid + 1
		} else if vs.thresholds[mid] > value {
			right = mid
		} else {
			return mid
		}
	}
	return left - 1
}

func (vs *VoterSet) TotalWeight() uint16 {
	return vs.totalWeight
}
func (vs *VoterSet) VoterWeight(index int) uint16 {
	return vs.weights[index]
}

func (vs *VoterSet) Count() int {
	return len(vs.voters)
}

func (vs *VoterSet) WriteVoterRaw(buffer *bytes.Buffer, i int) {
	weightBytes := shared.Uint16toBytes(vs.weights[i])
	buffer.Write(vs.voters[i].Bytes())
	buffer.Write(weightBytes[:])
}

func (vs *VoterSet) VoterIndex(address common.Address) int {
	voterData, ok := vs.voterDataMap[address]
	if !ok {
		return -1
	}
	return voterData.index
}
