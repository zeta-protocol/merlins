package superfluid

import (
	"github.com/merlins-labs/merlins/v15/x/superfluid/keeper"
	"github.com/merlins-labs/merlins/v15/x/superfluid/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker is called on every block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper, ek types.EpochKeeper) {
	numBlocksSinceEpochStart, err := ek.NumBlocksSinceEpochStart(ctx, k.GetEpochIdentifier(ctx))
	if err != nil {
		panic(err)
	}
	if numBlocksSinceEpochStart == 0 {
		k.AfterEpochStartBeginBlock(ctx)
	}
}
