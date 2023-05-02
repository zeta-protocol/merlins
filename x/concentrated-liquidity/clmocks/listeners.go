package clmocks

import (
	"github.com/merlins-labs/merlins/v15/x/concentrated-liquidity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ConcentratedLiquidityListenerMock struct {
	AfterConcentratedPoolCreatedCallCount    int
	AfterInitialPoolPositionCreatedCallCount int
	AfterLastPoolPositionRemovedCallCount    int
	AfterConcentratedPoolSwapCallCount       int
}

var _ types.ConcentratedLiquidityListener = &ConcentratedLiquidityListenerMock{}

func (l *ConcentratedLiquidityListenerMock) AfterConcentratedPoolCreated(ctx sdk.Context, sender sdk.AccAddress, poolId uint64) {
	l.AfterConcentratedPoolCreatedCallCount += 1
}

func (l *ConcentratedLiquidityListenerMock) AfterInitialPoolPositionCreated(ctx sdk.Context, sender sdk.AccAddress, poolId uint64) {
	l.AfterInitialPoolPositionCreatedCallCount += 1
}

func (l *ConcentratedLiquidityListenerMock) AfterLastPoolPositionRemoved(ctx sdk.Context, sender sdk.AccAddress, poolId uint64) {
	l.AfterLastPoolPositionRemovedCallCount += 1
}

func (l *ConcentratedLiquidityListenerMock) AfterConcentratedPoolSwap(ctx sdk.Context, sender sdk.AccAddress, poolId uint64) {
	l.AfterConcentratedPoolSwapCallCount += 1
}
