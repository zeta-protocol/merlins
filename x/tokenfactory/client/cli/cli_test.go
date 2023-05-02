package cli_test

import (
	"testing"

	"github.com/merlins-labs/merlins/furyutils/furycli"
	"github.com/merlins-labs/merlins/v15/x/tokenfactory/client/cli"
	"github.com/merlins-labs/merlins/v15/x/tokenfactory/types"
)

func TestGetCmdDenomAuthorityMetadata(t *testing.T) {
	desc, _ := cli.GetCmdDenomAuthorityMetadata()
	tcs := map[string]furycli.QueryCliTestCase[*types.QueryDenomAuthorityMetadataRequest]{
		"basic test": {
			Cmd: "uatom",
			ExpectedQuery: &types.QueryDenomAuthorityMetadataRequest{
				Denom: "uatom",
			},
		},
	}
	furycli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdDenomsFromCreator(t *testing.T) {
	desc, _ := cli.GetCmdDenomsFromCreator()
	tcs := map[string]furycli.QueryCliTestCase[*types.QueryDenomsFromCreatorRequest]{
		"basic test": {
			Cmd: "fury1test",
			ExpectedQuery: &types.QueryDenomsFromCreatorRequest{
				Creator: "fury1test",
			},
		},
	}
	furycli.RunQueryTestCases(t, desc, tcs)
}
