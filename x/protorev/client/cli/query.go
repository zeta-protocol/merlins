package cli

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/merlins-labs/merlins/furyutils/furycli"

	"github.com/merlins-labs/merlins/v15/x/protorev/types"
)

// NewCmdQuery returns the cli query commands for this module
func NewCmdQuery() *cobra.Command {
	cmd := furycli.QueryIndexCmd(types.ModuleName)

	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryParamsCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryNumberOfTradesCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryProfitsByDenomCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryAllProfitsCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryStatisticsByRouteCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryAllRouteStatisticsCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryTokenPairArbRoutesCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryAdminAccountCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryDeveloperAccountCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryMaxPoolPointsPerTxCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryMaxPoolPointsPerBlockCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryBaseDenomsCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryEnabledCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryPoolWeightsCmd)
	furycli.AddQueryCmd(cmd, types.NewQueryClient, NewQueryPoolCmd)

	return cmd
}

// NewQueryParamsCmd returns the command to query the module params
func NewQueryParamsCmd() (*furycli.QueryDescriptor, *types.QueryParamsRequest) {
	return &furycli.QueryDescriptor{
		Use:   "params",
		Short: "Query the module params",
	}, &types.QueryParamsRequest{}
}

// NewQueryNumberOfTradesCmd returns the command to query the number of trades executed by protorev
func NewQueryNumberOfTradesCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevNumberOfTradesRequest) {
	return &furycli.QueryDescriptor{
		Use:   "number-of-trades",
		Short: "Query the number of cyclic arbitrage trades protorev has executed",
	}, &types.QueryGetProtoRevNumberOfTradesRequest{}
}

// NewQueryProfitsByDenomCmd returns the command to query the profits of protorev by denom
func NewQueryProfitsByDenomCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevProfitsByDenomRequest) {
	return &furycli.QueryDescriptor{
		Use:   "profits-by-denom [denom]",
		Short: "Query the profits of protorev by denom",
		Long:  `{{.Short}}{{.ExampleHeader}}{{.CommandPrefix}} profits-by-denom ufury`,
	}, &types.QueryGetProtoRevProfitsByDenomRequest{}
}

// NewQueryAllProfitsCmd returns the command to query all profits of protorev
func NewQueryAllProfitsCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevAllProfitsRequest) {
	return &furycli.QueryDescriptor{
		Use:   "all-profits",
		Short: "Query all ProtoRev profits",
	}, &types.QueryGetProtoRevAllProfitsRequest{}
}

// NewQueryStatisticsByRoute returns the command to query the statistics of protorev by route
func NewQueryStatisticsByRouteCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevStatisticsByRouteRequest) {
	return &furycli.QueryDescriptor{
		Use:                "statistics-by-route [route]",
		Short:              "Query statistics about a specific arbitrage route",
		Long:               `{{.Short}}{{.ExampleHeader}}{{.CommandPrefix}} statistics-by-route [1,2,3]`,
		CustomFieldParsers: map[string]furycli.CustomFieldParserFn{"Route": parseRoute},
	}, &types.QueryGetProtoRevStatisticsByRouteRequest{}
}

// NewQueryAllRouteStatisticsCmd returns the command to query all route statistics of protorev
func NewQueryAllRouteStatisticsCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevAllRouteStatisticsRequest) {
	return &furycli.QueryDescriptor{
		Use:   "all-statistics",
		Short: "Query all ProtoRev statistics",
	}, &types.QueryGetProtoRevAllRouteStatisticsRequest{}
}

// NewQueryTokenPairArbRoutesCmd returns the command to query the token pair arb routes
func NewQueryTokenPairArbRoutesCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevTokenPairArbRoutesRequest) {
	return &furycli.QueryDescriptor{
		Use:   "hot-routes",
		Short: "Query the ProtoRev hot routes currently being used",
	}, &types.QueryGetProtoRevTokenPairArbRoutesRequest{}
}

// NewQueryAdminAccountCmd returns the command to query the admin account
func NewQueryAdminAccountCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevAdminAccountRequest) {
	return &furycli.QueryDescriptor{
		Use:   "admin-account",
		Short: "Query the admin account",
	}, &types.QueryGetProtoRevAdminAccountRequest{}
}

// NewQueryDeveloperAccountCmd returns the command to query the developer account
func NewQueryDeveloperAccountCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevDeveloperAccountRequest) {
	return &furycli.QueryDescriptor{
		Use:   "developer-account",
		Short: "Query the developer account",
	}, &types.QueryGetProtoRevDeveloperAccountRequest{}
}

// NewQueryMaxPoolPointsPerTxCmd returns the command to query the max pool points per tx
func NewQueryMaxPoolPointsPerTxCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevMaxPoolPointsPerTxRequest) {
	return &furycli.QueryDescriptor{
		Use:   "max-pool-points-per-tx",
		Short: "Query the max pool points per tx",
	}, &types.QueryGetProtoRevMaxPoolPointsPerTxRequest{}
}

// NewQueryMaxPoolPointsPerBlockCmd returns the command to query the max pool points per block
func NewQueryMaxPoolPointsPerBlockCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevMaxPoolPointsPerBlockRequest) {
	return &furycli.QueryDescriptor{
		Use:   "max-pool-points-per-block",
		Short: "Query the max pool points per block",
	}, &types.QueryGetProtoRevMaxPoolPointsPerBlockRequest{}
}

// NewQueryBaseDenomsCmd returns the command to query the base denoms
func NewQueryBaseDenomsCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevBaseDenomsRequest) {
	return &furycli.QueryDescriptor{
		Use:   "base-denoms",
		Short: "Query the base denoms used to construct arbitrage routes",
	}, &types.QueryGetProtoRevBaseDenomsRequest{}
}

// NewQueryEnabled returns the command to query the enabled status of protorev
func NewQueryEnabledCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevEnabledRequest) {
	return &furycli.QueryDescriptor{
		Use:   "enabled",
		Short: "Query whether protorev is currently enabled",
	}, &types.QueryGetProtoRevEnabledRequest{}
}

// NewQueryPoolWeightsCmd returns the command to query the pool weights of protorev
func NewQueryPoolWeightsCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevPoolWeightsRequest) {
	return &furycli.QueryDescriptor{
		Use:   "pool-weights",
		Short: "Query the pool weights used to determine how computationally expensive a route is",
	}, &types.QueryGetProtoRevPoolWeightsRequest{}
}

// NewQueryPoolCmd returns the command to query the pool id for a given denom pair stored via the highest liquidity method in ProtoRev
func NewQueryPoolCmd() (*furycli.QueryDescriptor, *types.QueryGetProtoRevPoolRequest) {
	return &furycli.QueryDescriptor{
		Use:   "pool [base_denom] [other_denom]",
		Short: "Query the pool id for a given denom pair stored via the highest liquidity method in ProtoRev",
	}, &types.QueryGetProtoRevPoolRequest{}
}

// convert a string array "[1,2,3]" to []uint64
func parseRoute(arg string, _ *pflag.FlagSet) (any, furycli.FieldReadLocation, error) {
	var route []uint64
	err := json.Unmarshal([]byte(arg), &route)
	if err != nil {
		return nil, furycli.UsedArg, err
	}
	return route, furycli.UsedArg, err
}
