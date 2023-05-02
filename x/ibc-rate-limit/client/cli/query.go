package cli

import (
	"github.com/spf13/cobra"

	"github.com/merlins-labs/merlins/furyutils/furycli"
	"github.com/merlins-labs/merlins/v15/x/ibc-rate-limit/client/queryproto"
	"github.com/merlins-labs/merlins/v15/x/ibc-rate-limit/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := furycli.QueryIndexCmd(types.ModuleName)

	cmd.AddCommand(
		furycli.GetParams[*queryproto.ParamsRequest](
			types.ModuleName, queryproto.NewQueryClient),
	)

	return cmd
}
