package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/merlins-labs/merlins/furyutils/furycli"
	"github.com/merlins-labs/merlins/v15/x/lockup/types"
)

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := furycli.TxIndexCmd(types.ModuleName)
	furycli.AddTxCmd(cmd, NewLockTokensCmd)
	furycli.AddTxCmd(cmd, NewBeginUnlockingAllCmd)
	furycli.AddTxCmd(cmd, NewBeginUnlockByIDCmd)
	furycli.AddTxCmd(cmd, NewForceUnlockByIdCmd)

	return cmd
}

func NewLockTokensCmd() (*furycli.TxCliDesc, *types.MsgLockTokens) {
	return &furycli.TxCliDesc{
		Use:   "lock-tokens [tokens]",
		Short: "lock tokens into lockup pool from user account",
		CustomFlagOverrides: map[string]string{
			"duration": FlagDuration,
		},
		Flags: furycli.FlagDesc{RequiredFlags: []*pflag.FlagSet{FlagSetLockTokens()}},
	}, &types.MsgLockTokens{}
}

// TODO: We should change the Use string to be unlock-all
func NewBeginUnlockingAllCmd() (*furycli.TxCliDesc, *types.MsgBeginUnlockingAll) {
	return &furycli.TxCliDesc{
		Use:   "begin-unlock-tokens",
		Short: "begin unlock not unlocking tokens from lockup pool for sender",
	}, &types.MsgBeginUnlockingAll{}
}

// NewBeginUnlockByIDCmd unlocks individual period lock by ID.
func NewBeginUnlockByIDCmd() (*furycli.TxCliDesc, *types.MsgBeginUnlocking) {
	return &furycli.TxCliDesc{
		Use:   "begin-unlock-by-id [id]",
		Short: "begin unlock individual period lock by ID",
		CustomFlagOverrides: map[string]string{
			"coins": FlagAmount,
		},
		Flags: furycli.FlagDesc{OptionalFlags: []*pflag.FlagSet{FlagSetUnlockTokens()}},
	}, &types.MsgBeginUnlocking{}
}

// NewForceUnlockByIdCmd force unlocks individual period lock by ID if proper permissions exist.
func NewForceUnlockByIdCmd() (*furycli.TxCliDesc, *types.MsgForceUnlock) {
	return &furycli.TxCliDesc{
		Use:   "force-unlock-by-id [id]",
		Short: "force unlocks individual period lock by ID",
		Long:  "force unlocks individual period lock by ID. if no amount provided, entire lock is unlocked",
		CustomFlagOverrides: map[string]string{
			"coins": FlagAmount,
		},
		Flags: furycli.FlagDesc{OptionalFlags: []*pflag.FlagSet{FlagSetUnlockTokens()}},
	}, &types.MsgForceUnlock{}
}
