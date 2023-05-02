package simapp

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"

	furysim "github.com/merlins-labs/merlins/v15/simulation/executor"
	"github.com/merlins-labs/merlins/v15/simulation/simtypes/simlogger"
	txfeetypes "github.com/merlins-labs/merlins/v15/x/txfees/types"
)

// Profile with:
// /usr/local/go/bin/go test -benchmem -run=^$ github.com/merlins-labs/merlins/simapp -bench ^BenchmarkFullAppSimulation$ -Commit=true -cpuprofile cpu.out
func BenchmarkFullAppSimulation(b *testing.B) {
	// -Enabled=true -NumBlocks=1000 -BlockSize=200 \
	// -Period=1 -Commit=true -Seed=57 -v -timeout 24h
	furysim.FlagEnabledValue = true
	furysim.FlagNumBlocksValue = 1000
	furysim.FlagBlockSizeValue = 200
	furysim.FlagCommitValue = true
	furysim.FlagVerboseValue = true
	// furysim.FlagPeriodValue = 1000
	fullAppSimulation(b, false)
}

func TestFullAppSimulation(t *testing.T) {
	// -Enabled=true -NumBlocks=1000 -BlockSize=200 \
	// -Period=1 -Commit=true -Seed=57 -v -timeout 24h
	furysim.FlagEnabledValue = true
	furysim.FlagNumBlocksValue = 200
	furysim.FlagBlockSizeValue = 25
	furysim.FlagCommitValue = true
	furysim.FlagVerboseValue = true
	furysim.FlagPeriodValue = 10
	furysim.FlagSeedValue = 11
	furysim.FlagWriteStatsToDB = true
	fullAppSimulation(t, true)
}

func fullAppSimulation(tb testing.TB, is_testing bool) {
	// TODO: Get SDK simulator fixed to have min fees possible
	txfeetypes.ConsensusMinFee = sdk.ZeroDec()
	config, db, logger, cleanup, err := furysim.SetupSimulation("goleveldb-app-sim", "Simulation")
	if err != nil {
		tb.Fatalf("simulation setup failed: %s", err.Error())
	}
	defer cleanup()
	// This file is needed to provide the correct path
	// to reflect.wasm test file needed for wasmd simulation testing.
	config.InitializationConfig.ParamsFile = "params.json"
	config.ExecutionDbConfig.UseMerkleTree = !is_testing

	// Run randomized simulation:
	_, _, simErr := furysim.SimulateFromSeed(
		tb,
		os.Stdout,
		MerlinsAppCreator(logger, db),
		MerlinsInitFns,
		config,
	)

	if simErr != nil {
		tb.Fatal(simErr)
	}

	if config.ExecutionDbConfig.UseMerkleTree {
		furysim.PrintStats(db)
	}
}

// TODO: Make another test for the fuzzer itself, which just has noOp txs
// and doesn't depend on the application.
func TestAppStateDeterminism(t *testing.T) {
	// if !furysim.FlagEnabledValue {
	// 	t.Skip("skipping application simulation")
	// }
	// TODO: Get SDK simulator fixed to have min fees possible
	txfeetypes.ConsensusMinFee = sdk.ZeroDec()

	config := furysim.NewConfigFromFlags()
	config.ExportConfig.ExportParamsPath = ""
	config.NumBlocks = 50
	config.BlockSize = 5
	config.OnOperation = false
	config.AllInvariants = false
	config.InitializationConfig.ChainID = helpers.SimAppChainID

	// This file is needed to provide the correct path
	// to reflect.wasm test file needed for wasmd simulation testing.
	config.InitializationConfig.ParamsFile = "params.json"

	numSeeds := 3
	numTimesToRunPerSeed := 5
	appHashList := make([]string, numTimesToRunPerSeed)

	for i := 0; i < numSeeds; i++ {
		config.Seed = rand.Int63()

		for j := 0; j < numTimesToRunPerSeed; j++ {
			logger := simlogger.NewSimLogger(log.TestingLogger())
			db := dbm.NewMemDB()

			fmt.Printf(
				"running non-determinism simulation; seed %d: %d/%d, attempt: %d/%d\n",
				config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
			)

			// Run randomized simulation:
			lastCommitId, _, simErr := furysim.SimulateFromSeed(
				t,
				os.Stdout,
				MerlinsAppCreator(logger, db),
				MerlinsInitFns,
				config,
			)

			require.NoError(t, simErr)

			appHash := lastCommitId.Hash
			appHashList[j] = fmt.Sprintf("%X", appHash)

			if j != 0 {
				require.Equal(
					t, appHashList[0], appHashList[j],
					"non-determinism in seed %d: %d/%d, attempt: %d/%d\n", config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
				)
			}
		}
	}
}
