package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*ConcentratedPoolExtension)(nil), nil)
	cdc.RegisterConcrete(&MsgCreatePosition{}, "merlins/cl-create-position", nil)
	cdc.RegisterConcrete(&MsgWithdrawPosition{}, "merlins/cl-withdraw-position", nil)
	cdc.RegisterConcrete(&MsgCollectFees{}, "merlins/cl-collect-fees", nil)
	cdc.RegisterConcrete(&MsgCollectIncentives{}, "merlins/cl-collect-incentives", nil)
	cdc.RegisterConcrete(&MsgCreateIncentive{}, "merlins/cl-create-incentive", nil)
	cdc.RegisterConcrete(&MsgFungifyChargedPositions{}, "merlins/cl-fungify-charged-positions", nil)
	cdc.RegisterConcrete(&TickSpacingDecreaseProposal{}, "merlins/cl-tick-spacing-dec-prop", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"merlins.concentratedliquidity.v1beta1.ConcentratedPoolExtension",
		(*ConcentratedPoolExtension)(nil),
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreatePosition{},
		&MsgWithdrawPosition{},
		&MsgCollectFees{},
		&MsgCollectIncentives{},
		&MsgCreateIncentive{},
		&MsgFungifyChargedPositions{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&TickSpacingDecreaseProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

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
