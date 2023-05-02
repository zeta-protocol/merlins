package model

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	"github.com/merlins-labs/merlins/v15/x/concentrated-liquidity/types"
	poolmanagertypes "github.com/merlins-labs/merlins/v15/x/poolmanager/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&Pool{}, "merlins/cl-pool", nil)
	cdc.RegisterConcrete(&MsgCreateConcentratedPool{}, "merlins/cl-create-pool", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"merlins.swaprouter.v1beta1.PoolI",
		(*poolmanagertypes.PoolI)(nil),
		&Pool{},
	)

	registry.RegisterInterface(
		"merlins.concentratedliquidity.v1beta1.ConcentratedPoolExtension",
		(*types.ConcentratedPoolExtension)(nil),
		&Pool{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateConcentratedPool{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_MsgCreator_serviceDesc)
}

// TODO: re-enable this when CL state-breakage PR is merged.
// return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterCodec(authzcodec.Amino)
	amino.Seal()
}
