package voters_test

import (
	"flare-fsc/client/shared/voters"
	"flare-fsc/utils"
	"math/big"
	"slices"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var (
	testVoters = []common.Address{
		common.HexToAddress("0xc783df8a850f42e7f7e57013759c285caa701eb6"),
		common.HexToAddress("0xead9c93b79ae7c1591b1fb5323bd777e86e150d4"),
		common.HexToAddress("0xe5904695748fe4a84b40b3fc79de2277660bd1d3"),
		common.HexToAddress("0x92561f28ec438ee9831d00d1d59fbdc981b762b2"),
		common.HexToAddress("0x2ffd013aaa7b5a7da93336c2251075202b33fb2b"),
	}

	testWeights = []uint16{100, 200, 300, 400, 500}
)

func TestInitialHashSeed(t *testing.T) {
	seed := voters.InitialHashSeed(big.NewInt(1), 2, 3)
	if seed != common.HexToHash("0x6e0c627900b24bd432fe7b1f713f1b0744091a646a9fe4a65a18dfed21f2949c") {
		t.Errorf("initial hash seed is not correct")
	}
}

func TestVoterSetInitialization(t *testing.T) {
	vs := voters.NewVoterSet(testVoters, testWeights)
	if vs == nil {
		t.Errorf("voter set is nil")
	} else if vs.TotalWeight() != 1500 {
		t.Errorf("total weight is not correct")
	}
}

func TestBinarySearch(t *testing.T) {
	testPairs := []uint16{0, 1, 99, 100, 101, 105, 299, 300, 301, 305, 599, 600, 601, 605, 999, 1000, 1001, 1005}

	t.Run("test1", func(t *testing.T) {
		vs := voters.NewVoterSet([]common.Address{common.HexToAddress("0xc783df8a850f42e7f7e57013759c285caa701eb6")}, []uint16{100})
		testResults := make([]int, 4)
		for i := 0; i <= 3; i++ {
			testResults[i] = vs.BinarySearch(testPairs[i])
		}
		cupaloy.SnapshotT(t, testResults)
	})

	t.Run("test2", func(t *testing.T) {
		vs := voters.NewVoterSet(testVoters, testWeights)
		test2Results := make([]int, len(testPairs))
		for i := 0; i < len(testPairs); i++ {
			test2Results[i] = vs.BinarySearch(testPairs[i])
		}
		cupaloy.SnapshotT(t, test2Results)
	})
}

func TestRandomNumberSequence(t *testing.T) {
	seed := voters.InitialHashSeed(big.NewInt(1), 1, 1)
	randoms := voters.RandomNumberSequence(seed, 5)

	cupaloy.SnapshotT(t, randoms)
}

func TestSelectVoters(t *testing.T) {
	vs := voters.NewVoterSet(testVoters, testWeights)
	seed := voters.InitialHashSeed(big.NewInt(1), 1, 1)
	voterSet, err := vs.RandomSelectThresholdWeightVoters(seed, 3000)

	voterSetHex := utils.Map(voterSet.ToSlice(), func(addr common.Address) string {
		return addr.Hex()
	})
	slices.Sort(voterSetHex)
	require.NoError(t, err)
	cupaloy.SnapshotT(t, voterSetHex)
}
