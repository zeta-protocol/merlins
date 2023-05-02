package v8

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/merlins-labs/merlins/v15/app/keepers"
)

// RunForkLogic executes height-gated on-chain fork logic for the Merlins v8
// upgrade.
func RunForkLogic(ctx sdk.Context, appKeepers *keepers.AppKeepers) {
	// Only proceed with v8 for mainnet, testnets need not adjust their pool incentives or unbonding.
	// https://github.com/merlins-labs/merlins/issues/1609
	if ctx.ChainID() != "merlins-1" {
		return
	}

	for i := 0; i < 100; i++ {
		ctx.Logger().Info("I am upgrading to v8")
	}
	ctx.Logger().Info("Applying Merlins v8 upgrade. Allowing direct unpooling for whitelisted pools")
	ctx.Logger().Info("Applying accelerated incentive updates per proposal 225")
	ApplyProp222Change(ctx, appKeepers.PoolIncentivesKeeper)
	ApplyProp223Change(ctx, appKeepers.PoolIncentivesKeeper)
	ApplyProp224Change(ctx, appKeepers.PoolIncentivesKeeper)
	ctx.Logger().Info("Registering state change for whitelisted pools for unpooling ")
	RegisterWhitelistedDirectUnbondPools(ctx, appKeepers.SuperfluidKeeper, appKeepers.GAMMKeeper)
}
