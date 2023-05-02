package types

import (
	"github.com/gogo/protobuf/proto"

	poolmanagertypes "github.com/merlins-labs/merlins/v15/x/poolmanager/types"
)

// CosmWasmExtension
type CosmWasmExtension interface {
	poolmanagertypes.PoolI

	GetCodeId() uint64

	GetInstantiateMsg() []byte

	GetContractAddress() string

	SetContractAddress(contractAddress string)

	GetStoreModel() proto.Message

	SetWasmKeeper(wasmKeeper WasmKeeper)
}
