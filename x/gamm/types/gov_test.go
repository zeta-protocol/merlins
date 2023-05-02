package types_test

import (
	"testing"

	proto "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"

	"github.com/merlins-labs/merlins/v15/x/gamm/types"
)

func TestUpdateMigrationRecordsProposalMarshalUnmarshal(t *testing.T) {
	tests := []struct {
		proposal *types.UpdateMigrationRecordsProposal
	}{
		{ // empty title
			proposal: &types.UpdateMigrationRecordsProposal{
				Title:       "",
				Description: "proposal to update migration records",
				Records:     []types.BalancerToConcentratedPoolLink(nil),
			},
		},
		{ // empty description
			proposal: &types.UpdateMigrationRecordsProposal{
				Title:       "title",
				Description: "",
				Records:     []types.BalancerToConcentratedPoolLink(nil),
			},
		},
		{ // empty records
			proposal: &types.UpdateMigrationRecordsProposal{
				Title:       "title",
				Description: "proposal to update migration records",
				Records:     []types.BalancerToConcentratedPoolLink(nil),
			},
		},
		{ // one record
			proposal: &types.UpdateMigrationRecordsProposal{
				Title:       "title",
				Description: "proposal to update migration records",
				Records: []types.BalancerToConcentratedPoolLink{
					{
						BalancerPoolId: 1,
						ClPoolId:       5,
					},
				},
			},
		},
		{ // two records
			proposal: &types.UpdateMigrationRecordsProposal{
				Title:       "title",
				Description: "proposal to update migration records",
				Records: []types.BalancerToConcentratedPoolLink{
					{
						BalancerPoolId: 1,
						ClPoolId:       5,
					},
					{
						BalancerPoolId: 2,
						ClPoolId:       6,
					},
				},
			},
		},
	}

	for _, test := range tests {
		bz, err := proto.Marshal(test.proposal)
		require.NoError(t, err)
		decoded := types.UpdateMigrationRecordsProposal{}
		err = proto.Unmarshal(bz, &decoded)
		require.NoError(t, err)
		require.Equal(t, *test.proposal, decoded)
	}
}
