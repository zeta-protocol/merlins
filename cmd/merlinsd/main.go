package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	merlins "github.com/merlins-labs/merlins/v15/app"
	"github.com/merlins-labs/merlins/v15/app/params"
	"github.com/merlins-labs/merlins/v15/cmd/merlinsd/cmd"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, merlins.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
