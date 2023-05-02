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
	cdc.RegisterConcrete(&MsgSuperfluidDelegate{}, "merlins/superfluid-delegate", nil)
	cdc.RegisterConcrete(&MsgSuperfluidUndelegate{}, "merlins/superfluid-undelegate", nil)
	cdc.RegisterConcrete(&MsgLockAndSuperfluidDelegate{}, "merlins/lock-and-superfluid-delegate", nil)
	cdc.RegisterConcrete(&MsgSuperfluidUnbondLock{}, "merlins/superfluid-unbond-lock", nil)
	cdc.RegisterConcrete(&MsgSuperfluidUndelegateAndUnbondLock{}, "merlins/sf-undelegate-and-unbond-lock", nil)
	cdc.RegisterConcrete(&SetSuperfluidAssetsProposal{}, "merlins/set-superfluid-assets-proposal", nil)
	cdc.RegisterConcrete(&UpdateUnpoolWhiteListProposal{}, "merlins/update-unpool-whitelist", nil)
	cdc.RegisterConcrete(&RemoveSuperfluidAssetsProposal{}, "merlins/del-superfluid-assets-proposal", nil)
	cdc.RegisterConcrete(&MsgUnPoolWhitelistedPool{}, "merlins/unpool-whitelisted-pool", nil)
	cdc.RegisterConcrete(&MsgUnlockAndMigrateSharesToFullRangeConcentratedPosition{}, "merlins/unlock-and-migrate", nil)
	cdc.RegisterConcrete(&MsgCreateFullRangePositionAndSuperfluidDelegate{}, "merlins/full-range-and-sf-delegate", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSuperfluidDelegate{},
		&MsgSuperfluidUndelegate{},
		&MsgLockAndSuperfluidDelegate{},
		&MsgSuperfluidUnbondLock{},
		&MsgSuperfluidUndelegateAndUnbondLock{},
		&MsgUnPoolWhitelistedPool{},
		&MsgUnlockAndMigrateSharesToFullRangeConcentratedPosition{},
		&MsgCreateFullRangePositionAndSuperfluidDelegate{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&SetSuperfluidAssetsProposal{},
		&RemoveSuperfluidAssetsProposal{},
		&UpdateUnpoolWhiteListProposal{},
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
