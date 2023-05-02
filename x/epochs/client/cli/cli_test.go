package cli_test

import (
	"testing"

	"github.com/merlins-labs/merlins/furyutils/furycli"
	"github.com/merlins-labs/merlins/x/epochs/client/cli"
	"github.com/merlins-labs/merlins/x/epochs/types"
)

func TestGetCmdCurrentEpoch(t *testing.T) {
	desc, _ := cli.GetCmdCurrentEpoch()
	tcs := map[string]furycli.QueryCliTestCase[*types.QueryCurrentEpochRequest]{
		"basic test": {
			Cmd: "day",
			ExpectedQuery: &types.QueryCurrentEpochRequest{
				Identifier: "day",
			},
		},
	}
	furycli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdEpochsInfo(t *testing.T) {
	desc, _ := cli.GetCmdEpochInfos()
	tcs := map[string]furycli.QueryCliTestCase[*types.QueryEpochsInfoRequest]{
		"basic test": {
			Cmd:           "",
			ExpectedQuery: &types.QueryEpochsInfoRequest{},
		},
	}
	furycli.RunQueryTestCases(t, desc, tcs)
}
