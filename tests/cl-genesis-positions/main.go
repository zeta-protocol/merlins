package main

import (
	"flag"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// operation defines the desired operation to be run by this script.
type operation int

const (
	// getData retrieves the data from the Uniswap subgraph and writes it to disk
	// under pathToFilesFromRoot + positionsFileName path.
	getData operation = iota
	// convertPositions converts the data from the Uniswap subgraph into Merlins
	// genesis. It reads pathToFilesFromRoot + positionsFileName path
	// run Merlins app via apptesting, creates positions and writes the genesis
	// under pathToFilesFromRoot + merlinsStateFileName path.
	convertPositions
	// mergeSubgraphAndLocalMerlinsGenesis merges the genesis created from the subgraph data
	// with the localmerlins genesis. This command is meant to be called inside the localmerlins
	// container during setup (see setup.sh). It reads the existing genesis from localmerlinsHomePath,
	// updates the concentrated liquidity section to append the CL pool created from the subgraph data,
	// its positions, ticks and accumulators.
	mergeSubgraphAndLocalMerlinsGenesis
)

const (
	pathToFilesFromRoot = "tests/cl-genesis-positions/"

	positionsFileName       = "subgraph_positions.json"
	merlinsGenesisFileName  = "genesis.json"
	bigbangPosiionsFileName = "bigbang_positions.json"

	localMerlinsHomePath = "/merlins/.merlinsd/"

	denom0 = "uusdc"
	denom1 = "ufury"
)

var (
	// This is lo-test1 address in localmerlins
	defaultCreatorAddresses = []sdk.AccAddress{sdk.MustAccAddressFromBech32("fury1cyyzpxplxdzkeea7kwsydadg87357qnahakaks"), sdk.MustAccAddressFromBech32("fury18s5lynnmx37hq4wlrw9gdn68sg2uxp5rgk26vv")}

	useKeyringAccounts bool

	writeGenesisToDisk bool

	writeBigBangConfigToDisk bool
)

func main() {
	var (
		desiredOperation int
		isLocalMerlins   bool
	)

	flag.BoolVar(&writeBigBangConfigToDisk, "big-bang", false, fmt.Sprintf("flag indicating whether to write the big bang config to disk at path %s", bigbangPosiionsFileName))
	flag.BoolVar(&writeGenesisToDisk, "genesis", false, fmt.Sprintf("flag indicating whether to write the genesis file to disk at path %s", merlinsGenesisFileName))
	flag.BoolVar(&useKeyringAccounts, "keyring", false, "flag indicating whether to use local test keyring accounts")
	flag.BoolVar(&isLocalMerlins, "localmerlins", false, "flag indicating whether this is being run inside the localmerlins container")
	flag.IntVar(&desiredOperation, "operation", 0, fmt.Sprintf("operation to run:\nget subgraph data: %v, convert subgraph positions to fury genesis: %v\nmerge converted subgraph genesis and localmerlins genesis: %v", getData, convertPositions, mergeSubgraphAndLocalMerlinsGenesis))

	flag.Parse()

	fmt.Println("isLocalMerlins:", isLocalMerlins)

	pathToSaveFilesAt := pathToFilesFromRoot
	if isLocalMerlins {
		pathToSaveFilesAt = ""
	}

	// Set this to one of the 'operation' values
	switch operation(desiredOperation) {
	// See definition for more info.
	case getData:
		fmt.Println("Getting data from Uniswap subgraph...")

		GetUniV3SubgraphData(pathToSaveFilesAt + positionsFileName)
		// See definition for more info.
	case convertPositions:
		fmt.Println("Converting positions from subgraph data to Merlins genesis...")

		var creatorAddresses []sdk.AccAddress
		if useKeyringAccounts {
			fmt.Println("Using local keyring addresses as creators")
			creatorAddresses = GetLocalKeyringAccounts()
		} else {
			fmt.Println("Using default creator addresses")
			creatorAddresses = defaultCreatorAddresses
		}

		ConvertSubgraphToMerlinsGenesis(creatorAddresses, pathToSaveFilesAt+positionsFileName)
		// See definition for more info.
	case mergeSubgraphAndLocalMerlinsGenesis:
		fmt.Println("Merging subgraph and local Merlins genesis...")
		clState, bankState := ConvertSubgraphToMerlinsGenesis(defaultCreatorAddresses, pathToSaveFilesAt+positionsFileName)

		EditLocalMerlinsGenesis(clState, bankState)
	default:
		panic("Invalid operation")
	}
}
