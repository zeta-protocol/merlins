package keeper

import (
	"github.com/gogo/protobuf/proto"

	gammtypes "github.com/merlins-labs/merlins/v15/x/gamm/types"
	"github.com/merlins-labs/merlins/v15/x/superfluid/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// This function calculates the fury equivalent worth of an LP share.
// It is intended to eventually use the TWAP of the worth of an LP share
// once that is exposed from the gamm module.
func (k Keeper) calculateFuryBackingPerShare(pool gammtypes.CFMMPoolI, furyInPool sdk.Int) sdk.Dec {
	twap := furyInPool.ToDec().Quo(pool.GetTotalShares().ToDec())
	return twap
}

func (k Keeper) SetFuryEquivalentMultiplier(ctx sdk.Context, epoch int64, denom string, multiplier sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixTokenMultiplier)
	priceRecord := types.FuryEquivalentMultiplierRecord{
		EpochNumber: epoch,
		Denom:       denom,
		Multiplier:  multiplier,
	}
	bz, err := proto.Marshal(&priceRecord)
	if err != nil {
		panic(err)
	}
	prefixStore.Set([]byte(denom), bz)
}

func (k Keeper) GetSuperfluidFURYTokens(ctx sdk.Context, denom string, amount sdk.Int) (sdk.Int, error) {
	multiplier := k.GetFuryEquivalentMultiplier(ctx, denom)
	if multiplier.IsZero() {
		return sdk.ZeroInt(), nil
	}

	decAmt := multiplier.Mul(amount.ToDec())
	_, err := k.GetSuperfluidAsset(ctx, denom)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	return k.GetRiskAdjustedFuryValue(ctx, decAmt.RoundInt()), nil
}

func (k Keeper) DeleteFuryEquivalentMultiplier(ctx sdk.Context, denom string) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixTokenMultiplier)
	prefixStore.Delete([]byte(denom))
}

func (k Keeper) GetFuryEquivalentMultiplier(ctx sdk.Context, denom string) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixTokenMultiplier)
	bz := prefixStore.Get([]byte(denom))
	if bz == nil {
		return sdk.ZeroDec()
	}
	priceRecord := types.FuryEquivalentMultiplierRecord{}
	err := proto.Unmarshal(bz, &priceRecord)
	if err != nil {
		panic(err)
	}
	return priceRecord.Multiplier
}

func (k Keeper) GetAllFuryEquivalentMultipliers(ctx sdk.Context) []types.FuryEquivalentMultiplierRecord {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixTokenMultiplier)
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()

	priceRecords := []types.FuryEquivalentMultiplierRecord{}
	for ; iterator.Valid(); iterator.Next() {
		priceRecord := types.FuryEquivalentMultiplierRecord{}

		err := proto.Unmarshal(iterator.Value(), &priceRecord)
		if err != nil {
			panic(err)
		}

		priceRecords = append(priceRecords, priceRecord)
	}
	return priceRecords
}
