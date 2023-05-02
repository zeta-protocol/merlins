package client

import (
	"github.com/merlins-labs/merlins/v15/x/concentrated-liquidity/client/cli"
	"github.com/merlins-labs/merlins/v15/x/concentrated-liquidity/client/rest"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	TickSpacingDecreaseProposalHandler = govclient.NewProposalHandler(cli.NewTickSpacingDecreaseProposal, rest.ProposalTickSpacingDecreaseRESTHandler)
)
