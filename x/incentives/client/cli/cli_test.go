package cli

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/merlins-labs/merlins/furyutils"
	"github.com/merlins-labs/merlins/furyutils/furycli"
	"github.com/merlins-labs/merlins/v15/x/incentives/types"
)

var testAddresses = furyutils.CreateRandomAccounts(3)

func TestGetCmdGauges(t *testing.T) {
	desc, _ := GetCmdGauges()
	tcs := map[string]furycli.QueryCliTestCase[*types.GaugesRequest]{
		"basic test": {
			Cmd: "--offset=2",
			ExpectedQuery: &types.GaugesRequest{
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			},
		},
	}
	furycli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdToDistributeCoins(t *testing.T) {
	desc, _ := GetCmdToDistributeCoins()
	tcs := map[string]furycli.QueryCliTestCase[*types.ModuleToDistributeCoinsRequest]{
		"basic test": {
			Cmd: "", ExpectedQuery: &types.ModuleToDistributeCoinsRequest{},
		},
	}
	furycli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdGaugeByID(t *testing.T) {
	desc, _ := GetCmdGaugeByID()
	tcs := map[string]furycli.QueryCliTestCase[*types.GaugeByIDRequest]{
		"basic test": {
			Cmd: "1", ExpectedQuery: &types.GaugeByIDRequest{Id: 1},
		},
	}
	furycli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdActiveGauges(t *testing.T) {
	desc, _ := GetCmdActiveGauges()
	tcs := map[string]furycli.QueryCliTestCase[*types.ActiveGaugesRequest]{
		"basic test": {
			Cmd: "--offset=2",
			ExpectedQuery: &types.ActiveGaugesRequest{
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			}},
	}
	furycli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdActiveGaugesPerDenom(t *testing.T) {
	desc, _ := GetCmdActiveGaugesPerDenom()
	tcs := map[string]furycli.QueryCliTestCase[*types.ActiveGaugesPerDenomRequest]{
		"basic test": {
			Cmd: "ufury --offset=2",
			ExpectedQuery: &types.ActiveGaugesPerDenomRequest{
				Denom:      "ufury",
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			}},
	}
	furycli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdUpcomingGauges(t *testing.T) {
	desc, _ := GetCmdUpcomingGauges()
	tcs := map[string]furycli.QueryCliTestCase[*types.UpcomingGaugesRequest]{
		"basic test": {
			Cmd: "--offset=2",
			ExpectedQuery: &types.UpcomingGaugesRequest{
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			}},
	}
	furycli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdUpcomingGaugesPerDenom(t *testing.T) {
	desc, _ := GetCmdUpcomingGaugesPerDenom()
	tcs := map[string]furycli.QueryCliTestCase[*types.UpcomingGaugesPerDenomRequest]{
		"basic test": {
			Cmd: "ufury --offset=2",
			ExpectedQuery: &types.UpcomingGaugesPerDenomRequest{
				Denom:      "ufury",
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			}},
	}
	furycli.RunQueryTestCases(t, desc, tcs)
}
