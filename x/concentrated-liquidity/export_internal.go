package concentrated_liquidity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/merlins-labs/merlins/v15/x/concentrated-liquidity/math"
)

func TickToSqrtPrice(tickIndex sdk.Int) (price sdk.Dec, err error) {
	return math.TickToSqrtPrice(tickIndex)
}
