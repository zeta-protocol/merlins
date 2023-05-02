package cli

import (
	"github.com/spf13/cobra"

	"github.com/merlins-labs/merlins/furyutils/furycli"
	"github.com/merlins-labs/merlins/x/epochs/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := furycli.QueryIndexCmd(types.ModuleName)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, GetCmdEpochInfos)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, GetCmdCurrentEpoch)

	return cmd
}

func GetCmdEpochInfos() (*furycli.QueryDescriptor, *types.QueryEpochsInfoRequest) {
	return &furycli.QueryDescriptor{
		Use:   "epoch-infos",
		Short: "Query running epoch infos.",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}}`,
		QueryFnName: "EpochInfos"}, &types.QueryEpochsInfoRequest{}
}

func GetCmdCurrentEpoch() (*furycli.QueryDescriptor, *types.QueryCurrentEpochRequest) {
	return &furycli.QueryDescriptor{
		Use:   "current-epoch",
		Short: "Query current epoch by specified identifier.",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} day`}, &types.QueryCurrentEpochRequest{}
}
