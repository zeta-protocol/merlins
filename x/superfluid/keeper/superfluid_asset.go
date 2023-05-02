package keeper

import (
	"github.com/merlins-labs/merlins/furyutils"
	"github.com/merlins-labs/merlins/v15/x/superfluid/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginUnwindSuperfluidAsset starts the deletion process for a superfluid asset.
// This current method is a stub, but is called when:
// * Governance removes a superfluid asset
// * A severe error in gamm occurs
//
// It should eventually begin unwinding all of the synthetic lockups for that asset
// and queue them for deletion.
// See https://github.com/merlins-labs/merlins/issues/864
func (k Keeper) BeginUnwindSuperfluidAsset(ctx sdk.Context, epochNum int64, asset types.SuperfluidAsset) {
	// Right now set the TWAP to 0, and delete the asset.
	k.SetFuryEquivalentMultiplier(ctx, epochNum, asset.Denom, sdk.ZeroDec())
	k.DeleteSuperfluidAsset(ctx, asset.Denom)
}

// Returns amount * (1 - k.RiskFactor(asset))
// Fow now, the risk factor is a global constant.
// It will move towards per pool functions.
func (k Keeper) GetRiskAdjustedFuryValue(ctx sdk.Context, amount sdk.Int) sdk.Int {
	minRiskFactor := k.GetParams(ctx).MinimumRiskFactor
	return amount.Sub(amount.ToDec().Mul(minRiskFactor).RoundInt())
}

// y = x - (x * minRisk)
// y = x (1 - minRisk)
// y / (1 - minRisk) = x

func (k Keeper) UnriskAdjustFuryValue(ctx sdk.Context, amount sdk.Dec) sdk.Dec {
	minRiskFactor := k.GetParams(ctx).MinimumRiskFactor
	return amount.Quo(sdk.OneDec().Sub(minRiskFactor))
}

func (k Keeper) AddNewSuperfluidAsset(ctx sdk.Context, asset types.SuperfluidAsset) error {
	// initialize fury equivalent multipliers
	epochIdentifier := k.GetEpochIdentifier(ctx)
	currentEpoch := k.ek.GetEpochInfo(ctx, epochIdentifier).CurrentEpoch
	return furyutils.ApplyFuncIfNoError(ctx, func(ctx sdk.Context) error {
		k.SetSuperfluidAsset(ctx, asset)
		err := k.UpdateFuryEquivalentMultipliers(ctx, asset, currentEpoch)
		return err
	})
}
