package v3

import "github.com/merlins-labs/merlins/v15/app/upgrades"

const (
	// UpgradeName defines the on-chain upgrade name for the Merlins v3 upgrade.
	UpgradeName = "v3"

	// UpgradeHeight defines the block height at which the Merlins v3 upgrade is
	// triggered.
	UpgradeHeight = 712_000
)

var Fork = upgrades.Fork{
	UpgradeName:    UpgradeName,
	UpgradeHeight:  UpgradeHeight,
	BeginForkLogic: RunForkLogic,
}
