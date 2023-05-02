#!/bin/bash

# 1) this script creates an empty directory in app/upgrades called "vX" where X is a previous version + 1 with an empty upgrade handler.
# 2) adds new version to app.go
# 3) update MERLINS_E2E_UPGRADE_VERSION variable in .vscode/launch.json
# 4) increases E2E_UPGRADE_VERSION in makefile by 1
# 5) bumps up previous e2e-init version in tests/e2e/containers/config.go

# Also insures that all the imports make use of a current module version from go mod:
# (see:    module=$(go mod edit -json | jq ".Module.Path")      in this script)
# Github workflow which calls this script can be found here: merlins/.github/workflows/auto-update-upgrade.yml

latest_version=0
for f in app/upgrades/*; do 
	s_f=(${f//// })
	version=${s_f[2]}
	num_version=${version//[!0-9]/}
	if [[ $num_version -gt $latest_version ]]; then
		LATEST_FILE=$f
		latest_version=$num_version
	fi
done
version_create=$1
new_file=./app/upgrades/${version_create}

mkdir $new_file
CONSTANTS_FILE=$new_file/constants.go
UPGRADES_FILE=$new_file/upgrades.go
touch $CONSTANTS_FILE
touch $UPGRADES_FILE

module=$(go mod edit -json | jq ".Module.Path")
module=${module%?}
path=${module%???}

bracks='"'
# set packages
echo -e "package ${version_create}\n" >> $CONSTANTS_FILE
echo -e "package ${version_create}\n" >> $UPGRADES_FILE

# imports
echo "import (" >> $CONSTANTS_FILE
echo "import (" >> $UPGRADES_FILE

# set imports for constants.go
echo -e "\t$module/app/upgrades$bracks\n" >> $CONSTANTS_FILE
echo -e '\tstore "github.com/cosmos/cosmos-sdk/store/types"' >> $CONSTANTS_FILE

# set imports for upgrades.go
echo -e '\tsdk "github.com/cosmos/cosmos-sdk/types"' >> $UPGRADES_FILE
echo -e '\t"github.com/cosmos/cosmos-sdk/types/module"' >> $UPGRADES_FILE
echo -e '\tupgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"\n' >> $UPGRADES_FILE
echo -e "\t$module/app/keepers$bracks" >> $UPGRADES_FILE
echo -e "\t$module/app/upgrades$bracks" >> $UPGRADES_FILE

# close import
echo ")" >> $UPGRADES_FILE
echo -e ")\n" >> $CONSTANTS_FILE
 
# constants.go logic
echo "// UpgradeName defines the on-chain upgrade name for the Merlins $version_create upgrade." >> $CONSTANTS_FILE
echo "const UpgradeName = ${bracks}$version_create$bracks" >> $CONSTANTS_FILE
echo "
var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
    },
}" >> $CONSTANTS_FILE
 
# upgrades.go logic
echo "
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	bpm upgrades.BaseAppParamManager,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// Run migrations before applying any other state changes.
		// NOTE: DO NOT PUT ANY STATE CHANGES BEFORE RunMigrations().
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		return migrations, nil
	}
}" >> $UPGRADES_FILE

# change app/app.go file
app_file=./app/app.go
UPGRADES_LINE=$(grep -F upgrades.Upgrade{ $app_file)
UPGRADES_LINE="${UPGRADES_LINE%?}, ${version_create}.Upgrade}"
sed -i "s|.*upgrades.Upgrade{.*|$UPGRADES_LINE|" $app_file 

PREV_IMPORT="v$latest_version $module/app/upgrades/v$latest_version$bracks"
NEW_IMPORT="$version_create $module/app/upgrades/$version_create$bracks"
sed -i "s|.*$PREV_IMPORT.*|\t$PREV_IMPORT\n\t$NEW_IMPORT|" $app_file

# change e2e version in makefile
sed -i "s/E2E_UPGRADE_VERSION := ${bracks}v$latest_version$bracks/E2E_UPGRADE_VERSION := ${bracks}$version_create$bracks/" ./Makefile

# bumps up prev e2e version
e2e_file=./tests/e2e/containers/config.go
PREV_MERLINS_DEV_TAG=$(curl -L -s 'https://registry.hub.docker.com/v2/repositories/furylabs/merlins-dev/tags?page=1&page_size=100'            | jq -r '.results[] | .name | select(.|test("^(?:v|)[0-9]+\\.[0-9]+(?:$|\\.[0-9]+$)"))' | grep --max-count=1 "")
PREV_MERLINS_E2E_TAG=$(curl -L -s 'https://registry.hub.docker.com/v2/repositories/furylabs/merlins-e2e-init-chain/tags?page=1&page_size=100' | jq -r '.results[] | .name | select(.|test("^(?:v|)[0-9]+\\.[0-9]+(?:$|\\.[0-9]+$)"))' | grep --max-count=1 "")

# previousVersionFuryTag  = PREV_MERLINS_DEV_TAG
if [[ $version_create == v$(($(echo $PREV_MERLINS_DEV_TAG | awk -F . '{print $1}')+1)) ]]; then	
    echo "Found previous merlins-dev tag $PREV_MERLINS_DEV_TAG"
	sed -i '/previousVersionFuryTag/s/".*"/'"\"$PREV_MERLINS_DEV_TAG\""'/' $e2e_file
else
    PREV_MERLINS_DEV_TAG=v$((${version_create:1}-1)).0.0
    echo "Using pre-defined merlins-dev tag: $PREV_MERLINS_DEV_TAG"
    sed -i '/previousVersionFuryTag/s/".*"/'"\"$PREV_MERLINS_DEV_TAG\""'/' $e2e_file
fi

# previousVersionInitTag  = PREV_MERLINS_E2E_TAG
if [[ $version_create == v$(($(echo $PREV_MERLINS_E2E_TAG | awk -F . '{print $1}' | grep -Eo '[0-9]*')+1)) ]]; then	
    echo "Found previous merlins-e2e-init-chain tag $PREV_MERLINS_E2E_TAG"
	sed -i '/previousVersionInitTag/s/".*"/'"\"$PREV_MERLINS_E2E_TAG\""'/' $e2e_file
else
    PREV_MERLINS_E2E_TAG=v$((${version_create:1}-1)).0.0
    echo "Using pre-defined merlins-e2e-init-chain tag: $PREV_MERLINS_E2E_TAG"
    sed -i '/previousVersionInitTag/s/".*"/'"\"$PREV_MERLINS_E2E_TAG\""'/' $e2e_file
fi

# update MERLINS_E2E_UPGRADE_VERSION in launch.json
sed -i "s/${bracks}MERLINS_E2E_UPGRADE_VERSION${bracks}: ${bracks}v$latest_version${bracks}/${bracks}MERLINS_E2E_UPGRADE_VERSION${bracks}: ${bracks}$version_create${bracks}/" ./.vscode/launch.json
