#!/bin/bash



# Archive node script
# NB:  you can also download archives at quicksync:
# https://quicksync.io/networks/merlins.html
# 2nd NB: you can change MERLINSD_PRUNING=nothing to MERLINSD_PRUNING=default OR you could also set the pruning settings manually with MERLINSD_PRUNING=custom
# 3rd NB: you might want to use this to test different databases, and to do that my recommended technique is like:
# go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb' -tags rocksdb ./...
# if you do not use the ldflags thing you won't use the chosen db for everything, so best use it.


# The goal of this script is to provide a way to audit:
# * Data availability
# * Database performance when synchronizing (will add variables for different DB's after a successful run)
# * Size on disk for archives when using different databases


export MERLINSD_PRUNING=nothing
export MERLINSD_DB_BACKEND=pebbledb
export MERLINSD_P2P_MAX_NUM_OUTBOUND_PEERS=500
export MERLINSD_P2P_MAX_NUM_INBOUND_PEERS=500
export MERLINSD_P2P_SEEDS=$(curl -s https://raw.githubusercontent.com/cosmos/chain-registry/master/merlins/chain.json | jq -r '[foreach .peers.seeds[] as $item (""; "\($item.id)@\($item.address)")] | join(",")')
export MERLINSD_P2P_LADDR=tcp://0.0.0.0:2001

# VERSION THREE
echo "v3 took" > howlong
git checkout v3.x
go mod edit -replace github.com/tendermint/tm-db=github.com/notional-labs/tm-db@136c7b6
go mod tidy
go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb' -tags pebbledb ./...
merlinsd init speedrun
wget -O ~/.merlinsd/config/genesis.json https://github.com/merlins-labs/networks/raw/main/merlins-1/genesis.json
merlinsd start --db_backend pebbledb
git reset --hard

# VERSION FOUR
echo "v4 took" >> howlong
git checkout v4.x
go mod edit -replace github.com/tendermint/tm-db=github.com/notional-labs/tm-db@136c7b6
go mod tidy
go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb' -tags pebbledb ./...
merlinsd start --db_backend pebbledb
git reset --hard


# VERSION SIX
echo "v6 took" >> howlong
git checkout v6.x
go mod edit -replace github.com/tendermint/tm-db=github.com/notional-labs/tm-db@136c7b6
go mod tidy
go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb' -tags pebbledb ./...
merlinsd start --db_backend pebbledb
git reset --hard



# VERSION SEVEN
echo "v7 took" >> howlong
git checkout v7.x
go mod edit -replace github.com/tendermint/tm-db=github.com/notional-labs/tm-db@136c7b6
go mod tidy
go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb' -tags pebbledb ./...
merlinsd start --db_backend pebbledb
git reset --hard



# VERSION EIGHT
echo "v8 took" >> howlong
git checkout v8.x
go mod edit -replace github.com/tendermint/tm-db=github.com/notional-labs/tm-db@136c7b6
go mod tidy
go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb' -tags pebbledb ./...
merlinsd start --db_backend pebbledb
git reset --hard


# VERSION NINE
echo "v9 took" >> howlong
git checkout v9.x
go mod edit -replace github.com/tendermint/tm-db=github.com/notional-labs/tm-db@136c7b6
go mod tidy
go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb' -tags pebbledb ./...time merlinsd start --db_backend pebbledb
merlinsd start --db_backend pebbledb
git reset --hard


# VERSION TEN
echo "v10 took" >> howlong
git checkout v10.x
go mod edit -replace github.com/tendermint/tm-db=github.com/notional-labs/tm-db@136c7b6
go mod tidy
go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb' -tags pebbledb ./...
merlinsd start --db_backend pebbledb

