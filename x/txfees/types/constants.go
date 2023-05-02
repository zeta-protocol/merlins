package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// ConsensusMinFee is a governance set parameter from prop 354 (https://www.mintscan.io/merlins/proposals/354)
// Its intended to be .0025 ufury / gas
var ConsensusMinFee sdk.Dec = sdk.NewDecWithPrec(25, 4)
