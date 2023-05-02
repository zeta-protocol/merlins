package v6

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/merlins-labs/merlins/v15/app/keepers"
)

// RunForkLogic executes height-gated on-chain fork logic for the Merlins v6
// upgrade.
//
// NOTE: All the height gated fork logic is actually in the Merlins ibc-go fork.
// See: https://github.com/merlins-labs/ibc-go/releases/tag/v2.0.2-fury
func RunForkLogic(ctx sdk.Context, _ *keepers.AppKeepers) {
	ctx.Logger().Info("Applying emergency hard fork for v6, allows IBC to create new channels.")
}
