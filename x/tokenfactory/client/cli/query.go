package cli

import (

	// "strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/merlins-labs/merlins/furyutils/furycli"
	"github.com/merlins-labs/merlins/v15/x/tokenfactory/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	cmd := furycli.QueryIndexCmd(types.ModuleName)

	furycli.AddQueryCmd(cmd, types.NewQueryClient, GetCmdDenomAuthorityMetadata)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, GetCmdDenomAuthorityMetadata)

	cmd.AddCommand(
		furycli.GetParams[*types.QueryParamsRequest](
			types.ModuleName, types.NewQueryClient),
	)

	return cmd
}

func GetCmdDenomAuthorityMetadata() (*furycli.QueryDescriptor, *types.QueryDenomAuthorityMetadataRequest) {
	return &furycli.QueryDescriptor{
		Use:   "denom-authority-metadata [denom] [flags]",
		Short: "Get the authority metadata for a specific denom",
		Long: `{{.Short}}{{.ExampleHeader}}
		{{.CommandPrefix}} uatom`,
	}, &types.QueryDenomAuthorityMetadataRequest{}
}

func GetCmdDenomsFromCreator() (*furycli.QueryDescriptor, *types.QueryDenomsFromCreatorRequest) {
	return &furycli.QueryDescriptor{
		Use:   "denoms-from-creator [creator address] [flags]",
		Short: "Returns a list of all tokens created by a specific creator address",
		Long: `{{.Short}}{{.ExampleHeader}}
		{{.CommandPrefix}} <address>`,
	}, &types.QueryDenomsFromCreatorRequest{}
}

// GetCmdDenomAuthorityMetadata returns the authority metadata for a queried denom
func GetCmdDenomBeforeSendHook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "denom-before-send-hook [denom] [flags]",
		Short: "Get the BeforeSend hook for a specific denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.BeforeSendHookAddress(cmd.Context(), &types.QueryBeforeSendHookAddressRequest{
				Denom: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
