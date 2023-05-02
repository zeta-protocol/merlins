package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	flag "github.com/spf13/pflag"

	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/merlins-labs/merlins/furyutils/furycli"
	clmodel "github.com/merlins-labs/merlins/v15/x/concentrated-liquidity/model"
	"github.com/merlins-labs/merlins/v15/x/concentrated-liquidity/types"
)

func NewTxCmd() *cobra.Command {
	txCmd := furycli.TxIndexCmd(types.ModuleName)
	furycli.AddTxCmd(txCmd, NewCreatePositionCmd)
	furycli.AddTxCmd(txCmd, NewWithdrawPositionCmd)
	furycli.AddTxCmd(txCmd, NewCreateConcentratedPoolCmd)
	furycli.AddTxCmd(txCmd, NewCollectFeesCmd)
	furycli.AddTxCmd(txCmd, NewCollectIncentivesCmd)
	furycli.AddTxCmd(txCmd, NewCreateIncentiveCmd)
	return txCmd
}

var poolIdFlagOverride = map[string]string{
	"poolid": FlagPoolId,
}

func NewCreateConcentratedPoolCmd() (*furycli.TxCliDesc, *clmodel.MsgCreateConcentratedPool) {
	return &furycli.TxCliDesc{
		Use:     "create-concentrated-pool [denom-0] [denom-1] [tick-spacing] [swap-fee]",
		Short:   "create a concentrated liquidity pool with the given denom pair, tick spacing, and swap fee",
		Long:    "denom-1 (the quote denom), tick spacing, and swap fees must all be authorized by the concentrated liquidity module",
		Example: "create-concentrated-pool uion ufury 1 0.01 --from val --chain-id merlins-1",
	}, &clmodel.MsgCreateConcentratedPool{}
}

func NewCreatePositionCmd() (*furycli.TxCliDesc, *types.MsgCreatePosition) {
	return &furycli.TxCliDesc{
		Use:                 "create-position [lower-tick] [upper-tick] [token-0] [token-1] [token-0-min-amount] [token-1-min-amount]",
		Short:               "create or add to existing concentrated liquidity position",
		Example:             "create-position [-69082] 69082 1000000000ufury 10000000uion 0 0 --pool-id 1 --from val --chain-id merlins-1",
		CustomFlagOverrides: poolIdFlagOverride,
		Flags:               furycli.FlagDesc{RequiredFlags: []*flag.FlagSet{FlagSetJustPoolId()}},
	}, &types.MsgCreatePosition{}
}

func NewWithdrawPositionCmd() (*furycli.TxCliDesc, *types.MsgWithdrawPosition) {
	return &furycli.TxCliDesc{
		Use:     "withdraw-position [position-id] [liquidity]",
		Short:   "withdraw from an existing concentrated liquidity position",
		Example: "withdraw-position 1 100317215 --from val --chain-id merlins-1",
	}, &types.MsgWithdrawPosition{}
}

func NewCollectFeesCmd() (*furycli.TxCliDesc, *types.MsgCollectFees) {
	return &furycli.TxCliDesc{
		Use:     "collect-fees [position-ids]",
		Short:   "collect fees from liquidity position(s)",
		Example: "collect-fees 1,5,7 --from val --chain-id merlins-1",
	}, &types.MsgCollectFees{}
}

func NewCollectIncentivesCmd() (*furycli.TxCliDesc, *types.MsgCollectIncentives) {
	return &furycli.TxCliDesc{
		Use:     "collect-incentives [position-ids]",
		Short:   "collect incentives from liquidity position(s)",
		Example: "collect-incentives 1,5,7 --from val --chain-id merlins-1",
	}, &types.MsgCollectIncentives{}
}

func NewCreateIncentiveCmd() (*furycli.TxCliDesc, *types.MsgCreateIncentive) {
	return &furycli.TxCliDesc{
		Use:                 "create-incentive [incentive-denom] [incentive-amount] [emission-rate] [start-time] [min-uptime]",
		Short:               "create an incentive record to emit incentives (per second) to a given pool",
		Example:             "create-incentive ufury 69082 0.02 100 2023-03-03 03:20:35.419543805 24h --pool-id 1 --from val --chain-id merlins-1",
		CustomFlagOverrides: poolIdFlagOverride,
		Flags:               furycli.FlagDesc{RequiredFlags: []*flag.FlagSet{FlagSetJustPoolId()}},
	}, &types.MsgCreateIncentive{}
}

func NewTickSpacingDecreaseProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tick-spacing-decrease-proposal [flags]",
		Args:  cobra.ExactArgs(0),
		Short: "Submit a tick spacing decrease proposal",
		Long: strings.TrimSpace(`Submit a tick spacing decrease proposal.

Passing in FlagPoolIdToTickSpacingRecords separated by commas would be parsed automatically to pairs of PoolIdToTickSpacing records.
Ex) --pool-tick-spacing-records=1,10,5,1 -> [(poolId 1, newTickSpacing 10), (poolId 5, newTickSpacing 1)]
Note: The new tick spacing value must be less than the current tick spacing value.

		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			content, err := parsePoolIdToTickSpacingRecordsArgsToContent(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")
	cmd.Flags().Bool(govcli.FlagIsExpedited, false, "If true, makes the proposal an expedited one")
	cmd.Flags().String(govcli.FlagProposal, "", "Proposal file path (if this path is given, other proposal flags are ignored)")
	cmd.Flags().String(FlagPoolIdToTickSpacingRecords, "", "The pool ID to new tick spacing records array")

	return cmd
}

func parsePoolIdToTickSpacingRecordsArgsToContent(cmd *cobra.Command) (govtypes.Content, error) {
	title, err := cmd.Flags().GetString(govcli.FlagTitle)
	if err != nil {
		return nil, err
	}

	description, err := cmd.Flags().GetString(govcli.FlagDescription)
	if err != nil {
		return nil, err
	}

	poolIdToTickSpacingRecords, err := parsePoolIdToTickSpacingRecords(cmd)
	if err != nil {
		return nil, err
	}

	content := &types.TickSpacingDecreaseProposal{
		Title:                      title,
		Description:                description,
		PoolIdToTickSpacingRecords: poolIdToTickSpacingRecords,
	}
	return content, nil
}

func parsePoolIdToTickSpacingRecords(cmd *cobra.Command) ([]types.PoolIdToTickSpacingRecord, error) {
	assetsStr, err := cmd.Flags().GetString(FlagPoolIdToTickSpacingRecords)
	if err != nil {
		return nil, err
	}

	assets := strings.Split(assetsStr, ",")

	if len(assets)%2 != 0 {
		return nil, fmt.Errorf("poolIdToTickSpacingRecords must be a list of pairs of poolId and newTickSpacing")
	}

	poolIdToTickSpacingRecords := []types.PoolIdToTickSpacingRecord{}
	i := 0
	for i < len(assets) {
		poolId, err := strconv.Atoi(assets[i])
		if err != nil {
			return nil, err
		}
		newTickSpacing, err := strconv.Atoi(assets[i+1])
		if err != nil {
			return nil, err
		}

		poolIdToTickSpacingRecords = append(poolIdToTickSpacingRecords, types.PoolIdToTickSpacingRecord{
			PoolId:         uint64(poolId),
			NewTickSpacing: uint64(newTickSpacing),
		})

		// increase counter by the next 2
		i = i + 2
	}

	return poolIdToTickSpacingRecords, nil
}
