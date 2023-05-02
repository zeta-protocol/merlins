# this script runs under the assumption that a three-validator environment is running on your local machine(multinode-local-testnet.sh)
# this script would do basic setup that has to be achieved to actual superfluid staking
# prior to running this script, have the following json file in the directory running this script
#
# stake-ufury.json
# {
# 	"weights": "5stake,5ufury",
# 	"initial-deposit": "1000000stake,1000000ufury",
# 	"swap-fee": "0.01",
# 	"exit-fee": "0.01",
# 	"future-governor": "168h"
# }

# create pool
merlinsd tx gamm create-pool --pool-file=./stake-ufury.json --from=validator1 --keyring-backend=test --chain-id=testing --yes --home=$HOME/.merlinsd/validator1
sleep 7

# test swap in pool created
merlinsd tx gamm swap-exact-amount-in 100000ufury 50000 --swap-route-pool-ids=1 --swap-route-denoms=stake --from=validator1 --keyring-backend=test --chain-id=testing --yes --home=$HOME/.merlinsd/validator1
sleep 7

# create a lock up with lockable duration 360h
merlinsd tx lockup lock-tokens 10000000000000000000gamm/pool/1 --duration=360h --from=validator1 --keyring-backend=test --chain-id=testing --broadcast-mode=block --yes --home=$HOME/.merlinsd/validator1
sleep 7

# submit and pass proposal for superfluid
merlinsd tx gov submit-proposal set-superfluid-assets-proposal --title="set superfluid assets" --description="set superfluid assets description" --superfluid-assets="gamm/pool/1" --deposit=10000000ufury --from=validator1 --chain-id=testing --keyring-backend=test --broadcast-mode=block --yes --home=$HOME/.merlinsd/validator1
sleep 7

merlinsd tx gov deposit 1 10000000stake --from=validator1 --keyring-backend=test --chain-id=testing --broadcast-mode=block --yes --home=$HOME/.merlinsd/validator1
sleep 7

merlinsd tx gov vote 1 yes --from=validator1 --keyring-backend=test --chain-id=testing --yes --home=$HOME/.merlinsd/validator1
sleep 7
merlinsd tx gov vote 1 yes --from=validator2 --keyring-backend=test --chain-id=testing --yes --home=$HOME/.merlinsd/validator2
sleep 7
