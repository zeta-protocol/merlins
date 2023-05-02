# run old binary on terminal1
go clean --modcache
git stash
git checkout v1.0.1
go install ./cmd/merlinsd/
Run the below commands
```
    #!/bin/bash
    rm -rf $HOME/.merlinsd/
    cd $HOME
    merlinsd init --chain-id=testing testing --home=$HOME/.merlinsd
    merlinsd keys add validator --keyring-backend=test --home=$HOME/.merlinsd
    merlinsd add-genesis-account $(merlinsd keys show validator -a --keyring-backend=test --home=$HOME/.merlinsd) 1000000000stake,1000000000valtoken --home=$HOME/.merlinsd
    merlinsd gentx validator 500000000stake --keyring-backend=test --home=$HOME/.merlinsd --chain-id=testing
    merlinsd gentx validator 500000000stake --commission-rate="0.0" --keyring-backend=test --home=$HOME/.merlinsd --chain-id=testing
    merlinsd collect-gentxs --home=$HOME/.merlinsd
    
    cat $HOME/.merlinsd/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="120s"' > $HOME/.merlinsd/config/tmp_genesis.json && mv $HOME/.merlinsd/config/tmp_genesis.json $HOME/.merlinsd/config/genesis.json
    cat $HOME/.merlinsd/config/genesis.json | jq '.app_state["staking"]["params"]["min_commission_rate"]="0.050000000000000000"' > $HOME/.merlinsd/config/tmp_genesis.json && mv $HOME/.merlinsd/config/tmp_genesis.json $HOME/.merlinsd/config/genesis.json

```

Create pool.json 
```
{
  "weights": "1stake,1valtoken",
  "initial-deposit": "100stake,20valtoken",
  "swap-fee": "0.01",
  "exit-fee": "0.01",
  "future-governor": "168h"
}
```

rm $HOME/.merlinsd/cfuryvisor/current -rf
cfuryvisor start

# operations on terminal2
merlinsd tx lockup lock-tokens 100stake --duration="5s" --from=validator --chain-id=testing --keyring-backend=test --yes
sleep 7
merlinsd tx gov submit-proposal software-upgrade upgrade-lockup-module-store-management --title="lockup module upgrade" --description="lockup module upgrade for gas efficiency"  --from=validator --upgrade-height=10 --deposit=10000000stake --chain-id=testing --keyring-backend=test -y
sleep 7
merlinsd tx gov vote 1 yes --from=validator --keyring-backend=test --chain-id=testing --yes
sleep 7
merlinsd tx gamm create-pool --pool-file="./pool.json"  --gas=3000000 --from=validator --chain-id=testing --keyring-backend=test --yes --broadcast-mode=block
sleep 7
merlinsd tx lockup lock-tokens 1000stake --duration="100s" --from=validator --chain-id=testing --keyring-backend=test --yes --broadcast-mode=block
sleep 7
merlinsd tx lockup lock-tokens 2000stake --duration="200s" --from=validator --chain-id=testing --keyring-backend=test --yes --broadcast-mode=block
sleep 7
merlinsd tx lockup lock-tokens 3000stake --duration="1s" --from=validator --chain-id=testing --keyring-backend=test --yes --broadcast-mode=block
sleep 7
merlinsd tx lockup begin-unlock-by-id 1 --from=validator --chain-id=testing --keyring-backend=test --yes --broadcast-mode=block
sleep 7
merlinsd tx lockup begin-unlock-by-id 3 --from=validator --chain-id=testing --keyring-backend=test --yes --broadcast-mode=block
sleep 7
merlinsd tx gov submit-proposal software-upgrade "v2" --title="lockup module upgrade" --description="lockup module upgrade for gas efficiency"  --from=validator --upgrade-height=20 --deposit=10000000stake --chain-id=testing --keyring-backend=test --yes  --broadcast-mode=block
sleep 7
merlinsd tx gov vote 1 yes --from=validator --keyring-backend=test --chain-id=testing --yes --broadcast-mode=block
merlinsd query gov proposal 1
merlinsd query upgrade plan
merlinsd query lockup account-locked-longer-duration $(merlinsd keys show -a --keyring-backend=test validator) 1s
merlinsd query gamm pools
merlinsd query staking validators
merlinsd query staking params

# on terminal1
Wait until consensus failure happen and stop binary using Ctrl + C
git checkout lockup_module_genesis_export
git checkout main

Update go mod file to use latest SDK changes: /Users/admin/go/pkg/mod/github.com/merlins-labs/cosmos-sdk@v0.42.5-0.20210622202917-f4f6a08ac64b
go get github.com/merlins-labs/cosmos-sdk@ea1ec79c739ba39639b9a24f824127ecc6650887
go: downloading github.com/merlins-labs/cosmos-sdk v0.42.5-0.20210630100106-ea1ec79c739b
Upgrade Merlins Cosmos SDK version to `v0.42.5-0.20210630100106-ea1ec79c739b`
go mod download github.com/cosmos/cosmos-sdk
git stash
git checkout min_commission_change_validation_change_ignore
go install ./cmd/merlinsd/
merlinsd start --home=$HOME/.merlinsd

# check on terminal2
merlinsd query lockup account-locked-longer-duration $(merlinsd keys show -a --keyring-backend=test validator) 1s
merlinsd query lockup account-locked-longer-duration $(merlinsd keys show -a --keyring-backend=test validator) 1s
merlinsd query lockup module-locked-amount
merlinsd query gamm pools
merlinsd query staking validators
merlinsd query staking params
merlinsd query bank balances $(merlinsd keys show -a --keyring-backend=test validator)
merlinsd tx staking edit-validator --commission-rate="0.1"  --from=validator --chain-id=testing --keyring-backend=test --yes --broadcast-mode=block
merlinsd tx staking edit-validator --commission-rate="0.08"  --from=validator --chain-id=testing --keyring-backend=test --yes --broadcast-mode=block

Result:
- pool exists
- lockup processed all correctly
- validator commission rate worked
- chain did not panic 