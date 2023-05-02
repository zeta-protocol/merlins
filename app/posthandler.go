package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	protorevkeeper "github.com/merlins-labs/merlins/v15/x/protorev/keeper"
)

func NewPostHandler(protoRevKeeper *protorevkeeper.Keeper) sdk.AnteHandler {
	protoRevDecorator := protorevkeeper.NewProtoRevDecorator(*protoRevKeeper)
	return sdk.ChainAnteDecorators(protoRevDecorator)
}
