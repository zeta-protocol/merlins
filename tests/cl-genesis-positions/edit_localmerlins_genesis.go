package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/server"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	tmjson "github.com/tendermint/tendermint/libs/json"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	merlinsApp "github.com/merlins-labs/merlins/v15/app"
	"github.com/merlins-labs/merlins/v15/x/concentrated-liquidity/model"

	cltypes "github.com/merlins-labs/merlins/v15/x/concentrated-liquidity/types"
	clgenesis "github.com/merlins-labs/merlins/v15/x/concentrated-liquidity/types/genesis"
	poolmanagertypes "github.com/merlins-labs/merlins/v15/x/poolmanager/types"
)

func EditLocalMerlinsGenesis(updatedCLGenesis *clgenesis.GenesisState, updatedBankGenesis *banktypes.GenesisState) {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(localMerlinsHomePath)
	config.Moniker = "localmerlins"

	genFile := config.GenesisFile()
	appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
	if err != nil {
		panic(err)
	}

	encodingConfig := merlinsApp.MakeEncodingConfig()
	cdc := encodingConfig.Marshaler

	// Concentrated liquidity genesis.
	var localMerlinsCLGenesis clgenesis.GenesisState
	cdc.MustUnmarshalJSON(appState[cltypes.ModuleName], &localMerlinsCLGenesis)

	// Pool manager genesis.
	var localMerlinsPoolManagerGenesis poolmanagertypes.GenesisState
	cdc.MustUnmarshalJSON(appState[poolmanagertypes.ModuleName], &localMerlinsPoolManagerGenesis)

	var localMerlinsBankGenesis banktypes.GenesisState
	cdc.MustUnmarshalJSON(appState[banktypes.ModuleName], &localMerlinsBankGenesis)

	nextPoolId := localMerlinsPoolManagerGenesis.NextPoolId
	localMerlinsPoolManagerGenesis.NextPoolId = nextPoolId + 1
	localMerlinsPoolManagerGenesis.PoolRoutes = append(localMerlinsPoolManagerGenesis.PoolRoutes, poolmanagertypes.ModuleRoute{
		PoolType: poolmanagertypes.Concentrated,
		PoolId:   nextPoolId,
	})
	appState[poolmanagertypes.ModuleName] = cdc.MustMarshalJSON(&localMerlinsPoolManagerGenesis)

	// Copy positions
	largestPositionId := uint64(0)
	for _, positionData := range updatedCLGenesis.PositionData {
		positionData.Position.PoolId = nextPoolId
		localMerlinsCLGenesis.PositionData = append(localMerlinsCLGenesis.PositionData, positionData)
		if positionData.Position.PositionId > largestPositionId {
			largestPositionId = positionData.Position.PositionId
		}
	}

	// Create map of pool balances
	balancesMap := map[string][]banktypes.Balance{}
	for _, balance := range updatedBankGenesis.Balances {
		if _, ok := balancesMap[balance.Address]; !ok {
			balancesMap[balance.Address] = []banktypes.Balance{}
		}
		balancesMap[balance.Address] = append(balancesMap[balance.Address], balance)
	}

	// Copy pool state, including ticks, incentive accums, records, and fee accumulators
	for _, pool := range updatedCLGenesis.PoolData {
		poolAny := pool.Pool

		var clPoolExt cltypes.ConcentratedPoolExtension
		err := cdc.UnpackAny(poolAny, &clPoolExt)
		if err != nil {
			panic(err)
		}

		clPool, error := clPoolExt.(*model.Pool)
		if !error {
			panic("Error converting pool")
		}
		clPool.Id = nextPoolId

		any, err := codectypes.NewAnyWithValue(clPool)
		if err != nil {
			panic(err)
		}
		anyCopy := *any

		for i := range pool.Ticks {
			pool.Ticks[i].PoolId = nextPoolId
		}

		for i := range pool.IncentiveRecords {
			pool.IncentiveRecords[i].PoolId = nextPoolId
		}

		for i := range pool.IncentivesAccumulators {
			pool.IncentivesAccumulators[i].Name = cltypes.KeyUptimeAccumulator(nextPoolId, uint64(i))
		}

		updatedPoolData := clgenesis.PoolData{
			Pool:                   &anyCopy,
			Ticks:                  pool.Ticks,
			IncentivesAccumulators: pool.IncentivesAccumulators,
			IncentiveRecords:       pool.IncentiveRecords,
			FeeAccumulator: clgenesis.AccumObject{
				Name:         cltypes.KeyFeePoolAccumulator(nextPoolId),
				AccumContent: pool.FeeAccumulator.AccumContent,
			},
		}

		// Update bank genesis with balances
		poolBalances := balancesMap[clPool.GetAddress().String()]
		localMerlinsBankGenesis.Balances = append(localMerlinsBankGenesis.Balances, poolBalances...)

		localMerlinsCLGenesis.PoolData = append(localMerlinsCLGenesis.PoolData, updatedPoolData)
	}

	localMerlinsCLGenesis.NextPositionId = largestPositionId + 1

	appState[cltypes.ModuleName] = cdc.MustMarshalJSON(&localMerlinsCLGenesis)

	// Persist updated bank genesis
	appState[banktypes.ModuleName] = cdc.MustMarshalJSON(&localMerlinsBankGenesis)

	appStateJSON, err := json.Marshal(appState)
	if err != nil {
		panic(err)
	}

	genDoc.AppState = appStateJSON

	genesisJson, err := tmjson.MarshalIndent(genDoc, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Writing genesis file to %s", localMerlinsHomePath)
	if err := WriteFile(filepath.Join(localMerlinsHomePath, "config", "genesis.json"), genesisJson); err != nil {
		panic(err)
	}
}

func WriteFile(path string, body []byte) error {
	_, err := os.Create(path)
	if err != nil {
		return err
	}

	return os.WriteFile(path, body, 0o600)
}
