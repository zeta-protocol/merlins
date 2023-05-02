package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	simapp "github.com/merlins-labs/merlins/v15/app"
	"github.com/merlins-labs/merlins/v15/x/superfluid"
	"github.com/merlins-labs/merlins/v15/x/superfluid/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var now = time.Now().UTC()

var testGenesis = types.GenesisState{
	Params: types.Params{
		MinimumRiskFactor: sdk.NewDecWithPrec(5, 1), // 50%
	},
	SuperfluidAssets: []types.SuperfluidAsset{
		{
			Denom:     DefaultGammAsset,
			AssetType: types.SuperfluidAssetTypeLPShare,
		},
	},
	FuryEquivalentMultipliers: []types.FuryEquivalentMultiplierRecord{
		{
			EpochNumber: 1,
			Denom:       DefaultGammAsset,
			Multiplier:  sdk.NewDec(1000),
		},
	},
	IntermediaryAccounts: []types.SuperfluidIntermediaryAccount{
		{
			Denom:   DefaultGammAsset,
			ValAddr: "furyvaloper1cyw4vw20el8e7ez8080md0r8psg25n0cq98a9n",
			GaugeId: 1,
		},
	},
	IntemediaryAccountConnections: []types.LockIdIntermediaryAccountConnection{
		{
			LockId:              1,
			IntermediaryAccount: "fury1hpgapnfl3thkevvl0jp3wqtk8jw7mpqumuuc2f",
		},
	},
}

func TestMarshalUnmarshalGenesis(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	ctx = ctx.WithBlockTime(now.Add(time.Second))

	encodingConfig := simapp.MakeEncodingConfig()
	appCodec := encodingConfig.Marshaler
	am := superfluid.NewAppModule(*app.SuperfluidKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.LockupKeeper, app.GAMMKeeper, app.EpochsKeeper, app.ConcentratedLiquidityKeeper)
	genesis := testGenesis
	app.SuperfluidKeeper.InitGenesis(ctx, genesis)

	genesisExported := am.ExportGenesis(ctx, appCodec)
	assert.NotPanics(t, func() {
		app := simapp.Setup(false)
		ctx := app.BaseApp.NewContext(false, tmproto.Header{})
		ctx = ctx.WithBlockTime(now.Add(time.Second))
		am := superfluid.NewAppModule(*app.SuperfluidKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.LockupKeeper, app.GAMMKeeper, app.EpochsKeeper, app.ConcentratedLiquidityKeeper)
		am.InitGenesis(ctx, appCodec, genesisExported)
	})
}

func TestInitGenesis(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	ctx = ctx.WithBlockTime(now.Add(time.Second))
	genesis := testGenesis
	app.SuperfluidKeeper.InitGenesis(ctx, genesis)

	params := app.SuperfluidKeeper.GetParams(ctx)
	require.Equal(t, params, genesis.Params)

	assets := app.SuperfluidKeeper.GetAllSuperfluidAssets(ctx)
	require.Equal(t, assets, genesis.SuperfluidAssets)

	multipliers := app.SuperfluidKeeper.GetAllFuryEquivalentMultipliers(ctx)
	require.Equal(t, multipliers, genesis.FuryEquivalentMultipliers)

	accounts := app.SuperfluidKeeper.GetAllIntermediaryAccounts(ctx)
	require.Equal(t, accounts, genesis.IntermediaryAccounts)

	connections := app.SuperfluidKeeper.GetAllLockIdIntermediaryAccountConnections(ctx)
	require.Equal(t, connections, genesis.IntemediaryAccountConnections)
}

func TestExportGenesis(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	ctx = ctx.WithBlockTime(now.Add(time.Second))
	genesis := testGenesis
	app.SuperfluidKeeper.InitGenesis(ctx, genesis)

	asset := types.SuperfluidAsset{
		Denom:     "gamm/pool/2",
		AssetType: types.SuperfluidAssetTypeLPShare,
	}
	app.SuperfluidKeeper.SetSuperfluidAsset(ctx, asset)
	savedAsset, err := app.SuperfluidKeeper.GetSuperfluidAsset(ctx, "gamm/pool/2")
	require.NoError(t, err)
	require.Equal(t, savedAsset, asset)

	genesisExported := app.SuperfluidKeeper.ExportGenesis(ctx)
	require.Equal(t, genesisExported.Params, genesis.Params)
	require.Equal(t, genesisExported.SuperfluidAssets, append(genesis.SuperfluidAssets, asset))
	require.Equal(t, genesis.FuryEquivalentMultipliers, genesis.FuryEquivalentMultipliers)
	require.Equal(t, genesis.IntermediaryAccounts, genesis.IntermediaryAccounts)
	require.Equal(t, genesis.IntemediaryAccountConnections, genesis.IntemediaryAccountConnections)
}
