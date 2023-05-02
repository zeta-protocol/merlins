package cli

import (
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/merlins-labs/merlins/furyutils/furycli"
	"github.com/merlins-labs/merlins/v15/x/concentrated-liquidity/client/queryproto"
	"github.com/merlins-labs/merlins/v15/x/concentrated-liquidity/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := furycli.QueryIndexCmd(types.ModuleName)
	furycli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetCmdPools)
	furycli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetUserPositions)
	furycli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetClaimableFees)
	furycli.AddQueryCmd(cmd, queryproto.NewQueryClient, GetClaimableIncentives)
	cmd.AddCommand(
		furycli.GetParams[*queryproto.ParamsRequest](
			types.ModuleName, queryproto.NewQueryClient),
	)
	return cmd
}

func GetUserPositions() (*furycli.QueryDescriptor, *queryproto.UserPositionsRequest) {
	return &furycli.QueryDescriptor{
			Use:   "user-positions [address]",
			Short: "Query user's positions",
			Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} user-positions fury12smx2wdlyttvyzvzg54y2vnqwq2qjateuf7thj`,
			Flags:               furycli.FlagDesc{OptionalFlags: []*flag.FlagSet{FlagSetJustPoolId()}},
			CustomFlagOverrides: poolIdFlagOverride},
		&queryproto.UserPositionsRequest{}
}

func GetCmdPools() (*furycli.QueryDescriptor, *queryproto.PoolsRequest) {
	return &furycli.QueryDescriptor{
		Use:   "pools",
		Short: "Query pools",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} pools`}, &queryproto.PoolsRequest{}
}

func GetClaimableFees() (*furycli.QueryDescriptor, *queryproto.ClaimableFeesRequest) {
	return &furycli.QueryDescriptor{
		Use:   "claimable-fees [positionID]",
		Short: "Query claimable fees",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} claimable-fees 53`}, &queryproto.ClaimableFeesRequest{}
}

func GetClaimableIncentives() (*furycli.QueryDescriptor, *queryproto.ClaimableIncentivesRequest) {
	return &furycli.QueryDescriptor{
		Use:   "claimable-incentives [positionID]",
		Short: "Query claimable incentives",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} claimable-fees 53`}, &queryproto.ClaimableIncentivesRequest{}
}
