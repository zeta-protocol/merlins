package valsetprefcli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/merlins-labs/merlins/furyutils"
	"github.com/merlins-labs/merlins/furyutils/furycli"
	"github.com/merlins-labs/merlins/v15/x/valset-pref/types"
)

func GetTxCmd() *cobra.Command {
	txCmd := furycli.TxIndexCmd(types.ModuleName)
	furycli.AddTxCmd(txCmd, NewSetValSetCmd)
	furycli.AddTxCmd(txCmd, NewDelValSetCmd)
	furycli.AddTxCmd(txCmd, NewUnDelValSetCmd)
	furycli.AddTxCmd(txCmd, NewReDelValSetCmd)
	furycli.AddTxCmd(txCmd, NewWithRewValSetCmd)
	return txCmd
}

func NewSetValSetCmd() (*furycli.TxCliDesc, *types.MsgSetValidatorSetPreference) {
	return &furycli.TxCliDesc{
		Use:              "set-valset [delegator_addr] [validators] [weights]",
		Short:            "Creates a new validator set for the delegator with valOperAddress and weight",
		Example:          "merlinsd tx valset-pref set-valset fury1... furyvaloper1abc...,furyvaloper1def...  0.56,0.44",
		NumArgs:          3,
		ParseAndBuildMsg: NewMsgSetValidatorSetPreference,
	}, &types.MsgSetValidatorSetPreference{}
}

func NewDelValSetCmd() (*furycli.TxCliDesc, *types.MsgDelegateToValidatorSet) {
	return &furycli.TxCliDesc{
		Use:     "delegate-valset [delegator_addr] [amount]",
		Short:   "Delegate tokens to existing valset using delegatorAddress and tokenAmount.",
		Example: "merlinsd tx valset-pref delegate-valset fury1... 100stake",
		NumArgs: 2,
	}, &types.MsgDelegateToValidatorSet{}
}

func NewUnDelValSetCmd() (*furycli.TxCliDesc, *types.MsgUndelegateFromValidatorSet) {
	return &furycli.TxCliDesc{
		Use:     "undelegate-valset [delegator_addr] [amount]",
		Short:   "UnDelegate tokens from existing valset using delegatorAddress and tokenAmount.",
		Example: "merlinsd tx valset-pref undelegate-valset fury1... 100stake",
		NumArgs: 2,
	}, &types.MsgUndelegateFromValidatorSet{}
}

func NewReDelValSetCmd() (*furycli.TxCliDesc, *types.MsgRedelegateValidatorSet) {
	return &furycli.TxCliDesc{
		Use:              "redelegate-valset [delegator_addr] [validators] [weights]",
		Short:            "Redelegate tokens from existing validators to new sets of validators",
		Example:          "merlinsd tx valset-pref redelegate-valset  fury1... furyvaloper1efg...,furyvaloper1hij...  0.56,0.44",
		NumArgs:          3,
		ParseAndBuildMsg: NewMsgReDelValidatorSetPreference,
	}, &types.MsgRedelegateValidatorSet{}
}

func NewWithRewValSetCmd() (*furycli.TxCliDesc, *types.MsgWithdrawDelegationRewards) {
	return &furycli.TxCliDesc{
		Use:     "withdraw-reward-valset [delegator_addr]",
		Short:   "Withdraw delegation reward form the new validator set.",
		Example: "merlinsd tx valset-pref withdraw-valset fury1...",
		NumArgs: 1,
	}, &types.MsgWithdrawDelegationRewards{}
}

func NewMsgSetValidatorSetPreference(clientCtx client.Context, args []string, fs *pflag.FlagSet) (sdk.Msg, error) {
	delAddr, err := sdk.AccAddressFromBech32(args[0])
	if err != nil {
		return nil, err
	}

	valset, err := ValidateValAddrAndWeight(args)
	if err != nil {
		return nil, err
	}

	return types.NewMsgSetValidatorSetPreference(
		delAddr,
		valset,
	), nil
}

func NewMsgReDelValidatorSetPreference(clientCtx client.Context, args []string, fs *pflag.FlagSet) (sdk.Msg, error) {
	delAddr, err := sdk.AccAddressFromBech32(args[0])
	if err != nil {
		return nil, err
	}

	valset, err := ValidateValAddrAndWeight(args)
	if err != nil {
		return nil, err
	}

	return types.NewMsgRedelegateValidatorSet(
		delAddr,
		valset,
	), nil
}

func ValidateValAddrAndWeight(args []string) ([]types.ValidatorPreference, error) {
	var valAddrs []string
	valAddrs = append(valAddrs, strings.Split(args[1], ",")...)

	weights, err := furyutils.ParseSdkDecFromString(args[2], ",")
	if err != nil {
		return nil, err
	}

	if len(valAddrs) != len(weights) {
		return nil, fmt.Errorf("the length of validator addresses and weights not matched")
	}

	if len(valAddrs) == 0 {
		return nil, fmt.Errorf("records is empty")
	}

	var valset []types.ValidatorPreference
	for i, val := range valAddrs {
		valset = append(valset, types.ValidatorPreference{
			ValOperAddress: val,
			Weight:         weights[i],
		})
	}

	return valset, nil
}
