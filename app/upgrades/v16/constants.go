package v16

import (
	"github.com/merlins-labs/merlins/v15/app/upgrades"

	store "github.com/cosmos/cosmos-sdk/store/types"

	cltypes "github.com/merlins-labs/merlins/v15/x/concentrated-liquidity/types"
)

// UpgradeName defines the on-chain upgrade name for the Merlins v16 upgrade.
const UpgradeName = "v16"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{cltypes.StoreKey},
		Deleted: []string{},
	},
}
