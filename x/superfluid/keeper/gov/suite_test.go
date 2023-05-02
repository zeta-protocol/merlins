package gov_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/merlins-labs/merlins/v15/app/apptesting"
	"github.com/merlins-labs/merlins/v15/x/superfluid/keeper"
	"github.com/merlins-labs/merlins/v15/x/superfluid/types"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	querier types.QueryServer
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.querier = keeper.NewQuerier(*suite.App.SuperfluidKeeper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
